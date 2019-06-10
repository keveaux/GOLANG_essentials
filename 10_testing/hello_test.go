package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, worsld"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Hello() string {
	return "Hello, world"
}
