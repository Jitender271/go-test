package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(2, 3)
	if result != 5 {
		t.Errorf("Expected -1 but got %d", result)
	}
}
