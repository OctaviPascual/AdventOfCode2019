package day13

import (
	"fmt"

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

type action string

const (
	readX    action = "read_x"
	readY    action = "read_y"
	readTile action = "read_tile"
)

type arcadeCabinet struct {
	screen      map[pixel]tile
	nextAction  action
	x, y, score int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		program: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	arcadeCabinet := arcadeCabinet{
		screen:     make(map[pixel]tile),
		nextAction: readX,
	}

	intcodeProgram, err := intcode.NewIntcodeProgram(
		d.program, intcode.MustNotInput, arcadeCabinet.onOutput(),
	)
	if err != nil {
		return "", err
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", arcadeCabinet.countBlockTiles()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	arcadeCabinet := arcadeCabinet{
		screen:     make(map[pixel]tile),
		nextAction: readX,
	}

	intcodeProgram, err := intcode.NewIntcodeProgram(
		"2"+d.program[1:], arcadeCabinet.onInputWithQuarters(), arcadeCabinet.onOutputWithQuarters(),
	)
	if err != nil {
		return "", err
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", arcadeCabinet.score), nil
}

func (ac *arcadeCabinet) onOutput() func(output int) {
	return func(output int) {
		switch ac.nextAction {
		case readX:
			ac.x = output
			ac.nextAction = readY
		case readY:
			ac.y = output
			ac.nextAction = readTile
		case readTile:
			ac.screen[pixel{ac.x, ac.y}] = tile(output)
			ac.nextAction = readX
		}
	}
}

func (ac *arcadeCabinet) onInputWithQuarters() func() int {
	return func() int {
		ballPixel, err := ac.tilePosition(ball)
		if err != nil {
			return 0
		}

		horizontalPaddlePixel, err := ac.tilePosition(horizontalPaddle)
		if err != nil {
			return 0
		}

		if horizontalPaddlePixel.x < ballPixel.x {
			return 1
		}
		if horizontalPaddlePixel.x > ballPixel.x {
			return -1
		}

		return 0
	}
}

func (ac *arcadeCabinet) onOutputWithQuarters() func(output int) {
	return func(output int) {
		switch ac.nextAction {
		case readX:
			ac.x = output
			ac.nextAction = readY
		case readY:
			ac.y = output
			ac.nextAction = readTile
		case readTile:
			if ac.x == -1 && ac.y == 0 {
				ac.score = output
			} else {
				ac.screen[pixel{ac.x, ac.y}] = tile(output)
			}
			ac.nextAction = readX
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
