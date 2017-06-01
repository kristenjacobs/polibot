package main

import (
	"testing"
)

func TestGetResponseString(t *testing.T) {
	for i := 0; i < 10; i++ {
		result := getResponseString()
		if len(result) < 1 {
			t.Error("Failed to return a valid response")
		}
	}
}
