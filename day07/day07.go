package day07

import (
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

type amplifier struct {
	id     rune
	phase  int
	input  chan int
	output chan int
}

type signalFn func(amplifiers []amplifier, program string) (int, error)

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
	phaseSettings := []int{0, 1, 2, 3, 4}
	signalFn := thrusterSignalInSeries
	maxThrusterSignal := getMaxThrusterSignal(d.program, phaseSettings, signalFn)
	return fmt.Sprintf("%d", maxThrusterSignal), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	phaseSettings := []int{5, 6, 7, 8, 9}
	signalFn := thrusterSignalWithFeedbackLoop
	maxThrusterSignal := getMaxThrusterSignal(d.program, phaseSettings, signalFn)
	return fmt.Sprintf("%d", maxThrusterSignal), nil
}

func getMaxThrusterSignal(program string, phaseSettings []int, signalFn signalFn) int {
	amplifiers := []amplifier{
		{id: 'A'},
		{id: 'B'},
		{id: 'C'},
		{id: 'D'},
		{id: 'E'},
	}
	allCombinations := generateAllPhaseSettingsCombinations(phaseSettings)

	maxThrusterSignal := 0
	for _, combination := range allCombinations {

		amplifiers[0].phase = combination[0]
		amplifiers[1].phase = combination[1]
		amplifiers[2].phase = combination[2]
		amplifiers[3].phase = combination[3]
		amplifiers[4].phase = combination[4]

		ts, _ := signalFn(amplifiers, program)
		if ts > maxThrusterSignal {
			maxThrusterSignal = ts
		}
	}
	return maxThrusterSignal
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

func thrusterSignalInSeries(amplifiers []amplifier, program string) (int, error) {
	wireInSeries(amplifiers)

	firstSignal := 0
	amplifiers[0].input <- firstSignal

	for _, amplifier := range amplifiers {
		onInput := func() int {
			return <-amplifier.input
		}

		onOutput := func(output int) {
			amplifier.output <- output
		}

		intcodeProgram, err := intcode.NewIntcodeProgram(program, onInput, onOutput)
		if err != nil {
			return 0, err
		}

		err = intcodeProgram.Run()
		if err != nil {
			return 0, err
		}
	}

	outputSignal := <-amplifiers[len(amplifiers)-1].output
	return outputSignal, nil
}

func wireInSeries(amplifiers []amplifier) {
	for i := range amplifiers {
		amplifiers[i].input = make(chan int, 2)
		amplifiers[i].input <- amplifiers[i].phase
	}

	for i := range amplifiers {
		if i < len(amplifiers)-1 {
			amplifiers[i].output = amplifiers[i+1].input
		} else {
			amplifiers[i].output = make(chan int, 1)
		}
	}
}

func thrusterSignalWithFeedbackLoop(amplifiers []amplifier, program string) (int, error) {
	wireWithFeedbackLoop(amplifiers)

	firstSignal := 0
	amplifiers[0].input <- firstSignal

	var group errgroup.Group
	for _, amplifier := range amplifiers {
		// create local variable for closures
		amplifier := amplifier

		onInput := func() int {
			return <-amplifier.input
		}

		onOutput := func(output int) {
			amplifier.output <- output
		}

		intcodeProgram, err := intcode.NewIntcodeProgram(program, onInput, onOutput)
		if err != nil {
			return 0, err
		}

		group.Go(func() error {
			return intcodeProgram.Run()
		})
	}
	err := group.Wait()
	if err != nil {
		return 0, err
	}

	outputSignal := <-amplifiers[len(amplifiers)-1].output
	return outputSignal, nil
}

func wireWithFeedbackLoop(amplifiers []amplifier) {
	for i := range amplifiers {
		amplifiers[i].input = make(chan int, 2)
		amplifiers[i].input <- amplifiers[i].phase
	}

	for i := range amplifiers {
		if i < len(amplifiers)-1 {
			amplifiers[i].output = amplifiers[i+1].input
		} else {
			amplifiers[i].output = amplifiers[0].input
		}
	}
}
