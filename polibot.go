package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func getResponseString() string {
	acceptableResponse := []string{
		"-> Given this is an ongoing investigation, I am not about to comment on specifics.",
		"-> What we need is a full and frank discussion about the issues that really matter",
		"-> Lets be absolutely clear, nothing has changed in our manifesto",
		"-> What this country deserves is a strong and stable leadership",
	}
	return acceptableResponse[rand.Intn(len(acceptableResponse))]
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", getResponseString())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
