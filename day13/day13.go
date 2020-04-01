package day13

import (
	"fmt"
	"time"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

type tile int

const (
	emtpy            tile = 0
	wall             tile = 1
	block            tile = 2
	horizontalPaddle tile = 3
	ball             tile = 4
)

type pixel struct {
	// x is the distance from the left, y the distance from the top
	x, y int
}

type arcadeCabinet struct {
	screen  map[pixel]tile
	program *intcode.Intcode
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		program: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
	if err != nil {
		return "", err
	}

	screen := make(map[pixel]tile)
	arcadeCabinet := arcadeCabinet{
		screen:  screen,
		program: intcodeProgram,
	}
	err = arcadeCabinet.run()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", arcadeCabinet.countBlockTiles()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram("2" + d.program[1:])
	if err != nil {
		return "", err
	}

	screen := make(map[pixel]tile)
	arcadeCabinet := arcadeCabinet{
		screen:  screen,
		program: intcodeProgram,
	}
	finalScore, err := arcadeCabinet.runWithQuarters()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", finalScore), nil
}

func (ac *arcadeCabinet) run() error {
	outputChannel := make(chan int)
	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- ac.program.Run(nil, outputChannel)
	}()

	for {
		x := <-outputChannel
		y := <-outputChannel
		tileID := <-outputChannel

		ac.screen[pixel{x, y}] = tile(tileID)

		select {
		case err := <-errorChannel:
			return err
		default:
		}
	}
}

func (ac *arcadeCabinet) runWithQuarters() (int, error) {
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- ac.program.Run(inputChannel, outputChannel)
	}()

	var score int
	for {
		select {
		case err := <-errorChannel:
			return score, err
		case x := <-outputChannel:
			y := <-outputChannel

			if x == -1 && y == 0 {
				score = <-outputChannel
			} else {
				tileID := <-outputChannel
				ac.screen[pixel{x, y}] = tile(tileID)
			}
		// Leave the program enough time to execute itself
		// If we can't read an output, we assume the program is waiting for an input
		// This is not ideal, but I couldn't come with a better way of knowing when the program is
		// expecting an input without changing Intcode API
		case <-time.After(200 * time.Microsecond):
			ballPixel, err := ac.tilePosition(ball)
			if err != nil {
				return 0, err
			}

			horizontalPaddlePixel, err := ac.tilePosition(horizontalPaddle)
			if err != nil {
				return 0, err
			}

			if horizontalPaddlePixel.x < ballPixel.x {
				inputChannel <- 1
			} else if horizontalPaddlePixel.x > ballPixel.x {
				inputChannel <- -1
			} else {
				inputChannel <- 0
			}
		}
	}
}

func (ac arcadeCabinet) countBlockTiles() int {
	total := 0
	for _, tile := range ac.screen {
		if tile == block {
			total++
		}
	}
	return total
}

func (ac arcadeCabinet) tilePosition(tileID tile) (pixel, error) {
	for pixel, tile := range ac.screen {
		if tile == tileID {
			return pixel, nil
		}
	}
	return pixel{}, fmt.Errorf("could not find the following tileID in the screen: %d", tileID)
}
