package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFoodListing(t *testing.T) {
	given := strings.NewReader(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)

	expected := FoodListing([]Food{
		{
			[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
			[]string{"dairy", "fish"},
		},
		{
			[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
			[]string{"dairy"},
		},
		{
			[]string{"sqjhc", "fvjkl"},
			[]string{"soy"},
		},
		{
			[]string{"sqjhc", "mxmxvkd", "sbzzf"},
			[]string{"fish"},
		},
	})

	actual := ParseFoodListing(given)
	assert.Equal(t, expected, actual)
}

func TestFindNonAllergenicIngredients(t *testing.T) {
	given := FoodListing([]Food{
		{
			[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
			[]string{"dairy", "fish"},
		},
		{
			[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
			[]string{"dairy"},
		},
		{
			[]string{"sqjhc", "fvjkl"},
			[]string{"soy"},
		},
		{
			[]string{"sqjhc", "mxmxvkd", "sbzzf"},
			[]string{"fish"},
		},
	})

	expectedNonAllergens := []string{"kfcds", "nhms", "sbzzf", "trh"}
	expectedAllergens := []string{"mxmxvkd", "sqjhc", "fvjkl"}

	actualNonAllergens, actualAllergens := given.FindNonAllergenicIngredients()

	assert.ElementsMatch(t, expectedNonAllergens, actualNonAllergens)
	assert.Equal(t, expectedAllergens, actualAllergens)

	expectedOccurrenceCount := map[string]int{
		"kfcds": 1,
		"nhms":  1,
		"sbzzf": 2,
		"trh":   1,
	}

	actualOccurrenceCount := given.CountOccurrencesFor(expectedNonAllergens)

	assert.Equal(t, expectedOccurrenceCount, actualOccurrenceCount)
}
