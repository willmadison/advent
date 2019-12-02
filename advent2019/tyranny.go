package advent2019

import (
	"bufio"
	"io"
	"strconv"
)

func DeriveTotalFuelRequirement(masses io.Reader) int {
	scanner := bufio.NewScanner(masses)

	var totalFuel int

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		totalFuel += DeriveFuelRequirementsByMass(mass)
	}

	return totalFuel
}

func DeriveTotalFuelRequirementIncludingFuelMass(masses io.Reader) int {
	scanner := bufio.NewScanner(masses)

	var totalFuel int

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuelNeeded := DeriveFuelRequirementsByMass(mass)

		totalFuel += fuelNeeded

		for fuelNeeded > 0 {
			fuelNeeded = DeriveFuelRequirementsByMass(fuelNeeded)

			if fuelNeeded > 0 {
				totalFuel += fuelNeeded
			}
		}
	}

	return totalFuel
}

func DeriveFuelRequirementsByMass(mass int) int {
	return mass/3 - 2
}
