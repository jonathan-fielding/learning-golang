package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func desserts(w http.ResponseWriter, req *http.Request) {
	arr := [4]string{"Apple Pie", "Cherry Bakewell", "Sticky Toffee Pudding", "Rhubarb Crumble"}
	var random = rand.Intn(len(arr))

	fmt.Fprintf(w, arr[random])
}

func main() {
	http.HandleFunc("/desserts", desserts)
	http.ListenAndServe(":8000", nil)
}
