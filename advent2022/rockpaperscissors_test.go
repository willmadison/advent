package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindFreshIngredients(t *testing.T) {
	given := strings.NewReader(`A Y
B X
C Z`)

	expectedTotalScore := 15

	totalScore, err := advent2022.FindTotalScoreForStrategyGuide(given, advent2022.GuideTypeMove)

	assert.Nil(t, err)

	assert.Equal(t, expectedTotalScore, totalScore)

	given = strings.NewReader(`A Y
B X
C Z`)

	expectedTotalScore = 12

	totalScore, err = advent2022.FindTotalScoreForStrategyGuide(given, advent2022.GuideTypeOutcome)

	assert.Nil(t, err)

	assert.Equal(t, expectedTotalScore, totalScore)
}
