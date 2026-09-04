package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := Greeting(); got == "" {
		t.Fatal("Greeting: got empty, want text")
	}
}
