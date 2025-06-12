package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Intn(size*10) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	var wg sync.WaitGroup
	maxes := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()

			start := chunkIndex * chunkSize
			if start >= len(data) {
				maxes[chunkIndex] = 0
				return
			}

			end := start + chunkSize
			if end > len(data) {
				end = len(data)
			}

			maxes[chunkIndex] = maximum(data[start:end])
		}(i)
	}

	wg.Wait()
	return maximum(maxes)
}

func main() {
	fmt.Printf("Generating %d random numbers...\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("\nFinding maximum in single thread...")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Max value: %d\nTime taken: %d µs\n", max, elapsed)

	fmt.Printf("\nFinding maximum using %d goroutines...\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Max value: %d\nTime taken: %d µs\n", max, elapsed)
}
