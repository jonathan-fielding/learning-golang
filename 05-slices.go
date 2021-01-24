package main

import "fmt"

func main() {
	s := make([]string, 3)

	s[0] = "apple"
	s[1] = "banana"
	s[2] = "cherry"

	s = append(s, "dates")

	fmt.Println("slice:", s)
}
