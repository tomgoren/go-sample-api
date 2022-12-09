package main

import "testing"

func TestContentTypeValidation(t *testing.T) {
	valid, _ := validateContentType("application/json")
	if !valid {
		t.Fatal("Something is wrong")
	}
}
