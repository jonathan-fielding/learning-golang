package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CatFacts struct {
	Status struct {
		Verified  bool `json:"verified"`
		SentCount int  `json:"sentCount"`
	} `json:"status"`
	Type      string    `json:"type"`
	Deleted   bool      `json:"deleted"`
	ID        string    `json:"_id"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	User      string    `json:"user"`
	Text      string    `json:"text"`
	Source    string    `json:"source"`
	V         int       `json:"__v"`
	Used      bool      `json:"used"`
}

func response(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/random?amount=1")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var catfacts CatFacts

	json.NewDecoder(resp.Body).Decode(&catfacts)

	fmt.Fprintf(w, catfacts.Text)
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":3000", nil)
}
