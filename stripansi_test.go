// go test -v stripansi_test.go stripansi.go
package stripansi

import (
	"fmt"
)

func Example_Strip() {
	fmt.Println(Strip("\x1b[38;5;140m foo\x1b[0m bar"))
	// Output:
	//  foo bar
}
