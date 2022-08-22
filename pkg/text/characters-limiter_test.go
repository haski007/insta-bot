package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharLimiterToWord(t *testing.T) {
	tt := []struct {
		name   string
		limit  int
		text   string
		expect string
	}{
		{
			name:   "500 wrods & limit 512 ",
			limit:  512,
			text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam lectus libero, rutrum vehicula porta vitae, maximus sit amet dolor. Donec rutrum nisl quis libero sagittis venenatis. Nullam eget efficitur nulla. Donec sapien justo, tempus vitae convallis vel, laoreet sit amet mi. Pellentesque orci neque, maximus nec rutrum vel, cursus nec ante. Vestibulum sit amet auctor libero. Sed id condimentum mi. Nunc est massa, dictum eu euismod lobortis, lobortis quis nunc. Praesent gravida aliquam faucibus. Quisque molestie sit amet massa nec pretium. Integer auctor urna eget convallis accumsan. Aenean dignissim, nisl in viverra egestas, diam ex semper justo, vitae facilisis ex leo vel felis. Nulla volutpat ullamcorper turpis, vitae suscipit nisi. Nullam sed velit vel nulla efficitur tempor. Etiam in leo purus. Nunc sed turpis condimentum, feugiat quam a, molestie mi.\n\nAliquam erat volutpat. Donec efficitur metus ut dolor finibus venenatis. Maecenas at hendrerit felis. Donec sit amet orci ac metus vulputate fermentum. Ut nec ex vitae sem maximus tempor. Vivamus elementum metus id auctor pretium. Aenean dapibus, lacus at tincidunt molestie, turpis odio lobortis nulla, at mollis mi quam at metus. Mauris euismod non arcu et placerat. Morbi at magna tincidunt, scelerisque justo ut, ultricies lorem. Integer vel aliquet enim. Integer tristique enim neque, at auctor augue faucibus lacinia.\n\nUt mattis urna id eleifend suscipit. Donec fringilla sed tellus sit amet faucibus. Sed porttitor hendrerit orci, ut sagittis nunc commodo eget. Sed convallis ligula ut scelerisque interdum. Donec consequat arcu vel leo finibus finibus. Quisque ac consequat mauris. Nulla facilisi. Proin a tincidunt nunc, blandit tempus lectus. Aliquam vitae commodo ex. Aenean ac ante in elit luctus pellentesque vitae in mauris.\n\nDonec quis facilisis nibh. Vestibulum nec elit non lorem euismod ornare. In quis enim vel nulla iaculis gravida. In hac habitasse platea dictumst. Nunc risus sapien, efficitur vitae pharetra vel, eleifend vel enim. Etiam id lobortis magna, a imperdiet tortor. Sed placerat efficitur vehicula. Ut leo felis, euismod ut scelerisque at, scelerisque vel odio. Morbi rhoncus sollicitudin fermentum. Cras maximus porta fringilla. Aenean efficitur ultricies sodales. Etiam efficitur nunc in consequat elementum. Sed efficitur sit amet orci eu feugiat.\n\nNunc maximus feugiat enim, quis ullamcorper velit imperdiet et. Donec ac nisi ac mi varius convallis. Nullam auctor at elit eget malesuada. Vivamus eu neque id libero finibus bibendum vel quis justo. Aliquam viverra ante at fringilla pretium. Ut porttitor sodales sodales. Vestibulum sagittis accumsan odio et feugiat. Nam sed nisl tortor. Aliquam nibh massa, suscipit vel varius et, cursus vitae dolor. Etiam vulputate malesuada laoreet. Quisque ac velit sodales, convallis leo in, interdum dolor. Phasellus ac facilisis lorem, vitae commodo diam. Duis eget turpis interdum, pellentesque lorem non, vulputate massa. Pellentesque facilisis sodales neque, non sodales augue sodales non.\n\nCurabitur et sollicitudin magna. Vestibulum tempus, tortor et accumsan venenatis, est mauris fermentum nulla, sed sollicitudin tortor eros quis ligula. Integer et arcu eu lorem vehicula egestas id ut dui. Nulla est tortor, tempus et viverra sed, porta at diam. Integer nec dolor venenatis massa pharetra luctus. Aliquam nec dui.",
			expect: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam lectus libero, rutrum vehicula porta vitae, maximus sit amet dolor. Donec rutrum nisl quis libero sagittis venenatis. Nullam eget efficitur nulla. Donec sapien justo, tempus vitae convallis vel, laoreet sit amet mi. Pellentesque orci neque, maximus nec rutrum vel, cursus nec ante. Vestibulum sit amet auctor libero. Sed id condimentum mi. Nunc est massa, dictum eu euismod lobortis, lobortis quis nunc. Praesent gravida aliquam faucibus. Quisque...",
		},
		{
			name:   "10 chars & limit 1",
			limit:  1,
			text:   "awd xasdqw",
			expect: "a...",
		},
		{
			name:   "one word, chars: 5, limit: 5",
			limit:  5,
			text:   "dawdklawnlddmad",
			expect: "dawdk...",
		},
		{
			name:   "lower than limit",
			limit:  20,
			text:   "123456789",
			expect: "123456789",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, CharLimiterToWord(tc.text, tc.limit))
		})
	}
}
