package day15

import (
	"errors"
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
	"github.com/OctaviPascual/AdventOfCode2019/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

type command int
type status int
type cell int

const (
	north command = 1
	south command = 2
	west  command = 3
	east  command = 4

	foundWall   status = 0
	moved       status = 1
	foundOxygen status = 2

	unknown cell = 0
	wall    cell = 1
	empty   cell = 2
	oxygen  cell = 3
)

func (c command) toInt() int {
	return int(c)
}

var initialPosition = position{x: 0, y: 0}

type position struct {
	x, y int
}

type repairDroid struct {
	program       *intcode.Intcode
	position      position
	space         map[position]cell
	inputChannel  chan int
	outputChannel chan int
	foundOxygen   bool
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

	repairDroid := newRepairDroid(intcodeProgram)
	fewestNumberOfCommandsToOxygen, err := repairDroid.fewestNumberOfCommandsToOxygen(initialPosition)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", fewestNumberOfCommandsToOxygen), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
	if err != nil {
		return "", err
	}

	repairDroid := newRepairDroid(intcodeProgram)
	minutesToFillWithOxygen, err := repairDroid.minutesToFillWithOxygen()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", minutesToFillWithOxygen), nil
}

func newRepairDroid(intcodeProgram *intcode.Intcode) *repairDroid {
	space := make(map[position]cell)
	space[initialPosition] = empty

	rd := &repairDroid{
		program:       intcodeProgram,
		position:      initialPosition,
		space:         space,
		inputChannel:  make(chan int),
		outputChannel: make(chan int),
	}
	rd.exploreAllSpace()
	return rd
}

func (rd *repairDroid) fewestNumberOfCommandsToOxygen(source position) (int, error) {
	oxygenPosition := oxygenPosition(rd.space)
	isTarget := func(position position) bool { return position == oxygenPosition }

	commands, err := commandsToTarget(source, isTarget, rd.space)
	if err != nil {
		return -1, errors.New("could not find path to oxygen")
	}

	return len(commands), nil
}

func (rd *repairDroid) minutesToFillWithOxygen() (int, error) {
	maxMinutesToFill := 0

	for position, cell := range rd.space {
		if cell != empty {
			continue
		}

		fewestNumberOfCommandsToOxygen, err := rd.fewestNumberOfCommandsToOxygen(position)
		if err != nil {
			return -1, err
		}

		maxMinutesToFill = util.Max(fewestNumberOfCommandsToOxygen, maxMinutesToFill)
	}

	return maxMinutesToFill, nil
}

func (rd *repairDroid) exploreAllSpace() {
	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- rd.program.Run(rd.inputChannel, rd.outputChannel)
	}()

	for {
		isTarget := func(position position) bool { return rd.space[position] == unknown }
		commands, err := commandsToTarget(rd.position, isTarget, rd.space)
		if err != nil {
			return
		}

		for _, command := range commands {
			rd.move(command)
		}
	}
}

func oxygenPosition(space map[position]cell) position {
	for position, cell := range space {
		if cell == oxygen {
			return position
		}
	}
	panic("could not find oxygen in space")
}

func (rd *repairDroid) move(command command) {
	rd.inputChannel <- command.toInt()

	status := status(<-rd.outputChannel)
	switch status {
	case foundWall:
		p := rd.position.nextPosition(command)
		rd.space[p] = wall
	case moved:
		rd.position = rd.position.nextPosition(command)
		rd.space[rd.position] = empty
	case foundOxygen:
		rd.position = rd.position.nextPosition(command)
		rd.foundOxygen = true
		rd.space[rd.position] = oxygen
	}
}

// nextPosition returns the position that results from applying command to p
func (p position) nextPosition(command command) position {
	switch command {
	case north:
		return position{x: p.x, y: p.y + 1}
	case south:
		return position{x: p.x, y: p.y - 1}
	case west:
		return position{x: p.x - 1, y: p.y}
	case east:
		return position{x: p.x + 1, y: p.y}
	default:
		panic(fmt.Sprintf("unknown command %d", command))
	}
}

// neighbours returns the four adjacent positions to p
func (p position) neighbours() []position {
	return []position{
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y - 1},
		{x: p.x, y: p.y + 1},
	}
}

// commandTo returns the command that must be performed to go from p to position
func (p position) commandTo(position position) command {
	switch {
	case position.x == p.x && position.y == p.y+1:
		return north
	case position.x == p.x && position.y == p.y-1:
		return south
	case position.x == p.x-1 && position.y == p.y:
		return west
	case position.x == p.x+1 && position.y == p.y:
		return east
	default:
		panic(fmt.Sprintf("positions %v and %v are not adjacent", p, position))
	}
}

// buildCommands returns the commands that are needed to move to target
func buildCommands(target position, parent map[position]position) []command {
	source, ok := parent[target]
	if !ok {
		return nil
	}
	return append(buildCommands(source, parent), source.commandTo(target))
}

// commandsToTarget returns the commands needed to go from the source position to a target position
func commandsToTarget(
	source position,
	isTarget func(position position) bool,
	space map[position]cell,
) ([]command, error) {
	parent := make(map[position]position)

	queue := []position{source}
	visited := map[position]bool{source: true}

	for len(queue) > 0 {
		current := queue[0]

		if isTarget(current) {
			return buildCommands(current, parent), nil
		}

		for _, successor := range current.neighbours() {
			if space[successor] == wall {
				continue
			}

			if _, ok := visited[successor]; !ok {
				visited[successor] = true
				parent[successor] = current
				queue = append(queue, successor)
			}
		}
		queue = queue[1:]
	}
	return nil, errors.New("no reachable target position found in space")
}
