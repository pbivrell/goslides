package main

func main() {
	Convert(int(5), int8(5))
}


// START OMIT
type integer interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr
}

func Convert[To, From integer](from From) To {
	to := To(from)
	if From(to) != from {
		panic("conversion out of range")
	}
	return to
}
// END OMIT
