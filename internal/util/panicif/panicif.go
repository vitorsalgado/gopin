package panicif

// Err panics if argument err is not NIL
func Err(err error) {
	if err != nil {
		panic(err)
	}
}
