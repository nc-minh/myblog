package stringUtil

func ShortenStr(str string, maxLength ...int) string {
	max := 1000
	if len(maxLength) > 0 {
		max = maxLength[0]
	}

	if len(str) > max {
		str = str[:max]
	}
	return str
}
