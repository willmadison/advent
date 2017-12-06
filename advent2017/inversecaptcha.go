package advent2017

func InverseCaptcha(captcha string) int {
	var sum int
	runes := []rune(captcha)
	previous := runes[len(runes)-1]

	for _, digit := range runes {
		if previous == digit {
			sum += int(digit - '0')
		}
		previous = digit
	}

	return sum
}
