package stripansi

import "testing"

func TestStrip(t *testing.T) {
	tests := []struct {
		name, input, want string
	}{
		{"plain text", "hello 世界", "hello 世界"},
		{"color", "\x1b[38;5;140m foo\x1b[0m bar", " foo bar"},
		{"hyperlink with ST", "\x1b]8;;https://example.com\x1b\\example\x1b]8;;\x1b\\", "example"},
		{"hyperlink with BEL", "\x1b]8;;https://example.com\aexample\x1b]8;;\a", "example"},
		{"hyperlink parameters", "\x1b]8;id=docs;https://example.com/a?x=1\x1b\\read docs\x1b]8;;\x1b\\", "read docs"},
		{"colored hyperlink", "\x1b[32m\x1b]8;;https://example.com\x1b\\世界\x1b]8;;\x1b\\\x1b[0m", "世界"},
		{"window title", "before\x1b]0;A title: with / punctuation\aafter", "beforeafter"},
		{"multiple sequences", "\x1b]0;first\x1b\\one\x1b]0;second\a two", "one two"},
		{"C1 controls", "\u009d8;;https://example.com\u009clabel\u009d8;;\u009c", "label"},
		{"ordinary bracket text", "]8;;https://example.com label", "]8;;https://example.com label"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strip(tt.input); got != tt.want {
				t.Errorf("Strip(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
