package listener

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/haski007/insta-bot/internal/bot/model"
	"github.com/haski007/insta-bot/pkg/file"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	tmpDirPath = "./resources"
)

var exprFindURL = regexp.MustCompile(`https?://[^\s]+`)

const (
	postSubstring = "/p/"
	reelSubstring = "/reel/"
)

// extractShortcode extracts the shortcode from Instagram URL
func extractShortcode(url string) (string, error) {
	// Remove query parameters
	if idx := strings.Index(url, "?"); idx != -1 {
		url = url[:idx]
	}

	// Remove trailing slash
	url = strings.TrimSuffix(url, "/")

	// Extract shortcode from /p/SHORTCODE or /reel/SHORTCODE
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid Instagram URL format")
	}

	shortcode := parts[len(parts)-1]
	if shortcode == "" {
		return "", fmt.Errorf("no shortcode found in URL")
	}

	return shortcode, nil
}

// downloadVideo downloads a video from URL and returns file bytes
func downloadVideo(videoURL string) (tgbotapi.FileBytes, error) {
	resp, err := http.Get(videoURL)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("download video: %w", err)
	}
	defer resp.Body.Close()

	videoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("read video bytes: %w", err)
	}

	return tgbotapi.FileBytes{
		Name:  fmt.Sprintf("video_%d.mp4", time.Now().UnixNano()),
		Bytes: videoBytes,
	}, nil
}

// downloadImage downloads an image from URL and returns file bytes
func downloadImage(imageURL string) (tgbotapi.FileBytes, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("download image: %w", err)
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("read image bytes: %w", err)
	}

	return tgbotapi.FileBytes{
		Name:  fmt.Sprintf("image_%d.jpg", time.Now().UnixNano()),
		Bytes: imageBytes,
	}, nil
}

// truncateCaption truncates caption to fit within character limit
func truncateCaption(caption string, limit int) string {
	if len(caption) <= limit {
		return caption
	}

	// Try to truncate at word boundary
	words := strings.Fields(caption)
	result := ""
	for _, word := range words {
		if len(result+" "+word) <= limit-3 { // -3 for "..."
			if result != "" {
				result += " "
			}
			result += word
		} else {
			break
		}
	}

	if result != caption {
		result += "..."
	}

	return result
}

