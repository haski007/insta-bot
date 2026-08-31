package listener

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/haski007/insta-bot/internal/clients/instloader"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/haski007/insta-bot/pkg/file"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isTikTokURL(text string) bool {
	return extractTikTokURL(text) != ""
}

func extractTikTokURL(text string) string {
	for _, raw := range exprFindURL.FindAllString(text, -1) {
		cleaned := strings.TrimRight(raw, ".,)!?]>")
		if isTikTokHost(cleaned) {
			return cleaned
		}
	}
	return ""
}

func isTikTokHost(raw string) bool {
	parsed, err := url.Parse(raw)
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	return host == "tiktok.com" || strings.HasSuffix(host, ".tiktok.com")
}

func tiktokMediaItems(media *tiktokapi.Media) []instloader.MediaItem {
	if media == nil {
		return nil
	}
	if len(media.Images) > 0 {
		items := make([]instloader.MediaItem, 0, len(media.Images))
		for _, img := range media.Images {
			items = append(items, instloader.MediaItem{URL: img})
		}
		return usableMediaItems(items)
	}
	if media.VideoURL == "" {
		return nil
	}
	return []instloader.MediaItem{{
		IsVideo:  true,
		VideoURL: media.VideoURL,
	}}
}

func (rcv *InstaBotService) msgTikTokTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	rawURL := extractTikTokURL(update.Message.Text)
	if rawURL == "" {
		return
	}

	media, err := rcv.tiktokApi.GetMedia(rawURL)
	if err != nil {
		rcv.log.WithError(err).Error("[msgTikTokTrigger] get media")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	items := tiktokMediaItems(media)
	if len(items) == 0 {
		rcv.log.Errorf("[msgTikTokTrigger] no downloadable media for %s", rawURL)
		rcv.SendError(chatID, "Could not download this TikTok "+emoji.SadFace)
		return
	}

	rcv.log.Infof("[msgTikTokTrigger] sending tiktok @%s id=%s (%d item(s))", media.Author, media.ID, len(items))
	if err := rcv.sendTikTokMedia(chatID, messageID, items); err != nil {
		rcv.log.WithError(err).Error("[msgTikTokTrigger] send media")
		rcv.SendError(chatID, ErrInternalServerError)
	}
}

func (rcv *InstaBotService) sendTikTokMedia(chatID int64, messageID int, items []instloader.MediaItem) error {
	for start := 0; start < len(items); start += telegramAlbumMax {
		end := start + telegramAlbumMax
		if end > len(items) {
			end = len(items)
		}
		chunk := items[start:end]

		if len(chunk) == 1 {
			if err := rcv.sendSingleTikTokItem(chatID, messageID, chunk[0]); err != nil {
				return err
			}
			continue
		}

		media, err := buildAlbumInputMedia(chunk, "")
		if err != nil {
			return fmt.Errorf("build album: %w", err)
		}
		album := tgbotapi.NewMediaGroup(chatID, media)
		album.ReplyToMessageID = messageID
		if _, err := rcv.bot.SendMediaGroup(album); err != nil {
			return fmt.Errorf("send media group: %w", err)
		}
	}
	return nil
}

func (rcv *InstaBotService) sendSingleTikTokItem(chatID int64, messageID int, item instloader.MediaItem) error {
	if item.IsVideo && item.VideoURL != "" {
		videoFile, err := downloadVideo(item.VideoURL)
		if err != nil {
			return fmt.Errorf("download video: %w", err)
		}
		videoConfig := tgbotapi.NewVideo(chatID, videoFile)
		videoConfig.ReplyToMessageID = messageID
		videoConfig.SupportsStreaming = true
		return rcv.ReplyVideo(chatID, messageID, videoConfig, "")
	}

	if item.URL == "" {
		return fmt.Errorf("tiktok item has no url")
	}
	imageFile, err := downloadImage(item.URL)
	if err != nil {
		return fmt.Errorf("download image: %w", err)
	}
	photoConfig := tgbotapi.NewPhoto(chatID, imageFile)
	photoConfig.ReplyToMessageID = messageID
	if _, err := rcv.bot.Send(photoConfig); err != nil {
		return fmt.Errorf("send photo: %w", err)
	}
	return nil
}

func getVideoFileBytes(filepath, name string) (photoFileBytes tgbotapi.FileBytes, err error) {
	defer func() {
		err = file.DeleteFile(filepath)
	}()
	photoBytes, err := os.ReadFile(filepath)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("read file err: %w", err)
	}
	photoFileBytes = tgbotapi.FileBytes{
		Name:  name,
		Bytes: photoBytes,
	}
	return photoFileBytes, nil
}
