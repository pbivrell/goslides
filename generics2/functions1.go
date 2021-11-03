package main

import (
	"fmt"
)

func main() {
	// START_CALLS1 OMIT
	Print[int](10)
	Print[string]("Apple")
	Print[float64](999.99)
	// END_CALLS1 OMIT

	// START_CALLS2 OMIT
	Print(1234)
	Print("pear")
	Print(23.45)
	Print([]int{1,2,3})
	// END_CALLS2 OMIT


}

// START_FUNC OMIT
func Print[T any](data T) {
	fmt.Println(data)
}
// END_FUNC OMIT
