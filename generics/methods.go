package main

import (
	"fmt"
)

func main() {
	words := []fmt.Stringer{
		A{
			A: "I",
		},
		BetterString("have"),
		B{
			B:  1,
			BB: 350.28,
		},
		BetterString("dollars"),
	}
	fmt.Println(Sentence(words))
}

type BetterString string

func (b BetterString) String() string {
	return string(b)
}

type A struct {
	A string
}

func (a A) String() string {
	return a.A
}

type B struct {
	B  int
	BB float64
}

func (b B) String() string {
	return fmt.Sprintf("%d,%0.2f", b.B, b.BB)
}

// fmt.Stringer interfaces is defined as the following
// type Stringer interface {
//    String() string
//}

// Sentence creates a sentences from the provided values. Add space between the values and a
// period at the end. Assumes words is not empty
// START OMIT
func Sentence[T fmt.Stringer](words []T) string {
	sentence := words[0].String()
	for i, word := range words {
		if i == 0 {
			continue
		}
		sentence += " " + word.String()
	}
	return sentence + "."
}

// END OMIT
