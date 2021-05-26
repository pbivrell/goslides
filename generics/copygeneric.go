package main

import "fmt"

func main() {

	var copiedString string
	dataString := "a"
	Copy(&copiedString, dataString)
	fmt.Println(copiedString)

	var copiedInt int
	dataInt := 1
	Copy(&copiedInt, dataInt)
	fmt.Println(copiedInt)

	var copiedAStruct AStruct
	dataAStruct := AStruct{}
	Copy(&copiedAStruct, dataAStruct)
	fmt.Println(copiedAStruct)
}

type AStruct struct{}

// START OMIT
func Copy[T any](dst *T, src T) {
	*dst = src
}

// END OMIT
