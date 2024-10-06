package main

import (
	"fmt"
	"math"
	"math/rand"
)

const a = 5
const b = 6

func main() {
	array := make([]int, rand.Intn(37)+1)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1567) + 1

	}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d. Исходное число: %d, Корень: %f \n", i+1, array[i], hui(float64(array[i])))

	}

}

func hui(a float64) float64 {
	b := math.Pow(a, 0.5)

	return b
}
