package main

import "fmt"

func main() {

	var copiedString string
	dataString := "a"
	CopyString(&copiedString, dataString)
	fmt.Println(copiedString)

	var copiedInt int
	dataInt := 1
	CopyInt(&copiedInt, dataInt)
	fmt.Println(copiedInt)

	var copiedAStruct AStruct
	dataAStruct := AStruct{}
	CopyAStruct(&copiedAStruct, dataAStruct)
	fmt.Println(copiedAStruct)
}

// START OMIT
// CopyString takes the value from src and stores it in the destination
func CopyString(dst *string, src string) {
	*dst = src
}

// CopyInt takes the value from src and stores it in the destination
func CopyInt(dst *int, src int) {
	*dst = src
}

type AStruct struct{}

// CopyAStruct takes the value from src and store it in the destintation
func CopyAStruct(dst *AStruct, src AStruct) {
	*dst = src
}

// END OMIT
