package main

import "testing"

func TestHello(t *testing.T) {
	// t.Fatal("not implemented")
	expected := "Hello, testing!"
	result := hello()

	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
