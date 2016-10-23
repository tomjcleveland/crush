package crush_test

import (
	"fmt"

	"github.com/tomjcleveland/crush"
)

func ExampleString() {
	helloWorld := `
        package main

        import "fmt"
        
        func main() {
            fmt.Println("Hello, world!")
        }`
	crushed, _ := crush.String(helloWorld)
	fmt.Println(crushed)
	// package main;import "fmt";func main() {fmt.Println("Hello, world!");};
	// Output: package main;import "fmt";func main() {fmt.Println("Hello, world!");};
}
