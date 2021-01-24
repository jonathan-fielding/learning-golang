package main

import (
	"fmt"
	"os"
)

func main() {
	var hello = "hello"
	var name = "world"

	if len(os.Args[1:]) == 1 {
		name = os.Args[1]
	}

	fmt.Println(hello + " " + name)
}
