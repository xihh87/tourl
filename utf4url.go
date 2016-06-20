// This program writes a stream from stdin to stdout, encoded as a URL.
package main

import (
	"bufio"
	"html/template"
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
		if c < utf8.RuneSelf {
			wt.Write(buf[:s])
		} else {
			wt.Write([]byte(template.URLQueryEscaper(string(buf[:s]))))
		}
	}
}
