package zengin

func LengthSJIS(str string) int {
	l := 0
	ss := []rune(str)
	for _, _ = range ss {
		l++
	}
	return l
}
