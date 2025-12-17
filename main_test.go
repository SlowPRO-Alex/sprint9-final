package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	// randSource источник псевдо случайных чисел.
	// Для повышения уникальности в качестве seed
	// используется текущее время в unix формате (в виде числа)
	randSource = rand.NewSource(time.Now().UnixNano())
	// randRange использует randSource для генерации случайных чисел
	randRange = rand.New(randSource)
)

func TestGenerateRandomElements(t *testing.T) {
	var size int
	s := generateRandomElements(size) // Отправляем 0
	assert.Equal(t, s, []int{})       // Проверяем, вернулся ли пустой слайс

	s = generateRandomElements(-1) // Отправляем отрицательное значение
	assert.Equal(t, s, []int{})    // Проверяем, вернулся ли пустой слайс
}

func TestMaximum(t *testing.T) {
	var s []int
	max := maximum(s)       // Отправляем пустой слайс
	assert.Equal(t, max, 0) // Проверяем, вернулся ли 0

	s = append(s, randRange.Intn(100))
	max = maximum(s)           // Отправляем слайс с 1 элементом
	assert.Equal(t, max, s[0]) // Проверяем, вернулся ли этот элемент

	s = append(s, randRange.Intn(100)*-1)
	max = maximum(s)           // Отправляем слайс с отрицательным числом
	assert.Equal(t, max, s[0]) // Проверяем, вернулось ли максимальное значение слайса

	s[1] = randRange.Intn(100)
	s[0] = s[1]
	s = append(s, s[0])
	max = maximum(s)           // Отправляем слайс с одинаковыми элементами
	assert.Equal(t, max, s[0]) // Проверяем, вернулось ли максимальное значение слайса

	s[0] = 100
	max = maximum(s)           // Отправляем слайс с максимальным элементом в начале
	assert.Equal(t, max, s[0]) // Проверяем, вернулось ли максимальное значение слайса

	s[len(s)-1] = 101
	max = maximum(s)                  // Отправляем слайс с максимальным элементом в конце
	assert.Equal(t, max, s[len(s)-1]) // Проверяем вернулось, ли максимальное значение слайса

	for i := 0; i < len(s); i++ {
		s[i] = s[i] * -1
	}
	max = maximum(s)           // Отправляем слайс только с отрицательными числами
	assert.Equal(t, max, s[1]) // Проверяем вернулось, ли максимальное значение слайса
}

func TestMaxChunks(t *testing.T) {
	var s []int
	max := maxChunks(s)     // Отправляем пустой слайс
	assert.Equal(t, max, 0) // Проверяем, вернулся ли 0

	s = make([]int, CHUNKS+1) // Создаем слайс с количеством элементов не делящихся без остатка на CHUNKS(8)
	for i := 0; i < len(s); i++ {
		s[i] = randRange.Intn(100) // Заполняем слайс
	}
	max = maxChunks(s)
	assert.NotEmpty(t, max) // Проверяем, что функция отработала и вернула значение

}
