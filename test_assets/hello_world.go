package main

import (
	"fmt"
	"strings"
)

// woah hey hello there

/* Here's a
   multiline
   comment
*/const hello = "hello" /* That's
   guaranteed to
   trip something
   up
*/const world = "world"

type punctuation struct {
	Bang  string `json:"bang"`
	comma string
}

var punct = punctuation{
	Bang:  "!",
	comma: ",",
}

func main() {
	fmt.Printf("%s%s %s%s\n", strings.Title(hello),
		punct.comma, world, punct.Bang)
}
