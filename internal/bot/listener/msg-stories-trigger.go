package listener

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// stories/username/MEDIA_ID — MEDIA_ID is numeric.
var storyURLRegexp = regexp.MustCompile(`(?i)instagram\.com/stories/([^/?#]+)/(\d+)`)

func extractStoryRef(rawURL string) (username, mediaID string, err error) {
	if strings.Contains(strings.ToLower(rawURL), "/stories/highlights/") {
		return "", "", fmt.Errorf("instagram highlights are not supported")
	}

	m := storyURLRegexp.FindStringSubmatch(rawURL)
	if len(m) < 3 {
		return "", "", fmt.Errorf("invalid Instagram story URL")
	}

	username = strings.TrimSpace(m[1])
	mediaID = m[2]
	if username == "" || mediaID == "" || strings.EqualFold(username, "highlights") {
		return "", "", fmt.Errorf("invalid Instagram story URL")
	}
	return username, mediaID, nil
}

func (rcv *InstaBotService) msgStoriesTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	rawURL := exprFindURL.FindString(update.Message.Text)

	username, mediaID, err := extractStoryRef(rawURL)
	if err != nil {
		rcv.log.WithError(err).Warnf("[msgStoriesTrigger] parse story url: %s", rawURL)
		if strings.Contains(err.Error(), "highlights") {
			rcv.SendError(chatID, "Instagram highlights are not supported yet "+emoji.SadFace)
			return
		}
		return
	}

	storyInfo, err := rcv.instloaderApi.GetStoryInfo(mediaID, username)
	if err != nil {
		rcv.log.WithError(err).Error("[msgStoriesTrigger] get story info from microservice")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}
	if storyInfo.Owner == "" {
		storyInfo.Owner = username
	}

	items := usableMediaItems(storyInfo.MediaItems())
	if len(items) == 0 {
		rcv.log.Errorf("[msgStoriesTrigger] no media URL available for story %s/%s", username, mediaID)
		rcv.SendError(chatID, "Story not available (expired, private, or blocked) "+emoji.SadFace)
		return
	}

	rcv.log.Infof("[msgStoriesTrigger] sending story @%s media_id=%s (%d item(s))", username, mediaID, len(items))
	if len(items) == 1 {
		rcv.sendSingleInstagramMedia(chatID, messageID, items[0], "", storyInfo)
		return
	}
	rcv.sendInstagramAlbum(chatID, messageID, items, "", storyInfo)
}
