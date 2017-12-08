package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/willmadison/advent/advent2017"
)

func main() {
	values := strings.Fields("5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6")
	var memory advent2017.Memory

	for _, value := range values {
		i, _ := strconv.Atoi(value)
		memory = append(memory, i)
	}

	_, cycleSize := advent2017.NumUniqueDistributions(memory)
	fmt.Println(cycleSize)
}
