package advent2018

import (
	"bytes"
	"strings"
	"unicode"
)

func OptimalReduction(polymer string) string {
	reductionByRune := map[rune]string{}

	for i := 'a'; i <= 'z'; i++ {
		upper := unicode.ToUpper(i)
		modified := strings.Replace(polymer, string(i), "", -1)
		modified = strings.Replace(modified, string(upper), "", -1)

		reductionByRune[i] = Reduce(modified)
	}

	minReduction := reductionByRune['a']

	for i := 'b'; i <= 'z'; i++ {
		if len(reductionByRune[i]) < len(minReduction) {
			minReduction = reductionByRune[i]
		}
	}

	return minReduction
}

func Reduce(polymer string) string {
	current, reduced := polymer, doReduce(polymer)

	for current != reduced {
		current, reduced = reduced, doReduce(reduced)
	}

	return reduced
}

func doReduce(polymer string) string {
	if polymer == "" {
		return polymer
	}

	chars := strings.Split(polymer, "")
	var buf bytes.Buffer

	var numReactions int

	for i := 0; i < len(chars)-1; {
		if !areReactive(chars[i], chars[i+1]) {
			buf.WriteString(chars[i])
			i++
		} else {
			i += 2
			numReactions++
		}

		if i == len(chars)-1 {
			buf.WriteString(chars[i])
		}
	}

	if buf.Len() == 0 && numReactions == 0 {
		return polymer
	}

	return buf.String()
}

func areReactive(a, b string) bool {
	return a != b && strings.ToLower(a) == strings.ToLower(b)
}
