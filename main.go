package main

import (
	"fmt"
	snip "github.com/dunstontc/snip/snippet"
)

func main() {
	file := "./ignore/ultisnips/c.snippets"
	matches := snip.ParseFile(file)
	for _, val := range matches {
		fmt.Printf("\n===\n%s", val)
	}
}
