package iteration

func Repeat(character string, repeatCounter uint8) string {
	var repeated string

	for i := uint8(0); i < repeatCounter; i++ {
		repeated += character
	}

	return repeated
}
