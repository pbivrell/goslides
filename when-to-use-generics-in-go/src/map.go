package main

import (
	"strconv"
	"fmt"
)

// START_FUNC OMIT
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
// END_FUNC OMIT

// START_MAIN OMIT
func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i 
}

type Named struct {
	Name string
}

func main() {

	s := []string{"1","2","29480", "-2"}
	i := Map(s, toInt)
	fmt.Println(i)

	named := Map(s, func( x string) Named {
		return Named{
			Name: x,
		}
	})
	fmt.Println(named)
}
// END_MAIN OMIT
