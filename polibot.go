package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	acceptableResponse := []string{
		"What we need is a full and frank discussion about the issues that really matter",
		"Lets be clear, nothing has changed in our manifesto, merely a clarification of the details",
		"What this country deserves is a strong and stable leadership",
	}
	fmt.Fprintf(w, "%s\n", acceptableResponse[rand.Intn(len(acceptableResponse))])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
