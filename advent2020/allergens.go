package advent2020

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

type FoodListing []Food

func (f FoodListing) FindNonAllergenicIngredients() ([]string, []string) {
	foodsByAllergen := map[string]map[int]struct{}{}

	allIngredients := map[string]struct{}{}

	for i, food := range f {
		for _, ingredient := range food.Ingredients {
			for _, allergen := range food.Allergens {
				if foodsByAllergen[allergen] == nil {
					foodsByAllergen[allergen] = map[int]struct{}{}
				}
				allIngredients[ingredient] = struct{}{}
				foodsByAllergen[allergen][i] = struct{}{}
			}
		}
	}

	potentiallyHarmfulIngredients := map[string]map[string]struct{}{}

	for allergen, recipies := range foodsByAllergen {
		ingredientsSeen := map[string]int{}

		for recipe := range recipies {
			for _, ingredient := range f[recipe].Ingredients {
				ingredientsSeen[ingredient]++
			}
		}

		for ingredient, numSeen := range ingredientsSeen {
			if numSeen == len(recipies) {
				if potentiallyHarmfulIngredients[allergen] == nil {
					potentiallyHarmfulIngredients[allergen] = map[string]struct{}{}
				}

				potentiallyHarmfulIngredients[allergen][ingredient] = struct{}{}
			}
		}
	}

	processed := map[string]struct{}{}

	for {
		removals := map[string][]string{}

		for allergen, ingredients := range potentiallyHarmfulIngredients {
			if _, matched := processed[allergen]; matched {
				continue
			}
			if len(ingredients) == 1 {
				processed[allergen] = struct{}{}

				for otherAllergen := range potentiallyHarmfulIngredients {
					if allergen == otherAllergen {
						continue
					}

					for ingredient := range ingredients {
						removals[ingredient] = append(removals[ingredient], otherAllergen)
					}
				}
			}
		}

		if len(removals) == 0 {
			break
		} else {
			for ingredient, allergens := range removals {
				for _, allergen := range allergens {
					delete(potentiallyHarmfulIngredients[allergen], ingredient)
				}
			}
		}
	}

	potentialAllergicIngredients := map[string]struct{}{}

	for _, ingredients := range potentiallyHarmfulIngredients {
		for ingredient := range ingredients {
			potentialAllergicIngredients[ingredient] = struct{}{}
		}
	}

	var nonAllergens []string

	for ingredient := range allIngredients {
		if _, present := potentialAllergicIngredients[ingredient]; !present {
			nonAllergens = append(nonAllergens, ingredient)
		}
	}

	var allergens []string

	for allergen := range potentiallyHarmfulIngredients {
		allergens = append(allergens, allergen)
	}

	sort.Strings(allergens)

	var allergicIngredients []string

	for _, allergen := range allergens {
		for ingredient := range potentiallyHarmfulIngredients[allergen] {
			allergicIngredients = append(allergicIngredients, ingredient)
		}
	}

	return nonAllergens, allergicIngredients
}

func (f FoodListing) CountOccurrencesFor(ingredients []string) map[string]int {
	occurrencesByIngredient := map[string]int{}

	for _, ingredient := range ingredients {
		occurrencesByIngredient[ingredient] = 0
	}

	for _, food := range f {
		for _, ingredient := range food.Ingredients {
			if _, present := occurrencesByIngredient[ingredient]; !present {
				continue
			}

			occurrencesByIngredient[ingredient]++
		}
	}

	return occurrencesByIngredient
}

type Food struct {
	Ingredients []string
	Allergens   []string
}

func ParseFoodListing(r io.Reader) FoodListing {
	var foods []Food

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		foods = append(foods, parseFood(scanner.Text()))
	}

	return FoodListing(foods)
}

func parseFood(value string) Food {
	food := Food{}

	foodParts := strings.Split(value, " (contains ")

	rawIngredientsList := foodParts[0]

	for _, ingredient := range strings.Split(rawIngredientsList, " ") {
		food.Ingredients = append(food.Ingredients, ingredient)
	}

	if len(foodParts) > 1 {
		rawAlergenList := strings.TrimRight(foodParts[1], ")")

		for _, allergen := range strings.Split(rawAlergenList, ", ") {
			food.Allergens = append(food.Allergens, allergen)
		}
	}

	return food
}
