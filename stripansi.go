package stripansi

import (
	"regexp"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

const osc = "(?:\x1b\\]|\u009d)[^\x07\x1b\u009c]*(?:\x07|\x1b\\\\|\u009c)"

var re = regexp.MustCompile(osc + "|" + ansi)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}
