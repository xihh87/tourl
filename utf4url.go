// This program writes a stream from stdin to stdout, encoded as a URL.
package main

import (
	"bufio"
	"html/template"
	"io"
	"os"
	"unicode/utf8"
)

// The UTF constants from https://golang.org/pkg/unicode/utf8/
const (
	RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
	RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
	MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
	UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wt := io.Writer(os.Stdout)
	s := 0
	buf := make([]byte, UTFMax)

	for c, _, err := rd.ReadRune(); err == nil; c, _, err = rd.ReadRune() {
		s = utf8.EncodeRune(buf, c)
		if !utf8.ValidRune(c) {
			continue
		}
		if c < RuneSelf {
			wt.Write(buf[:s])
		} else {
			wt.Write([]byte(template.URLQueryEscaper(string(buf[:s]))))
		}
	}
}
