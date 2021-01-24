package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

var pageCache = cache.New(5*time.Minute, 10*time.Minute)

func loadSite(w http.ResponseWriter, req *http.Request) {
	var url = "https://www.jonathanfielding.com" + req.URL.Path
	cachedResponse, found := pageCache.Get(req.URL.Path)

	fmt.Print(found)

	if found {
		fmt.Fprintf(w, cachedResponse.(string))
	} else {
		resp, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		pageCache.Set(req.URL.Path, bodyString, cache.DefaultExpiration)

		fmt.Fprintf(w, bodyString)
	}

}

func main() {
	http.HandleFunc("/", loadSite)
	http.ListenAndServe(":8000", nil)
}
