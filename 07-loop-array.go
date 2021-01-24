package main

import "fmt"

func main() {
	arr := [5]string{"apple", "banana", "cherry", "dates", "elderberry"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}
