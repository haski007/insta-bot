package listener

import (
	"strings"

	"github.com/haski007/insta-bot/internal/storage"
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdDropGPTConversationHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID
	username := update.Message.From.UserName

	if err := rcv.storage.DropConversation(&storage.DropConversationReq{
		Username: username,
		UserID:   userID,
		ChatID:   chatID,
	}); err != nil {
		rcv.log.WithError(err).Error("[cmdDropGPTConversationHandler] drop conversation")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	if err := rcv.SendMessage(chatID, "Conversation dropped "+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdDropGPTConversationHandler] send message")
		return
	}
}

func (rcv *InstaBotService) cmdSetSystemRoleHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	role := strings.TrimSpace(update.Message.CommandArguments())

	if role == "" {
		rcv.SendError(chatID, ErrNoArguments)
		return
	}

	if err := rcv.storage.SetSystemRoleForChat(chatID, role); err != nil {
		rcv.log.WithError(err).Error("[cmdSetSystemRole] set system role")
		rcv.SendError(chatID, ErrInternalServerError)
		return
	}

	if err := rcv.SendMessage(chatID, "New system role for this chat was applied "+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdSetSystemRole] send message")
		return
	}
}
