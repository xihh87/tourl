// Clean invalid characters on a UTF-8 file.
package main

import (
	"bufio"
	"io"
	"os"
	"unicode/utf8"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wt := io.Writer(os.Stdout)
	s := 0
	buf := make([]byte, utf8.UTFMax)

	for c, _, err := rd.ReadRune(); err == nil; c, _, err = rd.ReadRune() {
		s = utf8.EncodeRune(buf, c)
		if !utf8.ValidRune(c) {
			continue
		}
		wt.Write(buf[:s])
	}
}
