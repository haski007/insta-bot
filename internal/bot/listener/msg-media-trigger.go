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
	"github.com/haski007/insta-bot/internal/clients/instloader"
	"github.com/haski007/insta-bot/pkg/file"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	tmpDirPath = "./resources"
	// Telegram media groups accept 2–10 items per album.
	telegramAlbumMax = 10
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

func usableMediaItems(items []instloader.MediaItem) []instloader.MediaItem {
	out := make([]instloader.MediaItem, 0, len(items))
	for _, item := range items {
		if item.IsVideo && item.VideoURL != "" {
			out = append(out, item)
			continue
		}
		if item.URL != "" {
			out = append(out, item)
		}
	}
	return out
}

func (rcv *InstaBotService) sendInstagramFallback(chatID int64, postInfo instloader.PostInfo, isVideo bool) {
	kind := "Post"
	if isVideo {
		kind = "Video"
	}
	message := fmt.Sprintf("📸 Instagram %s\n\n👤 @%s\n❤️ %d likes\n💬 %d comments\n\n%s",
		kind, postInfo.Owner, postInfo.Likes, postInfo.Comments, truncateCaption(postInfo.Caption, rcv.captionCharsLimit))
	if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] send fallback message")
	}
}

func (rcv *InstaBotService) sendSingleInstagramMedia(
	chatID int64,
	messageID int,
	item instloader.MediaItem,
	caption string,
	postInfo instloader.PostInfo,
) {
	if item.IsVideo && item.VideoURL != "" {
		videoFile, err := downloadVideo(item.VideoURL)
		if err != nil {
			rcv.log.WithError(err).Error("[msgInstagramTrigger] download video")
			rcv.sendInstagramFallback(chatID, postInfo, true)
			return
		}

		videoConfig := tgbotapi.NewVideo(chatID, videoFile)
		videoConfig.Caption = caption
		videoConfig.ReplyToMessageID = messageID

		if err := rcv.ReplyVideo(chatID, messageID, videoConfig, caption); err != nil {
			rcv.log.WithError(err).Error("[msgInstagramTrigger] reply video")
			rcv.sendInstagramFallback(chatID, postInfo, true)
		}
		return
	}

	if item.URL == "" {
		rcv.log.Errorf("[msgInstagramTrigger] no media URL available for post %s", postInfo.Shortcode)
		return
	}

	imageFile, err := downloadImage(item.URL)
	if err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] download image")
		rcv.sendInstagramFallback(chatID, postInfo, false)
		return
	}

	photoConfig := tgbotapi.NewPhoto(chatID, imageFile)
	photoConfig.Caption = caption
	photoConfig.ReplyToMessageID = messageID

	if _, err := rcv.bot.Send(photoConfig); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] send photo")
		rcv.sendInstagramFallback(chatID, postInfo, false)
	}
}

func buildAlbumInputMedia(items []instloader.MediaItem, caption string) ([]interface{}, error) {
	media := make([]interface{}, 0, len(items))
	for i, item := range items {
		if item.IsVideo && item.VideoURL != "" {
			videoFile, err := downloadVideo(item.VideoURL)
			if err != nil {
				return nil, fmt.Errorf("download carousel video %d: %w", i, err)
			}
			input := tgbotapi.NewInputMediaVideo(videoFile)
			if i == 0 && caption != "" {
				input.Caption = caption
			}
			media = append(media, input)
			continue
		}

		if item.URL == "" {
			return nil, fmt.Errorf("carousel item %d has no url", i)
		}
		imageFile, err := downloadImage(item.URL)
		if err != nil {
			return nil, fmt.Errorf("download carousel image %d: %w", i, err)
		}
		input := tgbotapi.NewInputMediaPhoto(imageFile)
		if i == 0 && caption != "" {
			input.Caption = caption
		}
		media = append(media, input)
	}
	return media, nil
}

func (rcv *InstaBotService) sendInstagramAlbum(
	chatID int64,
	messageID int,
	items []instloader.MediaItem,
	caption string,
	postInfo instloader.PostInfo,
) {
	for start := 0; start < len(items); start += telegramAlbumMax {
		end := start + telegramAlbumMax
		if end > len(items) {
			end = len(items)
		}
		chunk := items[start:end]

		// Telegram requires at least 2 items for a media group.
		if len(chunk) == 1 {
			chunkCaption := ""
			if start == 0 {
				chunkCaption = caption
			}
			rcv.sendSingleInstagramMedia(chatID, messageID, chunk[0], chunkCaption, postInfo)
			continue
		}

		chunkCaption := ""
		if start == 0 {
			chunkCaption = caption
		}
		media, err := buildAlbumInputMedia(chunk, chunkCaption)
		if err != nil {
			rcv.log.WithError(err).Error("[msgInstagramTrigger] build album media")
			rcv.sendInstagramFallback(chatID, postInfo, false)
			return
		}

		album := tgbotapi.NewMediaGroup(chatID, media)
		album.ReplyToMessageID = messageID
		if _, err := rcv.bot.SendMediaGroup(album); err != nil {
			rcv.log.WithError(err).Errorf("[msgInstagramTrigger] send media group (%d items)", len(chunk))
			rcv.sendInstagramFallback(chatID, postInfo, false)
			return
		}
	}
}

func (rcv *InstaBotService) msgInstagramTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := exprFindURL.FindString(update.Message.Text)
	fmt.Println("url", url)

	if !strings.Contains(url, postSubstring) && !strings.Contains(url, reelSubstring) {
		return
	}

	shortcode, err := extractShortcode(url)
	if err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] extract shortcode")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	postInfo, err := rcv.instloaderApi.GetPostInfo(shortcode)
	if err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] get post info from microservice")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	caption := ""
	items := usableMediaItems(postInfo.MediaItems())
	if len(items) == 0 {
		rcv.log.Errorf("[msgInstagramTrigger] no media URL available for post %s", postInfo.Shortcode)
		return
	}

	if len(items) == 1 {
		rcv.sendSingleInstagramMedia(chatID, messageID, items[0], caption, postInfo)
		return
	}

	rcv.log.Infof("[msgInstagramTrigger] sending carousel with %d slides for %s", len(items), postInfo.Shortcode)
	rcv.sendInstagramAlbum(chatID, messageID, items, caption, postInfo)
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
