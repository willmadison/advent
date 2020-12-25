package advent2020

import (
	"bufio"
	"io"
	"strconv"
)

type Encryptor struct {
	LoopNumber int
}

type Door struct {
	Encryptor
}

type RoomCard struct {
	Encryptor
}

func (e *Encryptor) DeriveLoopNumber(publicKey int) {
	subjectNumber := 7

	start := 1

	var iterations int

	for start != publicKey {
		start *= subjectNumber
		start %= 20201227
		iterations++
	}

	e.LoopNumber = iterations
}

func (e Encryptor) Encrypt(publicKey int) int {
	subjectNumber := publicKey

	start := 1

	var iterations int

	for iterations < e.LoopNumber {
		start *= subjectNumber
		start %= 20201227
		iterations++
	}

	return start
}

func DetermineEncryptionKey(r io.Reader) int {
	publicKeys := readPublicKeys(r)

	var door Door
	var card RoomCard

	door.DeriveLoopNumber(publicKeys[0])
	card.DeriveLoopNumber(publicKeys[1])

	return door.Encrypt(publicKeys[1])
}

func readPublicKeys(r io.Reader) []int {
	var keys []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		keys = append(keys, i)
	}

	return keys
}
