package main

import "fmt"

func main() {

// START OMIT
	var copiedString string
	Copy(&copiedString, 2)
// END OMIT
}

func Copy[T any](dst *T, src T) {
	*dst = src
}

