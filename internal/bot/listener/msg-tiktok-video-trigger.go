package listener

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/pkg/file"
	"github.com/haski007/insta-bot/pkg/text"
	"mvdan.cc/xurls/v2"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgTikTokTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	fmt.Println("zalupa")

	xurlsStrict := xurls.Strict()
	output := xurlsStrict.FindAllString(update.Message.Text, -1)
	if len(output) < 1 {
		rcv.SendError(chatID, "invalid url in message")
		return
	}
	url := update.Message.Text

	videoData, err := rcv.tiktokApi.GetVideoDataFromUrl(url)
	if err != nil {
		rcv.log.WithError(err).Error("[msgTikTokTrigger] send media group")
		return
	}

	fileName := videoData.GetAuthor() + strconv.FormatInt(time.Now().Unix(), 10)
	videoPath, err := videoData.DownloadAsFile(os.TempDir(), fileName)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		if errC := rcv.NotifyCreator(fmt.Sprintf("[msgTikTokTrigger] can't download as file url: %s\n"+
			"from chat with: @%s\n",
			videoData.GetDownloadUrl(),
			update.Message.From.UserName)); errC != nil {
			rcv.log.WithError(err).Println("[msgTikTokTrigger] notify creator")
		}
		rcv.log.WithError(err).Error("[msgTikTokTrigger] download as file")
		return
	}

	fileBytes, err := getVideoFileBytes(videoPath, fileName)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		if errC := rcv.NotifyCreator(fmt.Sprintf("[msgTikTokTrigger] can't get file bytes videoPath: %s\n"+
			"from chat with: @%s\n",
			videoPath,
			update.Message.From.UserName)); errC != nil {
			rcv.log.WithError(err).Println("[msgTikTokTrigger] notify creator")
		}
		rcv.log.WithError(err).Error("[msgTikTokTrigger] download as file")
		return
	}

	var message = fmt.Sprintf("TikTokVideo:\n"+
		"Author: [%s](%s)\n"+
		"Source: [tiktok](%s)\n"+
		"title/caption:\n%s\n\n"+
		"shared by: @%s %s\n",
		videoData.GetAuthor(), publisher.TikTokBaseUrl+"@"+videoData.GetAuthor(),
		url,
		text.CharLimiterToWord(videoData.GetTitle(), rcv.captionCharsLimit),
		update.Message.From.UserName, update.Message.From.FirstName+" "+update.Message.From.LastName)
	videoMessage := tgbotapi.NewVideo(chatID, fileBytes)
	videoMessage.ParseMode = tgbotapi.ModeMarkdown
	videoMessage.Caption = message
	if _, err := rcv.bot.Send(videoMessage); err != nil {
		rcv.log.WithError(err).Error("[msgTikTokTrigger] send video message")
		return
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgTikTokTrigger] delete message")
	}
}
func getVideoFileBytes(filepath, name string) (tgbotapi.FileBytes, error) {
	defer file.DeleteFile(filepath)
	photoBytes, err := os.ReadFile(filepath)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("read file err: %w")
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  name,
		Bytes: photoBytes,
	}
	return photoFileBytes, nil
}
