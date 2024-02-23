package listener

import (
	"fmt"
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

func (rcv *InstaBotService) msgInstagramTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := exprFindURL.FindString(update.Message.Text)

	if !strings.Contains(url, postSubstring) && !strings.Contains(url, reelSubstring) {
		return
	}

	url = strings.ReplaceAll(url, "https://www.instagram.com", "https://www.ddinstagram.com")

	if err := rcv.SendMessageWithoutMarkdown(chatID, fmt.Sprintf("forwarder: @%s\n\nurl: %s", update.Message.From.UserName, url)); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] send message caption")
		return
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramTrigger] delete message")
		return
	}

	//content, err := rcv.instapi.GetPostContent(url)
	//if err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] get post content")
	//	rcv.SendError(chatID, ErrInternalServerError)
	//	return
	//}
	//
	//// ---> download videos and photos
	//videosBytes, err := downloadAndGetVideoFilesBytes(content.Video)
	//if err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] download videos")
	//	rcv.SendError(chatID, ErrInternalServerError)
	//	return
	//}
	//
	//imagesBytes, err := downloadAndGetImageFilesBytes(content.Image)
	//if err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] download images")
	//	rcv.SendError(chatID, ErrInternalServerError)
	//	return
	//}
	//
	//downloadedFilesBytes := append(videosBytes, imagesBytes...)
	//
	//mdg := tgbotapi.NewMediaGroup(chatID, downloadedFilesBytes)
	//if _, err := rcv.bot.SendMediaGroup(mdg); err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] send media group")
	//	return
	//}
	//
	//content.ArticleBody = escapeMarkdown(content.ArticleBody)
	//content.Author.Name = escapeMarkdown(content.Author.Name)
	//content.Author.Identifier.Value = escapeMarkdown(content.Author.Identifier.Value)
	//
	//var message = fmt.Sprintf("Instagram post  from author: [%s](%s)\n\n"+
	//	"Description: %s\n\n"+
	//	"source: [instagram](%s)\n\n"+
	//	"shared by: @%s %s",
	//	content.Author.Name,
	//	content.Author.GetProfileURL(),
	//	text.CharLimiterToWord(content.ArticleBody, rcv.captionCharsLimit),
	//	url,
	//	update.Message.From.UserName, update.Message.From.FirstName+" "+update.Message.From.LastName)
	//if err := rcv.SendMessage(chatID, message); err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] send message caption")
	//	return
	//}
	//
	//if err := rcv.DeleteMessage(chatID, messageID); err != nil {
	//	rcv.log.WithError(err).Error("[msgInstagramTrigger] delete message")
	//}
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
