package advent2025

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
	"github.com/willmadison/advent/internal/containers"
)

type InitializationInstruction struct {
	IndicatorLightsState       []bool
	IndicatorLightDesiredState []bool
	Buttons                    [][]int
	JoltageRequirements        []int
}

func MinimizeButtonPressesForMachineInitialization(r io.Reader) (int, error) {
	var instructions []InitializationInstruction

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		instruction, err := parseInitializationInstruction(line)
		if err != nil {
			return 0, err
		}

		instructions = append(instructions, instruction)
	}

	var totalPresses int

	for _, instruction := range instructions {
		buttonPresses := calculateMinButtonLightIndicatorPresses(instruction)

		totalPresses += buttonPresses
	}

	return totalPresses, nil
}

func MinimizeButtonPressesForProperMachineJoltage(r io.Reader) (int, error) {
	var instructions []InitializationInstruction

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		instruction, err := parseInitializationInstruction(line)
		if err != nil {
			return 0, err
		}

		instructions = append(instructions, instruction)
	}

	var totalPresses int

	for _, instruction := range instructions {
		buttonPresses := calculateMinButtonJoltagePresses(instruction)

		totalPresses += buttonPresses
	}

	return totalPresses, nil

}

func parseInitializationInstruction(line string) (InitializationInstruction, error) {
	instruction := InitializationInstruction{}

	parts := strings.Split(line, " ")
	rawIndicatorLights := parts[0]
	rawButtons := parts[1 : len(parts)-1]
	rawJoltageRequirements := parts[len(parts)-1]

	rawIndicatorLights = strings.Trim(rawIndicatorLights, "[]")
	indicatorLights := make([]bool, len(rawIndicatorLights))
	indicatorLightsInitialState := make([]bool, len(rawIndicatorLights))

	for i, v := range rawIndicatorLights {
		if v == '#' {
			indicatorLights[i] = true
		}
	}

	instruction.IndicatorLightsState = indicatorLightsInitialState
	instruction.IndicatorLightDesiredState = indicatorLights

	rawJoltageRequirements = strings.Trim(rawJoltageRequirements, "{}")

	var allButtons [][]int

	for _, rawButtom := range rawButtons {
		rawButtom = strings.Trim(rawButtom, "()")
		buttonStrings := strings.Split(rawButtom, ",")

		var buttons []int

		for _, buttonString := range buttonStrings {
			buttonValue, err := strconv.Atoi(buttonString)
			if err != nil {
				return InitializationInstruction{}, err
			}
			buttons = append(buttons, buttonValue)
		}

		allButtons = append(allButtons, buttons)
	}

	instruction.Buttons = allButtons

	joltageStrings := strings.Split(rawJoltageRequirements, ",")

	var joltageRequirements []int

	for _, v := range joltageStrings {
		joltage, err := strconv.Atoi(v)
		if err != nil {
			return InitializationInstruction{}, err
		}
		joltageRequirements = append(joltageRequirements, joltage)
	}

	instruction.JoltageRequirements = joltageRequirements

	return instruction, nil
}

func calculateMinButtonLightIndicatorPresses(instruction InitializationInstruction) int {
	numLights := len(instruction.IndicatorLightDesiredState)

	type state struct {
		lights  []bool
		presses int
	}

	stateToString := func(lights []bool) string {
		var sb strings.Builder
		for _, light := range lights {
			if light {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		return sb.String()
	}

	isDesiredState := func(lights []bool) bool {
		for i := 0; i < numLights; i++ {
			if lights[i] != instruction.IndicatorLightDesiredState[i] {
				return false
			}
		}
		return true
	}

	queue := containers.NewQueue[state]()
	queue.Enqueue(state{lights: make([]bool, numLights), presses: 0})
	visited := make(map[string]bool)
	visited[stateToString(instruction.IndicatorLightsState)] = true

	for queue.Size() > 0 {
		current, _ := queue.Dequeue()

		if isDesiredState(current.lights) {
			return current.presses
		}

		for _, button := range instruction.Buttons {
			newLights := make([]bool, numLights)
			copy(newLights, current.lights)

			for _, lightIndex := range button {
				newLights[lightIndex] = !newLights[lightIndex]
			}

			stateKey := stateToString(newLights)
			if !visited[stateKey] {
				visited[stateKey] = true
				queue.Enqueue(state{lights: newLights, presses: current.presses + 1})
			}
		}
	}

	return -1
}

func calculateMinButtonJoltagePresses(instruction InitializationInstruction) int {
	numButtons := len(instruction.Buttons)

	ctx := z3.NewContext(nil)
	solver := z3.NewSolver(ctx)

	intSort := ctx.IntSort()
	zero := ctx.FromInt(0, intSort).(z3.Int)
	one := ctx.FromInt(1, intSort).(z3.Int)

	buttons := make([]z3.Int, numButtons)
	for i := 0; i < numButtons; i++ {
		buttons[i] = ctx.IntConst("button_" + strconv.Itoa(i))
		solver.Assert(buttons[i].GE(zero))
	}

	for counterIdx, targetValue := range instruction.JoltageRequirements {
		var buttonsThatIncrement []z3.Int
		for buttonIdx, button := range instruction.Buttons {
			for _, affectedCounter := range button {
				if affectedCounter == counterIdx {
					buttonsThatIncrement = append(buttonsThatIncrement, buttons[buttonIdx])
					break
				}
			}
		}

		rhs := ctx.FromInt(int64(targetValue), intSort).(z3.Int)

		if len(buttonsThatIncrement) == 0 {
			if targetValue > 0 {
				solver.Assert(zero.Eq(one))
			}
		} else {
			sum := buttonsThatIncrement[0]
			for _, t := range buttonsThatIncrement[1:] {
				sum = sum.Add(t)
			}
			solver.Assert(sum.Eq(rhs))
		}
	}

	tot := ctx.IntConst("total")
	if len(buttons) > 0 {
		sumAll := buttons[0]
		for _, x := range buttons[1:] {
			sumAll = sumAll.Add(x)
		}
		solver.Assert(tot.Eq(sumAll))
	} else {
		solver.Assert(tot.Eq(zero))
	}

	minResult := -1
	for {
		sat, err := solver.Check()
		if !sat || err != nil {
			break
		}

		model := solver.Model()
		res := model.Eval(tot, true)
		val, _, _ := res.(z3.Int).AsInt64()

		minResult = int(val)

		cur := ctx.FromInt(val, intSort).(z3.Int)
		solver.Assert(tot.LT(cur))
	}

	return minResult
}
