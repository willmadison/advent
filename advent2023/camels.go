package advent2023

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Card int

const (
	Joker Card = iota
	_
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var cardsByLabel = map[rune]Card{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

type Rule int

const (
	NoJokers Rule = iota
	JokersWild
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []Card
	Type  HandType
}

func newHand(cards []Card, rule Rule) Hand {
	hand := Hand{Cards: cards}

	var handType HandType
	countsByCard := map[Card]int{}

	for _, card := range cards {
		countsByCard[card]++
	}

	if rule == JokersWild {
		numJokers := countsByCard[Joker]
		delete(countsByCard, Joker)

		mostFrequentCard := cards[0]
		highestFrequency := countsByCard[mostFrequentCard]

		for card, count := range countsByCard {
			if count > highestFrequency {
				mostFrequentCard = card
			}
		}

		// Greedily increment the largest count of cards we have by the number of Jokers
		countsByCard[mostFrequentCard] += numJokers
	}

	switch {
	case len(countsByCard) == 1:
		handType = FiveOfAKind
	case len(countsByCard) == 2:
		for _, count := range countsByCard {
			if count == 4 || count == 1 {
				handType = FourOfAKind
				break
			} else if count == 3 || count == 2 {
				handType = FullHouse
				break
			}
		}
	case len(countsByCard) == 5:
		handType = HighCard
	default:
		var pairs int
		var handTypeSet bool

		for _, count := range countsByCard {
			if count == 3 {
				handType = ThreeOfAKind
				handTypeSet = true
				break
			} else if count == 2 {
				pairs++
			}
		}

		if !handTypeSet {
			switch pairs {
			case 1:
				handType = OnePair
			case 2:
				handType = TwoPair
			}
		}

	}

	hand.Type = handType

	return hand
}

type Wager struct {
	Hand Hand
	Bid  int
}

func SortCamelCardWagers(r io.Reader, rules ...Rule) []Wager {
	wagers := parseWagers(r, rules...)

	sort.Slice(wagers, func(i, j int) bool {
		if wagers[i].Hand.Type == wagers[j].Hand.Type {
			for card := range wagers[i].Hand.Cards {
				if wagers[i].Hand.Cards[card] != wagers[j].Hand.Cards[card] {
					return wagers[i].Hand.Cards[card] < wagers[j].Hand.Cards[card]
				}
			}
		} else {
			return wagers[i].Hand.Type < wagers[j].Hand.Type
		}

		return false
	})

	return wagers
}

func parseWagers(r io.Reader, rules ...Rule) []Wager {
	var wagers []Wager

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		wagers = append(wagers, parseWager(scanner.Text(), rules...))
	}

	return wagers
}

func parseWager(line string, rules ...Rule) Wager {
	wagerParts := strings.Fields(line)

	bid, _ := strconv.Atoi(wagerParts[1])

	cards := []Card{}

	var rule Rule

	if len(rules) == 0 {
		rule = NoJokers
	} else {
		rule = rules[0]
	}

	for _, label := range wagerParts[0] {
		var card Card

		if label == 'J' {
			switch rule {
			case JokersWild:
				card = Joker
			default:
				card = Jack
			}
		} else {
			card = cardsByLabel[label]
		}

		cards = append(cards, card)
	}

	return Wager{Hand: newHand(cards, rule), Bid: bid}
}
