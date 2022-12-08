package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	// call the thing that we want to test
	got := Greeting()
	// specify what I should get back
	expected := "Hello, world!"
	// compare the result received with the expected value
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
