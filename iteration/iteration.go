package iteration

func Repeat(str string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return
}
