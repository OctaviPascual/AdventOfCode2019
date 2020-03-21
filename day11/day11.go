package day11

import (
	"fmt"
	"html"
	"math"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

type color int
type direction int

type orientation string

func (o orientation) turnLeft() orientation {
	switch o {
	case up:
		return left
	case down:
		return right
	case left:
		return down
	case right:
		return up
	}
	panic(fmt.Sprintf("unknown orientation: %s", o))
}

func (o orientation) turnRight() orientation {
	switch o {
	case up:
		return right
	case down:
		return left
	case left:
		return up
	case right:
		return down
	}
	panic(fmt.Sprintf("unknown orientation: %s", o))
}

const (
	black color = 0
	white color = 1

	turnLeft  direction = 0
	turnRight direction = 1

	up    orientation = "up"
	down  orientation = "down"
	left  orientation = "left"
	right orientation = "right"
)

var (
	blackEmoji = html.UnescapeString("&#" + strconv.Itoa(11035) + ";")
	whiteEmoji = html.UnescapeString("&#" + strconv.Itoa(11036) + ";")
)

type panel struct {
	x, y int
}

type grid map[panel]color

type robot struct {
	program     *intcode.Intcode
	panel       panel
	orientation orientation
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

	grid := grid(make(map[panel]color))
	robot := robot{
		program:     intcodeProgram,
		panel:       panel{0, 0},
		orientation: up,
	}
	robot.run(grid)

	return fmt.Sprintf("%d", grid.numberOfPaintedPanels()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
	if err != nil {
		return "", err
	}

	grid := grid(make(map[panel]color))
	grid[panel{0, 0}] = white
	robot := robot{
		program:     intcodeProgram,
		panel:       panel{0, 0},
		orientation: up,
	}
	robot.run(grid)

	topLeft := grid.topLeftPanel()
	bottomRight := grid.bottomRightPanel()

	return grid.renderHull(topLeft, bottomRight), nil
}

func (r *robot) run(grid grid) {
	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 2)
	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- r.program.Run(inputChannel, outputChannel)
	}()

	for {
		inputChannel <- int(grid[r.panel])

		colorToPaint := <-outputChannel
		r.paintPanel(grid, color(colorToPaint))

		directionToTurn := <-outputChannel
		r.turn(direction(directionToTurn))

		r.moveOneStepForward()

		select {
		case <-errorChannel:
			return
		default:
		}
	}
}

func (r robot) paintPanel(grid grid, colorToPaint color) {
	grid[r.panel] = colorToPaint
}

func (r *robot) turn(direction direction) {
	switch direction {
	case turnLeft:
		r.orientation = r.orientation.turnLeft()
	case turnRight:
		r.orientation = r.orientation.turnRight()
	}
}

func (r *robot) moveOneStepForward() {
	switch r.orientation {
	case up:
		r.panel = panel{r.panel.x, r.panel.y + 1}
	case down:
		r.panel = panel{r.panel.x, r.panel.y - 1}
	case right:
		r.panel = panel{r.panel.x + 1, r.panel.y}
	case left:
		r.panel = panel{r.panel.x - 1, r.panel.y}
	}
}

func (g grid) numberOfPaintedPanels() int {
	return len(g)
}

func (g grid) topLeftPanel() panel {
	xMin := math.MaxInt64
	yMin := math.MaxInt64
	for panel := range g {
		if panel.x < xMin {
			xMin = panel.x
		}
		if panel.y < yMin {
			yMin = panel.y
		}
	}
	return panel{xMin, yMin}
}

func (g grid) bottomRightPanel() panel {
	xMax := math.MinInt64
	yMax := math.MinInt64
	for panel := range g {
		if panel.x > xMax {
			xMax = panel.x
		}
		if panel.y > yMax {
			yMax = panel.y
		}
	}
	return panel{xMax, yMax}
}

func (g grid) renderHull(topLeft, bottomRight panel) string {
	rows := bottomRight.x - topLeft.x + 1
	cols := bottomRight.y - topLeft.y + 1

	// the way we traverse the matrix is a bit complex but needed to not get a rotated image
	var sb strings.Builder
	sb.WriteString("\n")
	for y := cols - 1; y >= 0; y-- {
		for x := 0; x < rows; x++ {
			switch g[panel{topLeft.x + x, topLeft.y + y}] {
			case black:
				sb.WriteString(blackEmoji)
			case white:
				sb.WriteString(whiteEmoji)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
