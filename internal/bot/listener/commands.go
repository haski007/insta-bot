package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Keep in sync with the command switch in StartPool.
var publicCommands = []tgbotapi.BotCommand{
	{Command: "help", Description: "Розбудити бота"},
	{Command: "sum", Description: "Підсумок останніх N повідомлень: /sum <N> [питання]"},
	{Command: "purge_history", Description: "Очистити збережену історію чату"},

	{Command: "set_system_role", Description: "Задати системну роль GPT для чату: /set_system_role <роль>"},
	{Command: "drop_my_gpt", Description: "Скинути мою розмову з GPT"},
	{Command: "drop_my_grok", Description: "Скинути мою розмову з Grok"},

	{Command: "disable_loader", Description: "Вимкнути завантаження Instagram/TikTok у чаті"},
	{Command: "enable_loader", Description: "Увімкнути завантаження Instagram/TikTok у чаті"},

	{Command: "lets_play", Description: "Покликати гравців у CS:GO: /lets_play [HH:MM]"},
	{Command: "lets_play_pubg", Description: "Покликати гравців у PUBG"},
	{Command: "lets_play_finals", Description: "Покликати гравців у The Finals"},
	{Command: "list_players", Description: "Список зареєстрованих гравців"},
	{Command: "reg_csgo_players", Description: "Додати гравців CS:GO: /reg_csgo_players @user ..."},
	{Command: "purge_csgo_players", Description: "Очистити гравців CS:GO"},
	{Command: "reg_pubg_players", Description: "Додати гравців PUBG: /reg_pubg_players @user ..."},
	{Command: "purge_pubg_players", Description: "Очистити гравців PUBG"},
	{Command: "reg_finals_players", Description: "Додати гравців The Finals: /reg_finals_players @user ..."},
	{Command: "purge_finals_players", Description: "Очистити гравців The Finals"},
	{Command: "set_email", Description: "Прив'язати email до юзера: /set_email <username> <email>"},

	{Command: "sub_to_startup", Description: "Підписати чат на стартап-розсилку"},
	{Command: "unsub_to_startup", Description: "Відписати чат від стартап-розсилки"},
	{Command: "arc", Description: "Поточні події ARC Raiders"},
	{Command: "sub_arc_events", Description: "Підписати чат на події ARC Raiders"},
	{Command: "unsub_arc_events", Description: "Відписати чат від подій ARC Raiders"},
}

var creatorCommands = []tgbotapi.BotCommand{
	{Command: "ukraine_for_ukrainians", Description: "Увімкнути перевірку англіцизмів у чаті"},
	{Command: "unsub_ukraine_for_ukrainians", Description: "Вимкнути перевірку англіцизмів у чаті"},
	{Command: "ignore", Description: "Не перевіряти англіцизми юзера: /ignore @нікнейм"},
	{Command: "unignore", Description: "Знову перевіряти англіцизми юзера: /unignore @нікнейм"},
	{Command: "fuck", Description: "Ставити 💩 на повідомлення юзера: /fuck @username"},
	{Command: "unfuck", Description: "Прибрати юзера з 💩-списку: /unfuck @username"},
	{Command: "spam", Description: "Заспамити повідомленням: /spam <N> ..."},
	{Command: "w", Description: "Написати в інший чат від бота: /w <chat_id> [текст]"},
	{Command: "stream_chat", Description: "Пересилати повідомлення чату сюди: /stream_chat <chat_id>"},
	{Command: "stop_stream_chat", Description: "Зупинити пересилання: /stop_stream_chat <chat_id>"},
	{Command: "get_streams", Description: "Які чати пересилаються сюди"},
	{Command: "set_quality", Description: "Максимальна якість YouTube: /set_quality <N>"},
	{Command: "test", Description: "Тестовий інвойс"},
}

// RegisterCommands publishes the command hints shown in Telegram's "/" menu:
// public commands everywhere, plus creator-only ones in the creator's private chat.
func (rcv *InstaBotService) RegisterCommands() error {
	if _, err := rcv.bot.Request(tgbotapi.NewSetMyCommands(publicCommands...)); err != nil {
		return fmt.Errorf("set default commands: %w", err)
	}

	all := make([]tgbotapi.BotCommand, 0, len(publicCommands)+len(creatorCommands))
	all = append(all, publicCommands...)
	all = append(all, creatorCommands...)
	scope := tgbotapi.NewBotCommandScopeChat(rcv.creatorID)
	if _, err := rcv.bot.Request(tgbotapi.NewSetMyCommandsWithScope(scope, all...)); err != nil {
		return fmt.Errorf("set creator commands: %w", err)
	}
	return nil
}
