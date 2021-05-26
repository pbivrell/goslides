package main

import "fmt"

func main() {
	// Max value should be 1023.39
	fmt.Println(FloatMax(2.3, 5.0, 1023.39, 1023.38999999999))

	// Max value should be 3298
	fmt.Println(IntMax(1, 2, 3, 4, 5, 7, 3298, -213))

}

func FloatMax(floats ...float64) float64 {
	var max float64
	for _, v := range floats {
		if v > max {
			max = v
		}
	}
	return max
}

func IntMax(ints ...int64) int64 {
	var max int64
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

// START OMIT
type Ordered interface {
	type int, float64
}

func Max[T Ordered](things ...T) T {
	var max T
	for _, v := range things {
		if v > max {
			max = v
		}
	}
	return max
}

// END OMIT
