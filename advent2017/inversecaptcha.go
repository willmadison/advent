package advent2017

func InverseCaptcha(captcha string) int {
	offset := len(captcha) / 2

	var sum int
	runes := []rune(captcha)

	for i, digit := range runes {
		if  digit == runes[(i+offset)%len(runes)] {
			sum += int(digit - '0')
		}
	}

	return sum
}
