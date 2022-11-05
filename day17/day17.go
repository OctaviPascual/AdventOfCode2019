package day17

import (
	"fmt"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

type pixel rune

const (
	scaffold      pixel = '#'
	openSpace     pixel = '.'
	newLine       pixel = '\n'
	robotUp       pixel = '^'
	robotDown     pixel = 'v'
	robotLeft     pixel = '<'
	robotRight    pixel = '>'
	robotTumbling pixel = 'X'
)

type camera struct {
	output [][]pixel
}

type position struct {
	i, j int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		program: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	var view strings.Builder
	onOuput := func(output int) { view.WriteRune(rune(output)) }

	intcodeProgram, err := intcode.NewIntcodeProgram(d.program, intcode.MustNotInput, onOuput)
	if err != nil {
		return "", fmt.Errorf("could not create intcode program: %w", err)
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", fmt.Errorf("could not run intcode program: %w", err)
	}

	// remove last newLine from the view as it ends with \n\n
	camera := newCamera(view.String()[:view.Len()-1])
	return fmt.Sprintf("%d", camera.calibrate()), nil
}

func newCamera(view string) *camera {
	var output [][]pixel
	var row []pixel
	for _, c := range view {
		if pixel(c) == newLine {
			output = append(output, row)
			row = nil
			continue
		}
		row = append(row, pixel(c))
	}
	return &camera{
		output: output,
	}
}

func (c *camera) calibrate() int {
	sumAlignmentParameters := 0

	rows := len(c.output)
	cols := len(c.output[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			position := position{i, j}
			if c.isScaffoldIntersection(position) {
				sumAlignmentParameters += position.alignmentParameters()
			}
		}
	}

	return sumAlignmentParameters
}

func (c *camera) isScaffold(position position) bool {
	if position.i < 0 || position.i >= len(c.output) {
		return false
	}
	if position.j < 0 || position.j >= len(c.output[0]) {
		return false
	}
	return c.output[position.i][position.j] == scaffold
}

func (c *camera) isScaffoldIntersection(position position) bool {
	return c.isScaffold(position) &&
		c.isScaffold(position.up()) &&
		c.isScaffold(position.down()) &&
		c.isScaffold(position.left()) &&
		c.isScaffold(position.right())
}

func (p position) up() position {
	return position{p.i - 1, p.j}
}

func (p position) down() position {
	return position{p.i + 1, p.j}
}

func (p position) left() position {
	return position{p.i, p.j - 1}
}

func (p position) right() position {
	return position{p.i, p.j + 1}
}

func (p position) alignmentParameters() int {
	return p.i * p.j
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}
