package main

package main

import (
	"fmt"
)

func main() {
	reviews := []Review{
		{
			Title:  "Star Wars: Episode I - The Phantom Menace",
			Rating: 3.5,
		},{
			Title:  "Star Wars: Episode II - Attack of the Clones",
			Rating: 2.7,
		},{
			Title:  "Star Wars: Episode III - Revenge of the Sith",
			Rating: 6.4,
		},{
			Title:  "Star Wars: Episode IV - New Hope",
			Rating: 8.8,
		},{
			Title:  "Star Wars: Episode V - Empire Strikes Back",
			Rating: 9.2,
		},{
			Title:  "Star Wars: Episode VI - Return of the Jedi",
			Rating: 8.7,
		},
	}

	Sort(reviews)

	fmt.Println(reviews)
}

type Review struct {
	Title  string
	Rating float64
}

func (r Review) Less(o Review) bool {
	return r.Rating < o.Rating
}

// START OMIT
func Sort[Elem interface{ Less(Elem) bool }](list []Elem) {

	for i, v1 := range list {
		for j, v2 := range list {
			if v1.Less(v2) {
				list[i], list[j] = list[j], list[i]
			}
		}

	}
}
// END OMIT
