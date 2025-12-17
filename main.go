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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size < 1 {
		return []int{}
	}
	s := make([]int, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		s[i] = rand.Intn(SIZE)
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup
	switch {
	case len(data) == 0:
		return 0
	case len(data)%CHUNKS != 0:
		{
			j := float64(CHUNKS) - float64(len(data)%CHUNKS)/(1/float64(CHUNKS))
			for j > 0 {
				data = append(data, 0)
				j--
			}
		}
	}
	result := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(ch []int) {
			// уменьшаем счётчик, когда горутина завершает работу
			defer wg.Done()
			result = append(result, maximum(data[len(data)/CHUNKS*i:len(data)/CHUNKS*i+len(data)/CHUNKS]))
		}(data)
	}
	wg.Wait()
	return maximum(result)
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	tStart := time.Now()
	max := maximum(data)
	elapsed := time.Since(tStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	tStart = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(tStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
