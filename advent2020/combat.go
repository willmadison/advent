package advent2020

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Card int

type Deck struct {
	Player int
	Queue  *cardQueue
}

func (d Deck) Score() int {
	var score int

	for d.Queue.Size() > 0 {
		topCard, _ := d.Queue.Peek()
		score += (d.Queue.Size() * int(topCard))
		d.Queue.Dequeue()
	}

	return score
}

func (d Deck) Serialize() string {
	return fmt.Sprintf("%v:%v", d.Player, d.Queue.data)
}

func (d *Deck) PlaceCards(winner, loser Card) {
	d.Queue.Enqueue(winner)
	d.Queue.Enqueue(loser)
}

func SerializeRound(firstPlayer, secondPlayer Deck) string {
	return fmt.Sprintf("%v && %v", firstPlayer.Serialize(), secondPlayer.Serialize())
}

func NewDeck(player int, cards []Card) Deck {
	return Deck{player, NewCardQueue(cards)}
}

type CardQueue interface {
	Enqueue(Card)
	Dequeue() (Card, error)
	Peek() (Card, error)
	Size() int
}

type cardQueue struct {
	data []Card
	size int
}

func NewCardQueue(cards []Card) *cardQueue {
	if len(cards) > 0 {
		return &cardQueue{data: cards, size: len(cards)}
	}

	return &cardQueue{data: []Card{}}
}

func (c *cardQueue) Enqueue(card Card) {
	c.data = append(c.data, card)
	c.size++
}

func (c *cardQueue) Dequeue() (Card, error) {
	if c.size > 0 {
		value := c.data[0]
		c.size--
		c.data = c.data[1:]

		return value, nil
	}

	return Card(math.MinInt64), errors.New("no such element")
}

func (c *cardQueue) Peek() (Card, error) {
	if c.size > 0 {
		value := c.data[0]
		return value, nil
	}

	return Card(math.MinInt64), errors.New("no such element")
}

func (c cardQueue) Size() int {
	return c.size
}

func ParseDecks(r io.Reader) (Deck, Deck) {
	player := 1

	var player1Deck, player2Deck Deck
	var cards []Card

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Player") {
			playerNumber := extractPlayerNumber(line)

			if playerNumber != player {
				player1Deck = NewDeck(player, cards)

				player = playerNumber
				cards = nil
			}

			continue
		}

		card, _ := strconv.Atoi(line)
		cards = append(cards, Card(card))
	}

	player2Deck = NewDeck(player, cards)

	return player1Deck, player2Deck
}

func extractPlayerNumber(value string) int {
	trimmed := strings.TrimPrefix(value, "Player ")
	trimmed = strings.TrimRight(trimmed, ":")

	number, _ := strconv.Atoi(trimmed)

	return number
}

func Battle(firstDeck, secondDeck *Deck) *Deck {
	for firstDeck.Queue.Size() > 0 && secondDeck.Queue.Size() > 0 {
		firstCard, _ := firstDeck.Queue.Dequeue()
		secondCard, _ := secondDeck.Queue.Dequeue()

		if firstCard > secondCard {
			firstDeck.PlaceCards(firstCard, secondCard)
		} else {
			secondDeck.PlaceCards(secondCard, firstCard)
		}
	}

	if firstDeck.Queue.Size() == 0 {
		return secondDeck
	}

	return firstDeck
}

func RecursiveBattle(firstDeck, secondDeck *Deck) *Deck {
	roundsPlayed := map[string]struct{}{}

	for firstDeck.Queue.Size() > 0 && secondDeck.Queue.Size() > 0 {
		round := SerializeRound(*firstDeck, *secondDeck)

		if _, played := roundsPlayed[round]; played {
			return firstDeck
		}

		roundsPlayed[round] = struct{}{}

		firstCard, _ := firstDeck.Queue.Dequeue()
		secondCard, _ := secondDeck.Queue.Dequeue()

		var winningPlayer int

		if firstDeck.Queue.Size() >= int(firstCard) && secondDeck.Queue.Size() >= int(secondCard) {
			firstDestination := make([]Card, int(firstCard))
			secondDestination := make([]Card, int(secondCard))

			copy(firstDestination, firstDeck.Queue.data[:int(firstCard)])
			copy(secondDestination, secondDeck.Queue.data[:int(secondCard)])

			firstCopy := NewDeck(firstDeck.Player, firstDestination)
			secondCopy := NewDeck(secondDeck.Player, secondDestination)

			winner := RecursiveBattle(&firstCopy, &secondCopy)

			winningPlayer = winner.Player
		} else {
			if firstCard > secondCard {
				winningPlayer = 1
			} else {
				winningPlayer = 2
			}
		}

		if winningPlayer == 1 {
			firstDeck.PlaceCards(firstCard, secondCard)
		} else {
			secondDeck.PlaceCards(secondCard, firstCard)
		}
	}

	if firstDeck.Queue.Size() == 0 {
		return secondDeck
	}

	return firstDeck
}
