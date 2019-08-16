package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jerry")
	want := "Hello, Jerry"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}