func (rcv *InstaBotService) msgInstagramTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := exprFindURL.FindString(update.Message.Text)

	// if update.Message.From.UserName == "dmlitvin" {
	// 	rcv.Reply(chatID, messageID, "Хуй тобі, а не відео")
	// 	return
	// }

	if !strings.Contains(url, postSubstring) && !strings.Contains(url, reelSubstring) {
		return
	}

	// Extract shortcode from the URL
	shortcode, err := extractShortcode(url)
	if err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] extract shortcode")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	// Get post info from Python microservice using the client
	postInfo, err := rcv.instloaderApi.GetPostInfo(shortcode)
	if err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] get post info from microservice")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	// Send message with post info
	caption := fmt.Sprintf("❤️ Likes: %d\n%s", postInfo.Likes, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))

	// Download and send video if it's a video post
	if postInfo.IsVideo && postInfo.VideoURL != "" {
		// Download the video
		videoFile, err := downloadVideo(postInfo.VideoURL)
		if err != nil {
			rcv.log.WithError(err).Error("[msgInstagramTrigger] download video")
			// Fallback to text message if download fails
			message := fmt.Sprintf("📸 Instagram Video\n\n👤 @%s\n❤️ %d likes\n💬 %d comments\n\n%s",
				postInfo.Owner, postInfo.Likes, postInfo.Comments, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))
			if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
				rcv.log.WithError(err).Error("[msgInstagramTrigger] send fallback message")
			}
			return
		}

		videoConfig := tgbotapi.NewVideo(chatID, videoFile)
		videoConfig.Caption = caption
		videoConfig.ReplyToMessageID = messageID

		if err := rcv.ReplyVideo(chatID, messageID, videoConfig, caption); err != nil {
			rcv.log.WithError(err).Error("[msgInstagramTrigger] reply video")
			// Fallback to text message if video fails
			message := fmt.Sprintf("📸 Instagram Video\n\n👤 @%s\n❤️ %d likes\n💬 %d comments\n\n%s",
				postInfo.Owner, postInfo.Likes, postInfo.Comments, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))
			if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
				rcv.log.WithError(err).Error("[msgInstagramTrigger] send fallback message")
			}
		}
	} else {
		// Handle non-video content (images)
		if postInfo.URL != "" {
			// Download the image
			imageFile, err := downloadImage(postInfo.URL)
			if err != nil {
				rcv.log.WithError(err).Error("[msgInstagramTrigger] download image")
				// Fallback to text message if download fails
				message := fmt.Sprintf("📸 Instagram Post\n\n👤 @%s\n❤️ %d likes\n💬 %d comments\n\n%s",
					postInfo.Owner, postInfo.Likes, postInfo.Comments, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))
				if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
					rcv.log.WithError(err).Error("[msgInstagramTrigger] send fallback message")
				}
				return
			}

			// Send the image with caption
			photoConfig := tgbotapi.NewPhoto(chatID, imageFile)
			photoConfig.Caption = caption
			photoConfig.ReplyToMessageID = messageID

			if _, err := rcv.bot.Send(photoConfig); err != nil {
				rcv.log.WithError(err).Error("[msgInstagramTrigger] send photo")
				// Fallback to text message if photo fails
				message := fmt.Sprintf("📸 Instagram Post\n\n👤 @%s\n❤️ %d likes\n💬 %d comments\n\n%s",
					postInfo.Owner, postInfo.Likes, postInfo.Comments, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))
				if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
					rcv.log.WithError(err).Error("[msgInstagramTrigger] send fallback message")
				}
			}
		} else {
			rcv.log.Error("[msgInstagramTrigger] no media URL available for post %s", postInfo.Shortcode)
		}
	}
}

func downloadAndGetVideoFilesBytes(videos []*model.Video) ([]interface{}, error) {
	var downloadedFilesBytes []interface{}
	for _, v := range videos {
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ".mp4")
		filePath, err := v.DownloadAsFile(tmpDirPath, fileName)
		if err != nil {
			return nil, fmt.Errorf("download video err: %w", err)
		}
		fileBytes, err := getFileBytes(filePath, fileName)
		if err != nil {
			return nil, fmt.Errorf("get file bytes err: %w", err)
		}

		if err := file.DeleteFile(filePath); err != nil {
			return nil, fmt.Errorf("[msgInstagramTrigger] image delete file err: %w", err)
		}

		downloadedFilesBytes = append(downloadedFilesBytes, tgbotapi.NewInputMediaVideo(fileBytes))
	}
	return downloadedFilesBytes, nil
}

func downloadAndGetImageFilesBytes(videos []*model.Image) ([]interface{}, error) {
	var downloadedFilesBytes []interface{}
	for _, v := range videos {
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ".jpg")
		filePath, err := v.DownloadAsFile(tmpDirPath, fileName)
		if err != nil {
			return nil, fmt.Errorf("download video err: %w", err)
		}
		fileBytes, err := getFileBytes(filePath, fileName)
		if err != nil {
			return nil, fmt.Errorf("get file bytes err: %w", err)
		}

		if err := file.DeleteFile(filePath); err != nil {
			return nil, fmt.Errorf("[msgInstagramTrigger] image delete file err: %w", err)
		}

		downloadedFilesBytes = append(downloadedFilesBytes, tgbotapi.NewInputMediaPhoto(fileBytes))
	}
	return downloadedFilesBytes, nil
}

func getFileBytes(filepath, name string) (tgbotapi.FileBytes, error) {
	photoBytes, err := os.ReadFile(filepath)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("read file err: %w", err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  name,
		Bytes: photoBytes,
	}
	return photoFileBytes, nil
}
