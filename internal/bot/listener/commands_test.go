package listener

import (
	"os"
	"regexp"
	"testing"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/require"
)

func TestCommandHintsMatchReceiver(t *testing.T) {
	src, err := os.ReadFile("receiver.go")
	require.NoError(t, err)

	handled := map[string]bool{}
	for _, m := range regexp.MustCompile(`case "([a-z0-9_]+)":`).FindAllStringSubmatch(string(src), -1) {
		handled[m[1]] = true
	}
	handled["w"] = true
	require.NotEmpty(t, handled)

	hinted := map[string]bool{}
	for _, c := range append(append([]tgbotapi.BotCommand{}, publicCommands...), creatorCommands...) {
		require.Falsef(t, hinted[c.Command], "duplicate hint %q", c.Command)
		hinted[c.Command] = true
		require.Regexp(t, `^[a-z0-9_]{1,32}$`, c.Command)
		require.LessOrEqualf(t, utf8.RuneCountInString(c.Description), 256, "description of %q is too long", c.Command)
		require.Truef(t, handled[c.Command], "hint %q has no handler in receiver.go", c.Command)
	}
	for name := range handled {
		require.Truef(t, hinted[name], "command %q is handled but has no hint", name)
	}
}
