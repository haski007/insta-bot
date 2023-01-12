package listener

import (
	"errors"
	"fmt"
	"os"

	"github.com/haski007/insta-bot/internal/clients/instapi"

	"github.com/haski007/insta-bot/pkg/text"

	"github.com/haski007/insta-bot/internal/bot"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/pkg/file"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	tmpDirPath = "./resources"
)

func (rcv *InstaBotService) msgMediaTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := update.Message.Text

	mediaInfo, err := rcv.instapi.GetMediaInfoFromURL(rcv.ctx, url)
	if err != nil {
		rcv.log.WithError(err).Println("[msgMediaTrigger] GetMediaInfoFromURL")
		rcv.SendError(chatID, fmt.Sprintf("can't get media from url [%s]", url))
		return
	}
	var downloadedFiles []interface{}
	mediaType := publisher.MediaType(mediaInfo.GetMediaType())
	switch mediaType {
	case publisher.Video:
		name := *mediaInfo.GetUser().Username + mediaInfo.GetPk()

		fileData, err := downloadResource(mediaInfo.GetVideoUrl(), name, mediaType)
		if err != nil {
			if errors.Is(err, bot.ErrWrongFileFormat) {
				rcv.SendError(chatID, "wrong file format, or not supported yet, write please to creator: @pdemian")
			}
			rcv.log.WithError(err).
				WithFields(map[string]interface{}{
					"mediaType":   mediaInfo.GetMediaType(),
					"productType": mediaInfo.GetProductType(),
				}).
				Println("[msgMediaTrigger] download resource")
			return
		}

		downloadedFiles = append(downloadedFiles, fileData)

	case publisher.Photo:
		name := *mediaInfo.GetUser().Username + mediaInfo.GetPk()

		fileData, err := downloadResource(mediaInfo.GetThumbnailUrl(), name, mediaType)
		if err != nil {
			if errors.Is(err, bot.ErrWrongFileFormat) {
				rcv.SendError(chatID, "wrong file format, or not supported yet, write please to creator: @pdemian")
			}
			rcv.log.WithError(err).
				WithFields(map[string]interface{}{
					"mediaType":   mediaInfo.GetMediaType(),
					"productType": mediaInfo.GetProductType(),
				}).
				Println("[msgMediaTrigger] download resource")
			return
		}

		downloadedFiles = append(downloadedFiles, fileData)

	case publisher.Album:
		downloaded, err := downloadResources(mediaInfo)
		if err != nil {
			if errors.Is(err, bot.ErrWrongFileFormat) {
				rcv.SendError(chatID, "wrong file format, or not supported yet, write please to creator: @pdemian")
			}
			rcv.log.WithError(err).
				WithFields(map[string]interface{}{
					"mediaType":   mediaInfo.GetMediaType(),
					"productType": mediaInfo.GetProductType(),
				}).
				Println("[msgMediaTrigger] download resources")
			return
		}
		downloadedFiles = append(downloadedFiles, downloaded...)
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgMediaTrigger] delete message")
	}

	mdg := tgbotapi.NewMediaGroup(chatID, downloadedFiles)
	if _, err := rcv.bot.SendMediaGroup(mdg); err != nil {
		rcv.log.WithError(err).Error("[msgMediaTrigger] send media group")
		return
	}

	var message = fmt.Sprintf("Media from author: [%s %s](%s)\n\n"+
		"Description: %s\n\n"+
		"source: [instagram](%s)\n\n"+
		"shared by: @%s %s",
		*mediaInfo.GetUser().Username,
		*mediaInfo.GetUser().FullName,
		publisher.InstagramBaseUrl+*mediaInfo.GetUser().Username,
		text.CharLimiterToWord(mediaInfo.GetCaptionText(), rcv.captionCharsLimit),
		url,
		update.Message.From.UserName, update.Message.From.FirstName+" "+update.Message.From.LastName)
	if err := rcv.SendMessage(chatID, message); err != nil {
		rcv.log.WithError(err).Error("[msgMediaTrigger] send message caption")
		return
	}
}

func downloadResources(mediaInfo *instapi.Media) (filesData []interface{}, err error) {
	for _, r := range mediaInfo.Resources {
		name := *mediaInfo.GetUser().Username + r.GetPk()

		fileData, err := downloadResource(r.GetThumbnailUrl(), name, publisher.MediaType(r.GetMediaType()))
		if err != nil {
			return nil, fmt.Errorf("download resource err: %w", err)
		}

		filesData = append(filesData, fileData)
	}

	return
}

func downloadResource(url string, name string, mediaType publisher.MediaType) (interface{}, error) {
	var extention string
	switch mediaType {
	case publisher.Photo:
		extention = ".jpg"
	case publisher.Video:
		extention = ".mp4"

	}

	filepath, err := file.Download(url, tmpDirPath, name+extention)
	if err != nil {
		return nil, fmt.Errorf("[msgMediaTrigger] download media file err: %w", err)
	}

	switch mediaType {
	case publisher.Photo:
		photoBytes, err := getFileBytes(filepath, name)
		if err != nil {
			return nil, fmt.Errorf("[msgMediaTrigger] photo getFileBytes %w", err)
		}
		if err := file.DeleteFile(filepath); err != nil {
			return nil, fmt.Errorf("[msgMediaTrigger] photo delete file err: %w", err)
		}

		return tgbotapi.NewInputMediaPhoto(photoBytes), nil
	case publisher.Video:
		videoBytes, err := getFileBytes(filepath, name)
		if err != nil {
			return nil, fmt.Errorf("[msgMediaTrigger] video getFileBytes %w", err)
		}
		if err := file.DeleteFile(filepath); err != nil {
			return nil, fmt.Errorf("[msgMediaTrigger] video delete file err: %w", err)
		}

		return tgbotapi.NewInputMediaVideo(videoBytes), nil
	}

	return nil, bot.ErrWrongFileFormat
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
