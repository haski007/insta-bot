package listener

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

// normalizeUkraineIgnoreUsername normalizes Telegram @username (latin, digits, underscore; 1–32 runes).
func normalizeUkraineIgnoreUsername(s string) (string, bool) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "@")
	s = strings.ToLower(s)
	if len(s) < 1 || len(s) > 32 {
		return "", false
	}
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '_' {
			return "", false
		}
	}
	return s, true
}

func (rcv *InstaBotService) cmdUkraineAnglicismIgnore(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		_ = rcv.SendMessageWithoutMarkdown(update.Message.Chat.ID, "Цю команду може використовувати лише власник бота."+emoji.NoEntry)
		return
	}
	chatID := update.Message.Chat.ID

	sub, err := rcv.storage.IsChatSubscribedToUkraineForUkrainians(chatID)
	if err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismIgnore] subscribed check")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Не вдалося перевірити підписку"+emoji.NoEntry)
		return
	}
	if !sub {
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Спочатку увімкни моніторинг: /ukraine_for_ukrainians"+emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendError(chatID, ErrNoArguments+" Використання: /ignore @нікнейм")
		return
	}
	user, ok := normalizeUkraineIgnoreUsername(args[0])
	if !ok {
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Некоректний нікнейм. Приклад: /ignore @username"+emoji.NoEntry)
		return
	}

	if err := rcv.storage.UkraineAnglicismIgnoreAdd(chatID, user); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismIgnore] add")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Не вдалося зберегти"+emoji.NoEntry)
		return
	}

	msg := fmt.Sprintf("Повідомлення від @%s не отримуватимуть відповіді щодо англіцизмів. Зняти: /unignore @%s", user, user)
	if err := rcv.SendMessageWithoutMarkdown(chatID, msg+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismIgnore] send")
		_ = rcv.NotifyCreator(fmt.Sprintf("[cmdUkraineAnglicismIgnore] send: %s\n", err))
	}
}

func (rcv *InstaBotService) cmdUkraineAnglicismUnignore(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		_ = rcv.SendMessageWithoutMarkdown(update.Message.Chat.ID, "Цю команду може використовувати лише власник бота."+emoji.NoEntry)
		return
	}
	chatID := update.Message.Chat.ID

	sub, err := rcv.storage.IsChatSubscribedToUkraineForUkrainians(chatID)
	if err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismUnignore] subscribed check")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Не вдалося перевірити підписку"+emoji.NoEntry)
		return
	}
	if !sub {
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Чат не підписаний на моніторинг"+emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendError(chatID, ErrNoArguments+" Використання: /unignore @нікнейм")
		return
	}
	user, ok := normalizeUkraineIgnoreUsername(args[0])
	if !ok {
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Некоректний нікнейм. Приклад: /unignore @username"+emoji.NoEntry)
		return
	}

	if err := rcv.storage.UkraineAnglicismIgnoreRemove(chatID, user); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismUnignore] remove")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "Не вдалося оновити список"+emoji.NoEntry)
		return
	}

	msg := fmt.Sprintf("Знято ігнор для @%s — знову перевірятимемо англіцизми.", user)
	if err := rcv.SendMessageWithoutMarkdown(chatID, msg+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineAnglicismUnignore] send")
		_ = rcv.NotifyCreator(fmt.Sprintf("[cmdUkraineAnglicismUnignore] send: %s\n", err))
	}
}
