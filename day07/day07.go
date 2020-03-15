package day07

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

type amplifier struct {
	id          rune
	firstInput  int
	secondInput int
}

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		program: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	amplifiers := []amplifier{
		{id: 'A'},
		{id: 'B'},
		{id: 'C'},
		{id: 'D'},
		{id: 'E'},
	}
	maxThrusterSignal := getMaxThrusterSignal(amplifiers, d.program)
	return fmt.Sprintf("%d", maxThrusterSignal), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func getMaxThrusterSignal(amplifiers []amplifier, program string) int {
	allCombinations := generateAllPhaseSettingsCombinations([]int{0, 1, 2, 3, 4})
	maxThrusterSignal := 0
	for _, combination := range allCombinations {

		amplifiers[0].firstInput = combination[0]
		amplifiers[1].firstInput = combination[1]
		amplifiers[2].firstInput = combination[2]
		amplifiers[3].firstInput = combination[3]
		amplifiers[4].firstInput = combination[4]

		ts, _ := thrusterSignal(amplifiers, program)
		if ts > maxThrusterSignal {
			maxThrusterSignal = ts
		}
	}
	return maxThrusterSignal
}

func thrusterSignal(amplifiers []amplifier, program string) (int, error) {
	secondInput := 0
	for _, amplifier := range amplifiers {

		amplifier.secondInput = secondInput

		intcodeProgram, err := intcode.NewIntcodeProgram(program)
		if err != nil {
			return 0, err
		}

		input := []int{amplifier.firstInput, amplifier.secondInput}
		output, err := intcodeProgram.RunWithInput(input...)
		if err != nil {
			return 0, err
		}
		secondInput = output[0]
	}
	return secondInput, nil
}

func generateAllPhaseSettingsCombinations(phaseSettings []int) [][]int {
	var phaseSettingsCombinations [][]int
	permute(phaseSettings, &phaseSettingsCombinations, 0)
	return phaseSettingsCombinations
}

func permute(phaseSettings []int, phaseSettingsCombinations *[][]int, i int) {
	if i > len(phaseSettings) {
		combination := make([]int, len(phaseSettings))
		copy(combination, phaseSettings)
		*phaseSettingsCombinations = append(*phaseSettingsCombinations, combination)
		return
	}

	permute(phaseSettings, phaseSettingsCombinations, i+1)

	for j := i + 1; j < len(phaseSettings); j++ {
		phaseSettings[i], phaseSettings[j] = phaseSettings[j], phaseSettings[i]
		permute(phaseSettings, phaseSettingsCombinations, i+1)
		phaseSettings[i], phaseSettings[j] = phaseSettings[j], phaseSettings[i]
	}
}
