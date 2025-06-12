package main

import (
	"testing"
)

func TestGenerateRandomElements_Empty(t *testing.T) {
	got := generateRandomElements(0)
	if got != nil {
		t.Error("Expected nil for size 0")
	}
}

func TestGenerateRandomElements_Size(t *testing.T) {
	size := 100
	got := generateRandomElements(size)
	if len(got) != size {
		t.Errorf("Expected length %d, got %d", size, len(got))
	}
}

func TestMaximum_Empty(t *testing.T) {
	if got := maximum([]int{}); got != 0 {
		t.Error("Expected 0 for empty slice")
	}
}

func TestMaximum_Single(t *testing.T) {
	if got := maximum([]int{5}); got != 5 {
		t.Error("Expected 5 for single element")
	}
}

func TestMaximum_Multiple(t *testing.T) {
	if got := maximum([]int{1, 3, 2, 5, 4}); got != 5 {
		t.Error("Expected 5 for [1,3,2,5,4]")
	}
}

func TestMaxChunks_Empty(t *testing.T) {
	if got := maxChunks([]int{}); got != 0 {
		t.Error("Expected 0 for empty slice")
	}
}

func TestMaxChunks_Single(t *testing.T) {
	if got := maxChunks([]int{5}); got != 5 {
		t.Error("Expected 5 for single element")
	}
}

func TestMaxChunks_Small(t *testing.T) {
	if got := maxChunks([]int{1, 2, 3, 4, 5}); got != 5 {
		t.Error("Expected 5 for [1,2,3,4,5]")
	}
}

func TestMaxChunks_Large(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i + 1
	}
	if got := maxChunks(data); got != 1000 {
		t.Error("Expected 1000 for large slice")
	}
}
