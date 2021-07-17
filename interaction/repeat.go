package interaction

const repeatCount = 5

func Repeat(character string, times int) string {
	var repeated string

	if times < 1 {
		times = repeatCount
	}

	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
