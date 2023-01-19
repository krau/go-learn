package iteration

func Repeat(character string, repeatCount int) string {
	var repeated string
	// 迭代，使用 for
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
