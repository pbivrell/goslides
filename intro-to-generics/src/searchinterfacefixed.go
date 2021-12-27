package main

import "fmt"

func main() {
	// START OMIT
	var myString string
	myString = Search("b", []interface{}{"a", "b", "c"}).(string)
	fmt.Println(myString)
	// END OMIT

}

func Search(find interface{}, list []interface{}) (item interface{}) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
