package main

import "testing"

func TestGreeting(t *testing.T) {
	result := Greeting()
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("Return %s, but wanted %s", result, expected)
	}
}
