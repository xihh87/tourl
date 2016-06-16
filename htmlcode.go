package main
import (
	"fmt"
	"html/template"
)

func main() {
	fmt.Println(template.URLQueryEscaper("ë"))
}
