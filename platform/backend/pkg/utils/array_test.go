package utils_test

import (
	"backend/pkg/utils"
	"log"
	"testing"
)

func TestArrayIntersection(t *testing.T) {
	intArr1 := []int{-2, -1, 0, 1, 2, 3}
	intArr2 := []int{2, 3, 4, 5}
	log.Println(utils.ArrayIntersection(intArr1, intArr2))

	// one nil pass
	intArr1 = nil
	log.Println(utils.ArrayIntersection(intArr1, intArr2))

	// both nil pass
	intArr2 = nil
	log.Println(utils.ArrayIntersection(intArr1, intArr2))

	float64Arr1 := []float64{1.2, 1.5, 1.888, -1.77}
	float64Arr2 := []float64{0.99, 1.888, -1.77}
	log.Println(utils.ArrayIntersection(float64Arr1, float64Arr2))
}
