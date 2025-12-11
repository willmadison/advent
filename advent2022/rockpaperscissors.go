package advent2022

import (
	"bufio"
	"io"
	"strings"
)

type GuideType int

const (
	GuideTypeMove GuideType = iota
	GuideTypeOutcome
)

type OpponentMove rune

var (
	OpponentRock     OpponentMove = 'A'
	OpponentPaper    OpponentMove = 'B'
	OpponentScissors OpponentMove = 'C'
)

type YourMove rune

var (
	Rock     YourMove = 'X'
	Paper    YourMove = 'Y'
	Scissors YourMove = 'Z'
)

type Outcome rune

var (
	Loss Outcome = 'X'
	Draw Outcome = 'Y'
	Win  Outcome = 'Z'
)

type round struct {
	Opponent OpponentMove
	You      YourMove
}

type roundOutcome struct {
	Opponent OpponentMove
	Outcome  Outcome
}

func (r roundOutcome) determineYourMove() round {
	var round round

	round.Opponent = r.Opponent

	switch r.Opponent {
	case OpponentRock:
		switch r.Outcome {
		case Draw:
			round.You = Rock
		case Win:
			round.You = Paper
		case Loss:
			round.You = Scissors
		}
	case OpponentPaper:
		switch r.Outcome {
		case Draw:
			round.You = Paper
		case Win:
			round.You = Scissors
		case Loss:
			round.You = Rock
		}
	case OpponentScissors:
		switch r.Outcome {
		case Draw:
			round.You = Scissors
		case Win:
			round.You = Rock
		case Loss:
			round.You = Paper
		}
	}

	return round
}

func (r round) score() int {
	var score int

	switch r.Opponent {
	case OpponentRock:
		switch r.You {
		case Rock:
			score += 1
			score += 3
		case Paper:
			score += 2
			score += 6
		case Scissors:
			score += 3
		}
	case OpponentPaper:
		switch r.You {
		case Rock:
			score += 1
		case Paper:
			score += 2
			score += 3
		case Scissors:
			score += 3
			score += 6
		}
	case OpponentScissors:
		switch r.You {
		case Rock:
			score += 1
			score += 6
		case Paper:
			score += 2
		case Scissors:
			score += 3
			score += 3
		}
	}

	return score
}

func FindTotalScoreForStrategyGuide(r io.Reader, guideType GuideType) (int, error) {
	scanner := bufio.NewScanner(r)

	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()

		var round round
		var err error

		if guideType == GuideTypeMove {
			round, err = parseRoundFromStrategyGuideLine(line)
			if err != nil {
				return 0, err
			}
		} else {
			var roundOutcome roundOutcome
			roundOutcome, err = parseRoundOutcomeFromStrategyGuideLine(line)
			if err != nil {
				return 0, err
			}
			round = roundOutcome.determineYourMove()
		}

		totalScore += round.score()
	}

	return totalScore, nil
}

func parseRoundFromStrategyGuideLine(line string) (round, error) {
	parts := strings.Split(line, " ")

	return round{Opponent: OpponentMove(parts[0][0]), You: YourMove(parts[1][0])}, nil
}

func parseRoundOutcomeFromStrategyGuideLine(line string) (roundOutcome, error) {
	parts := strings.Split(line, " ")

	return roundOutcome{Opponent: OpponentMove(parts[0][0]), Outcome: Outcome(parts[1][0])}, nil
}
