package main

import (
	"fmt"
	"html"
	"net/http"
)

type album struct {
	ID         string `json:"id"`
	ArtistName string `json:"artistName"`
	Song       string `json:"song"`
	numbers    string `josn:"number_of_songs"`
}

func main() {
	res := []album{album{"1", "joe", "wake me up", "first"}, album{"2", "wish", "sleepless nightn", "second"}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	http.HandleFunc("/getalbums", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%+ v\n", res))
	})
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s=%s\n", key, value)
		}
	})

	http.ListenAndServe(":8081", nil)
}
