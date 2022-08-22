package listener

import (
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/internal/bot/publisher"
)

func (rcv *InstaBotService) msgStoriesTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := update.Message.Text

	storyInfo, err := rcv.instapi.GetStoryInfoFromUrl(rcv.ctx, url)
	if err != nil {
		rcv.log.WithError(err).Println("[msgStoriesTrigger] GetStoryInfoFromUrl")
		if err := rcv.SendError(chatID, fmt.Sprintf("can't get media from url [%s]", url)); err != nil {
			rcv.log.WithError(err).Println("[msgStoriesTrigger] send error")
		}
		return
	}
	var downloadedFiles []interface{}
	mediaType := publisher.MediaType(storyInfo.GetMediaType())
	switch mediaType {
	case publisher.Video:
		name := *storyInfo.GetUser().Username + storyInfo.GetPk()

		fileData, err := downloadResource(storyInfo.GetVideoUrl(), name, mediaType)
		if err != nil {
			if errors.Is(err, bot.ErrWrongFileFormat) {
				rcv.SendError(chatID, "wrong file format, or not supported yet, write please to creator: @pdemian")
			}
			rcv.log.WithError(err).
				WithFields(map[string]interface{}{
					"mediaType":   storyInfo.GetMediaType(),
					"productType": storyInfo.GetProductType(),
				}).
				Println("[msgStoriesTrigger] download resource")
			return
		}

		downloadedFiles = append(downloadedFiles, fileData)

	case publisher.Photo:
		name := *storyInfo.GetUser().Username + storyInfo.GetPk()

		fileData, err := downloadResource(storyInfo.GetThumbnailUrl(), name, mediaType)
		if err != nil {
			if errors.Is(err, bot.ErrWrongFileFormat) {
				rcv.SendError(chatID, "wrong file format, or not supported yet, write please to creator: @pdemian")
			}
			rcv.log.WithError(err).
				WithFields(map[string]interface{}{
					"mediaType":   storyInfo.GetMediaType(),
					"productType": storyInfo.GetProductType(),
				}).
				Println("[msgStoriesTrigger] download resource")
			return
		}

		downloadedFiles = append(downloadedFiles, fileData)
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgStoriesTrigger] delete message")
	}

	mdg := tgbotapi.NewMediaGroup(chatID, downloadedFiles)
	if _, err := rcv.bot.SendMediaGroup(mdg); err != nil {
		rcv.log.WithError(err).Error("[msgStoriesTrigger] send media group")
		return
	}

	var message = fmt.Sprintf("Story:\n"+
		"Author: [%s %s](%s)\n"+
		"Source: [instagram](%s)\n\n"+
		"shared by: @%s %s",
		*storyInfo.GetUser().Username,
		*storyInfo.GetUser().FullName,
		fmt.Sprintf(publisher.InstagramBaseUrl+*storyInfo.GetUser().Username),
		url,
		update.Message.From.UserName, update.Message.From.FirstName+" "+update.Message.From.LastName)
	if err := rcv.SendMessage(chatID, message); err != nil {
		rcv.log.WithError(err).Error("[msgStoriesTrigger] send message caption")
		return
	}
}
