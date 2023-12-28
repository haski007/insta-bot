package chatgpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_splitMessagesByTokens(t *testing.T) {
	tt := []struct {
		name     string
		messages []string
		limit    int
		want     [][]string
	}{
		{
			name: "success split 3 messages by 10 tokens limit",
			messages: []string{
				"1 2 3 4 5 6 7 8 9 10",
				"1 2 3 4 5 6 7 8 9 10",
				"1 2 3 4 5 6 7 8 9 10",
			},
			limit: 10,
			want: [][]string{
				{"1 2 3 4 5 6 7 8 9 10"},
				{"1 2 3 4 5 6 7 8 9 10"},
				{"1 2 3 4 5 6 7 8 9 10"},
			},
		},
		{
			name: "words as tokens",
			messages: []string{
				"hello world",
				"zalupa hui chlen piska",
				"taste this",
			},
			limit: 5,
			want: [][]string{
				{"hello world"},
				{"zalupa hui chlen piska"},
				{"taste this"},
			},
		},
		{
			name: "all fits in one limit",
			messages: []string{
				"hello world",
				"zalupa hui chlen piska",
				"taste this",
			},
			limit: 8,
			want: [][]string{
				{"hello world", "zalupa hui chlen piska", "taste this"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equalf(t, tc.want, splitMessagesByTokens(tc.messages, tc.limit), "splitMessagesByTokens(%v, %v)", tc.messages, tc.limit)
		})
	}
}
