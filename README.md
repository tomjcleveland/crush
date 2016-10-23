# Crush
Crush Go source files into a single line.

## Lol why
 * You want to obfuscate your shitty code.
 * You're trying to cheat at [code golf](http://codegolf.stackexchange.com/).
 * You're tired of those Go nerds [telling you how to live your life](https://blog.golang.org/go-fmt-your-code).

## Quickstart
```go
package main

import (
	"fmt"

	"github.com/tomjcleveland/crush"
)

const helloWorld = `package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}`

func main() {
	crushed, _ := crush.String(helloWorld)
	fmt.Println(crushed)
	// package main;import "fmt";func main() {fmt.Println("Hello, world!");};
}
```

## API
Yeah it's pretty complicated better check the [Godocs](https://godoc.org/github.com/tomjcleveland/crush).

## Aside
If you ever decided to compile a formal [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_Form) description using the [Go Spec](https://golang.org/ref/spec), you'd actually discover that pretty much *nothing* we would call Go conforms to this grammar, because most of us use newlines (`\n`) instead of semicolons (`;`).

The Go people sell this as a [feature](https://golang.org/ref/spec#Semicolons), which ok I guess it is, but then can't you *update the frigging EBNF?*

I'm working on a tool that validates code based on EBNF, but it doesn't work for Go unless I crush it first, so that's actually why I wrote this.