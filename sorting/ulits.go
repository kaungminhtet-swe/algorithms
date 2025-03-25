package sorting

import (
	"math/rand"
)

func IsSorted(instances []int) bool {
	for i := 1; i < len(instances); i++ {
		if instances[i-1] > instances[i] {
			return false
		}
	}
	return true
}

func GenerateRandomSlices(length uint) []int {
	result := make([]int, length)

	for i := range result {
		result[i] = rand.Intn(10000000)
	}

	return result
}
