package main

import "testing"

func TestMain(t *testing.T) {
	got := Hello("Ram")
	want := "Hello, Ram!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
