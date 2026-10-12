package listener

import (
	"net/url"
	"strings"

	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isYouTubeShortsURL(text string) bool {
	return extractYouTubeShortsURL(text) != ""
}

func extractYouTubeShortsURL(text string) string {
	for _, raw := range exprFindURL.FindAllString(text, -1) {
		cleaned := strings.TrimRight(raw, ".,)!?]>")
		if isYouTubeShortsLink(cleaned) {
			return cleaned
		}
	}
	return ""
}

func isYouTubeShortsLink(raw string) bool {
	parsed, err := url.Parse(raw)
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	if host != "youtube.com" && !strings.HasSuffix(host, ".youtube.com") {
		return false
	}
	path := strings.ToLower(parsed.Path)
	return strings.Contains(path, "/shorts/")
}

func (rcv *InstaBotService) msgYoutubeShortsTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	rawURL := extractYouTubeShortsURL(update.Message.Text)
	if rawURL == "" {
		return
	}

	media, err := rcv.youtubeApi.GetMedia(rawURL)
	if err != nil {
		rcv.log.WithError(err).Error("[msgYoutubeShortsTrigger] get media")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}
	if media.VideoURL == "" {
		rcv.log.Errorf("[msgYoutubeShortsTrigger] no downloadable media for %s", rawURL)
		rcv.SendError(chatID, "Could not download this YouTube Short "+emoji.SadFace)
		return
	}

	rcv.log.Infof("[msgYoutubeShortsTrigger] sending youtube short %q (%s)", media.Title, media.Quality)
	videoFile, err := downloadVideo(media.VideoURL)
	if err != nil {
		rcv.log.WithError(err).Error("[msgYoutubeShortsTrigger] download video")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	videoConfig := tgbotapi.NewVideo(chatID, videoFile)
	videoConfig.ReplyToMessageID = messageID
	videoConfig.SupportsStreaming = true
	if err := rcv.ReplyVideo(chatID, messageID, videoConfig, ""); err != nil {
		rcv.log.WithError(err).Error("[msgYoutubeShortsTrigger] send video")
		rcv.SendError(chatID, ErrInternalServerError)
	}
}
