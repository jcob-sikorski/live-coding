// The Growing Slice
// Write a program that appends elements to a slice inside a loop.
// Use testing.B (benchmarking) to see how many allocations happen
// when you don't pre-allocate with make([]int, 0, capacity).

// Goal: Understand how dynamic resizing triggers heap allocations.

package main

import "testing"

// AppendWithoutPrealloc appends n elements to a slice without initial capacity.
func AppendWithoutPrealloc(n int) {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
}

// AppendWithPrealloc appends n elements to a slice with pre-allocated capacity.
func AppendWithPrealloc(n int) {
	// We pre-allocate the capacity so no resizing is needed.
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
}

// BenchmarkWithoutPrealloc tests performance when the slice has to grow dynamically.
func BenchmarkWithoutPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithoutPrealloc(1000)
	}
}

// BenchmarkWithPrealloc tests performance when capacity is known upfront.
func BenchmarkWithPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithPrealloc(1000)
	}
}
