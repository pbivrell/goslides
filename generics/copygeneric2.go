package main

import "fmt"

func main() {

// START OMIT
	var copiedString string
	dataString := "a"
	
	var copiedInt int
	dataInt := 1
	
	var copiedAStruct AStruct
	dataAStruct := AStruct{}
	
	Copy[string](&copiedString, dataString)

	Copy[int](&copiedInt, dataInt)

	Copy[AStruct](&copiedAStruct, dataAStruct)
// END OMIT
}

type AStruct struct{}

func Copy[T any](dst *T, src T) {
	*dst = src
}

