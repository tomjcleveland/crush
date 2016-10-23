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
