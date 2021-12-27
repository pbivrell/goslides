package main


import (
	"fmt"
)

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	Positive() bool
}

type Paul int

func (p Paul) Positive() bool {
	return p > 0
}

func Apple[T SignedInt](x T) {
	if x < 0  {
		fmt.Println(x)
	}
	fmt.Println(x.Positive())
}

func main() {
	
	x := Paul(1)

	Apple(x)

}
