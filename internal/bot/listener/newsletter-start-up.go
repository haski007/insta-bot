package listener

import (
	"fmt"

	"github.com/haski007/insta-bot/pkg/emoji"

	"github.com/haski007/insta-bot/internal/bot/listener/transform"
	"github.com/haski007/insta-bot/internal/storage"
	"github.com/sashabaranov/go-openai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdSubToStartupHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.SubscribeChatToStartup(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdSubToStartupHandler] subscribe chat to startup")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[cmdSubToStartupHandler] subscribe chat to startup: %s\n", err))
		return
	}

	if err := rcv.SendMessage(chatID, "You have successfully subscribed to start-up newsletter"+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdSubToStartupHandler] send message")
		rcv.NotifyCreator(fmt.Sprintf("[cmdSubToStartupHandler] send message: %s\n", err))
		return
	}
}

func (rcv *InstaBotService) cmdUnsubToStartupHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.UnsubscribeChatToStartup(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdUnsubToStartupHandler] unsubscribe chat to startup")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[cmdUnsubToStartupHandler] unsubscribe chat to startup: %s\n", err))
		return
	}

	if err := rcv.SendMessage(chatID, "You have unsubscribed from start-up newsletter "+emoji.Basket); err != nil {
		rcv.log.WithError(err).Error("[cmdSubToStartupHandler] send message")
		rcv.NotifyCreator(fmt.Sprintf("[cmdSubToStartupHandler] send message: %s\n", err))
		return
	}
}

func (rcv *InstaBotService) sendStartupNewsletter() error {
	newsletters, err := rcv.storage.GetAllStartupNewsletters()
	if err != nil {
		return fmt.Errorf("get all startup newsletters: %s", err)
	}

	for chatID, newsletter := range newsletters {
		gptMessages := transform.ReplicasToGPTMessagesHistory(newsletter)

		gptMessages = append(gptMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: `Generate new startup ideas`,
		})

		rsp, err := rcv.gpt.Conversation(rcv.ctx, gptMessages)
		if err != nil {
			return fmt.Errorf("get conversation answer: %w", err)
		}

		replicasToSave := []storage.Replica{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: `Generate new startup ideas`,
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: rsp,
			},
		}

		message := fmt.Sprintf("Ну що ви котусики, тримайте нову порцію стартап ідей на сьогодні:\n\n%s", rsp)
		if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
			return fmt.Errorf("send message without markdown: %w", err)
		}

		if err := rcv.storage.PushStartupNewsletter(chatID, replicasToSave); err != nil {
			return fmt.Errorf("push startup newsletter: %w", err)
		}

		rcv.log.WithFields(map[string]interface{}{
			"chatID": chatID,
		}).Info("new start-up ideas newsletter")
	}
	return nil
}
