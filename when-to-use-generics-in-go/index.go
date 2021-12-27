package main

import (
	"fmt"
)

type Indexable[T any] interface {
	~[]T | ~string
}

func main() {
	
	slice := "apple"

	val := Get[uint8, string](0, slice)

	fmt.Println(val)

}

func Get[T any, E Indexable[T]] (idx int, items E) T {
	return items[idx]
}
