package main
import (
	"bufio"
	"fmt"
	"html/template"
	"os"
)

func main() {
	consolereader := bufio.NewReader(os.Stdin)

	input, err := consolereader.ReadString('\n')

	if err == nil {
		fmt.Print(template.URLQueryEscaper(input))
	}
}
