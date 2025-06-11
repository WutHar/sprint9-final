package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero size", 0},
		{"small size", 10},
		{"large size", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if tt.size == 0 && got != nil {
				t.Errorf("generateRandomElements(0) = %v, want nil", got)
			}
			if tt.size > 0 && len(got) != tt.size {
				t.Errorf("len(generateRandomElements(%d)) = %d, want %d", tt.size, len(got), tt.size)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"multiple elements", []int{1, 3, 2, 5, 4}, 5},
		{"all equal", []int{2, 2, 2}, 2},
		{"negative numbers", []int{-1, -3, -2}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"less than chunks", []int{1, 2, 3, 4, 5}, 5},
		{"exact chunks size", make([]int, 8), 0},
		{"multiple chunks", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if tt.name == "exact chunks size" {
				if got != 0 {
					t.Errorf("maxChunks() = %v, want 0 for zero-initialized slice", got)
				}
			} else if got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
