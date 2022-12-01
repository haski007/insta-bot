package listener

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/haski007/insta-bot/pkg/text"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgYoutubeTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := update.Message.Text

	rcv.log.WithFields(map[string]interface{}{
		"from": update.Message.From.UserName,
		"chat": update.Message.Chat.Title,
	}).Debugln("received url to download video...")
	videoData, err := rcv.youtubeApi.GetVideoByUrl(url, 1080)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] get video by url")
		return
	}

	if videoData.GetDownloadUrl() == "" {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] get video by url")
	}

	fileName := videoData.Author + strconv.FormatInt(time.Now().UnixNano(), 10) + ".mp4"
	rcv.log.WithFields(map[string]interface{}{
		"from":        update.Message.From.UserName,
		"chat":        update.Message.Chat.Title,
		"video_title": videoData.Title,
		"filename":    fileName,
	}).Debugln("downloading video...")
	videoPath, err := videoData.DownloadAsFile(os.TempDir(), fileName)
	if err != nil {
		if errC := rcv.NotifyCreator(fmt.Sprintf("[msgYoutubeTrigger] can't download as file url: %s\n"+
			"from chat with: @%s\n",
			videoData.GetDownloadUrl(),
			update.Message.From.UserName)); errC != nil {
			rcv.log.WithError(err).Println("[msgYoutubeTrigger] notify creator")
		}
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] download as file")
		return
	}

	fileBytes, err := getVideoFileBytes(videoPath, fileName)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		if errC := rcv.NotifyCreator(fmt.Sprintf("[msgYoutubeTrigger] can't get file bytes videoPath: %s\n"+
			"from chat with: @%s\n",
			videoPath,
			update.Message.From.UserName)); errC != nil {
			rcv.log.WithError(err).Println("[msgYoutubeTrigger] notify creator")
		}
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] download as file")
		return
	}
	rcv.log.WithFields(map[string]interface{}{
		"from":        update.Message.From.UserName,
		"chat":        update.Message.Chat.Title,
		"video_title": videoData.Title,
	}).Debugln("Sending video to a chat...")
	var message = fmt.Sprintf("YouTube video:\n"+
		"Author: *%s*\n"+
		"Source: [youtube](%s)\n"+
		"title/caption: *%s*\n"+
		"Published: *%s*\nDuration: *%s*\nQuality: *%s*\n\n"+
		"shared by: *@%s %s*\n",
		videoData.GetAuthor(),
		url,
		text.CharLimiterToWord(videoData.GetTitle(), rcv.captionCharsLimit),
		videoData.CreatedAt.Format("02-01-2006 15:04:05"), videoData.Duration.String(), videoData.QualityLabel,
		update.Message.From.UserName, update.Message.From.FirstName+" "+update.Message.From.LastName)
	videoMessage := tgbotapi.NewVideo(chatID, fileBytes)
	videoMessage.ParseMode = tgbotapi.ModeMarkdown
	videoMessage.Caption = message
	if _, err := rcv.bot.Send(videoMessage); err != nil {
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] send video message")
		return
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgYoutubeTrigger] delete message")
	}
}
