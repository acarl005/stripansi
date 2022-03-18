package stripansi

import (
	"fmt"
	"testing"
)

// go test -v stripansi_test.go stripansi.go
func ExampleStrip() {
	fmt.Println(Strip("\x1b[38;5;140m foo\x1b[0m bar"))
	// Output:
	//  foo bar
}

// go test -bench=.
func BenchmarkStrip(b *testing.B) {
	str := "\x1b[38;5;140m foo\x1b[0m bar"
	for i := 0; i < b.N; i++ {
		Strip(str)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/dreddsa5dies/stripansi
cpu: Intel(R) Core(TM) i7-3667U CPU @ 2.00GHz
BenchmarkStrip-4   	  509816	      2374 ns/op	      32 B/op	       3 allocs/op
PASS
ok  	github.com/dreddsa5dies/stripansi	1.249s
*/
