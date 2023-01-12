package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) getChatMembers(chatID int64) ([]tgbotapi.ChatMember, error) {
	cnt, err := rcv.bot.GetChatMembersCount(tgbotapi.ChatMemberCountConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: chatID,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("get chat members count err: %w", err)
	}

	var members []tgbotapi.ChatMember
	for i := 0; i < cnt; i++ {
		m, err := rcv.bot.GetChatMember(tgbotapi.GetChatMemberConfig{ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             chatID,
			SuperGroupUsername: "",
			UserID:             int64(i),
		}})
		if err != nil {
			return nil, fmt.Errorf("get chat member #%d err: %s", i, err)
		}

		members = append(members, m)
	}

	return members, nil
}
