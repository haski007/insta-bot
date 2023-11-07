package listener

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func afterChar(s string, c rune) string {
	// Find the index of the character
	index := strings.IndexRune(s, c)
	if index == -1 {
		// Character not found
		return ""
	}
	// Return the substring after the character
	return s[index+1:]
}

func (rcv *InstaBotService) msgGPTextToSpeech(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	username := update.Message.From.UserName

	var voice string
	prompt := strings.TrimPrefix(update.Message.Text, "~")

	voice = regexp.MustCompile(`{(\w*)}`).FindString(prompt)
	voice = strings.Trim(voice, "{}")

	if voice != "" {
		prompt = afterChar(prompt, '}')
	}

	rcv.log.WithFields(map[string]interface{}{
		"from":   username,
		"prompt": prompt,
		"voice":  voice,
	}).Debugln("new request to text-to-speech")

	mp3Bytes, err := rcv.gpt.TextToSpeech(rcv.ctx, voice, prompt)
	if err != nil {
		rcv.log.WithError(err).Error("[msgGPTextToSpeech] text to speech")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGPTextToSpeech] text to speech: %s\n", err))
		return
	}

	os.WriteFile("audio.mp3", mp3Bytes, 0644)

	fileBytes := tgbotapi.FileBytes{
		Name:  "generated by " + update.Message.From.FirstName, // This is what the file will be called
		Bytes: mp3Bytes,
	}

	audioConfig := tgbotapi.NewAudio(chatID, fileBytes)

	rcv.log.WithFields(map[string]interface{}{
		"from":   username,
		"prompt": prompt,
	}).Info("new request to text-to-speech")

	if err := rcv.ReplyAudio(chatID, messageID, audioConfig); err != nil {
		rcv.log.WithError(err).Error("[msgGPTextToSpeech] send audio")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGPTextToSpeech] send audio: %s\n", err))
		return
	}
}
