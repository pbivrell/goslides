package main

import (
	"fmt"
)

// START_STRINGER OMIT
type Stringer interface {
	String() string
}

func AddStrings[T Stringer](s1, s2 T) string {
	return s1.String() + s2.String()
}
// END_STRINGER OMIT

// START_ONLYSIGNED OMIT
type OnlySignedInt interface {
	int | int8 | int16 | int32 | int64 
}

func Celsius[T OnlySignedInt](F T) T {
	return (9/5) * (F - 32)
}
// END_ONLYSIGNED OMIT

// START_BETTERTYPE OMIT
type OtherInt int
// END_BETTERTYPE OMIT

// START_SIGNED OMIT
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 
}
// END_SIGNED OMIT

// START_ABSOLUTE OMIT
type PostiveSignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 
	Positive() bool
}

type BetterInt int

func (b BetterInt) Positive() bool {
	return b > 0
}
// END_ABSOLUTE OMIT

// START_MAIN OMIT
func main() {
	fmt.Println(Celsius(32))
	fmt.Println(Celsius(int8(32)))
}
// END_MAIN OMIT
