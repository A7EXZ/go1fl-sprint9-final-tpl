package main

import (
	"math"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"negative size", -10, 0},
		{"minimum int size", math.MinInt, 0},
		{"zero size", 0, 0},
		{"one element", 1, 1},
		{"small slice", 10, 10},
		{"large slice", 10_000, 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.wantLen {
				t.Fatalf("length = %d, want %d", len(data), tt.wantLen)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("data[%d] = %d, want a positive number", i, value)
				}
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
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"maximum at start", []int{9, 2, 3}, 9},
		{"maximum in middle", []int{2, 9, 3}, 9},
		{"maximum at end", []int{2, 3, 9}, 9},
		{"equal elements", []int{7, 7, 7}, 7},
		{"zeros", []int{0, 0, 0}, 0},
		{"negative elements", []int{-5, -2, -8}, -2},
		{"mixed signs", []int{-5, 0, 3}, 3},
		{"maximum int", []int{1, math.MaxInt, 2}, math.MaxInt},
		{"minimum int", []int{math.MinInt}, math.MinInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
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
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"equal elements", []int{7, 7, 7}, 7},
		{"zeros", []int{0, 0, 0}, 0},
		{"short negative slice", []int{-5, -2, -8}, -2},
		{"negative chunks", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"maximum int", []int{1, math.MaxInt, 2}, math.MaxInt},
		{"minimum int", []int{math.MinInt}, math.MinInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksBoundaries(t *testing.T) {
	for size := 1; size <= CHUNKS*2+1; size++ {
		data := make([]int, size)
		for i := range data {
			data[i] = 1
		}
		for position := range data {
			data[position] = 2
			if got := maxChunks(data); got != 2 {
				t.Fatalf("size %d, maximum at %d: got %d, want 2", size, position, got)
			}
			data[position] = 1
		}
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(100_003)
	want := maximum(data)
	if got := maxChunks(data); got != want {
		t.Errorf("maxChunks = %d, maximum = %d", got, want)
	}
}
