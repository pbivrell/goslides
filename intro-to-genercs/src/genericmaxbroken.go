package main

import "fmt"

func main() {
	// Max value should be 1023.39
	fmt.Println(Max(2.3, 5.0, 1023.39, 1023.38999999999))

	// Max value should be 3298
	fmt.Println(Max(1, 2, 3, 4, 5, 7, 3298, -213))

}

// START OMIT
func Max[T ?](things ...T) T {
	var max T
	for _, v := range things  {
		if v > max {
			max = v
		}
	}
	return max
}
// END OMIT
