package stripansi

import "testing"

func TestStripSGRSubparameters(t *testing.T) {
	for _, tc := range []struct{ name, input, want string }{
		{"rgb", "\x1b[38:2::255:128:0mcolor\x1b[0m", "color"},
		{"background", "\x1b[48:2:0:0:128:255m背景\x1b[0m", "背景"},
		{"indexed", "\x1b[38:5:140mcolor\x1b[0m", "color"},
		{"combined", "\x1b[1;38:2::255:0:0mred\x1b[0m", "red"},
		{"underline", "\x1b[4:3mtext\x1b[0m", "text"},
		{"c1", "\u009b38:5:140mcolor\u009b0m", "color"},
		{"legacy", "\x1b[38;5;140m foo\x1b[0m bar", " foo bar"},
		{"plain", "12:30; status [38:2::255:0:0m]", "12:30; status [38:2::255:0:0m]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Strip(tc.input); got != tc.want {
				t.Errorf("Strip(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
