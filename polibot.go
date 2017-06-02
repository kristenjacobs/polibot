package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func getResponseString() string {
	acceptableResponse := []string{
		"...Given this is an ongoing investigation, I am unable comment on specific details. However, what I can say is that, as a party, we will do everything in our power to ensure that this issue get resolved in both a quick and decisive manner.",
		"What we need is a full and frank discussion about the issues that really matter",
		"??????? Lets be clear, nothing has changed in our manifesto, merely a clarification of the details",
		"What this country deserves is a strong and stable leadership",
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
