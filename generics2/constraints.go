package main

// START_STRINGER OMIT
type Stringer interface {
	String() string
}
// END_STRINGER OMIT

// START_WRITER OMIT
type Writer interface{
	Write(p []byte) (n int, err error)
}
// END_WRITER OMIT

// START_ONLYSIGNED OMIT
type OnlySignedInt interface {
	int | int8 | int16 | int32 | int64 
}
// END_ONLYSIGNED OMIT

// START_BETTERTYPE OMIT
type BetterInt int
// END_BETTERTYPE OMIT

// START_SIGNED OMIT
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 
}
// END_SIGNED OMIT

// START_ABSOLUTE OMIT
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 
	Positive() bool
}

type BetterInt int

func (b BetterInt) Positive() bool {
	return b > 0
}
// END_ABSOLUTE OMIT
