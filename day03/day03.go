package day03

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/model"
)

type direction rune

const (
	Up    direction = 'U'
	Right           = 'R'
	Down            = 'D'
	Left            = 'L'
)

type segment struct {
	direction direction
	length    int
}

type wire []segment

type port struct {
	x, y int
}

var centralPort = port{
	x: 0,
	y: 0,
}

type circuit struct {
	ports1 map[port]int
	ports2 map[port]int
}

type distance uint
type signalDelay uint

type day struct {
	wire1 wire
	wire2 wire
}

func NewDay(input string) (model.Day, error) {
	paths := strings.Split(input, "\n")
	if len(paths) != 2 {
		return nil, fmt.Errorf("invalid number of wire paths %d", len(paths))
	}

	path1, err := parseWire(paths[0])
	if err != nil {
		return nil, fmt.Errorf("invalid wire path %s: %w", paths[0], err)
	}

	path2, err := parseWire(paths[1])
	if err != nil {
		return nil, fmt.Errorf("invalid wire path %s: %w", paths[1], err)
	}

	return &day{
		wire1: path1,
		wire2: path2,
	}, nil
}

func parseWire(wireString string) (wire, error) {
	segmentsString := strings.Split(wireString, ",")
	wire := make(wire, 0, len(segmentsString))
	for _, segmentString := range segmentsString {
		segment, err := parseSegment(segmentString)
		if err != nil {
			return nil, fmt.Errorf("invalid segment %s: %w", segmentString, err)
		}
		wire = append(wire, segment)
	}
	return wire, nil
}

func parseSegment(segmentString string) (segment, error) {
	if len(segmentString) < 2 {
		return segment{}, fmt.Errorf("invalid segment %s", segmentString)
	}

	direction := direction(segmentString[0])
	switch direction {
	case Up, Left, Right, Down:
	default:
		return segment{}, fmt.Errorf("invalid direction %c", segmentString[0])
	}

	length, err := strconv.Atoi(segmentString[1:])
	if err != nil {
		return segment{}, fmt.Errorf("invalid length %s: %w", segmentString[1:], err)
	}

	return segment{
		direction: direction,
		length:    length,
	}, nil
}

func (d distance) String() string {
	return fmt.Sprintf("%d", d)
}

func (d day) SolvePartOne() (model.Answer, error) {
	circuit := circuit{
		ports1: getPortsFromWire(d.wire1),
		ports2: getPortsFromWire(d.wire2),
	}
	return circuit.findMinimumDistance()
}

func (sd signalDelay) String() string {
	return fmt.Sprintf("%d", sd)
}

func (d day) SolvePartTwo() (model.Answer, error) {
	circuit := circuit{
		ports1: getPortsFromWire(d.wire1),
		ports2: getPortsFromWire(d.wire2),
	}
	return circuit.findMinimumSignalDelay()
}

func getPortsFromWire(wire wire) map[port]int {
	ports := make(map[port]int, len(wire))
	currentPort := centralPort
	steps := 1
	for _, segment := range wire {
		for i := 0; i < segment.length; i++ {
			currentPort = currentPort.nextPort(segment.direction)
			if _, ok := ports[currentPort]; !ok {
				ports[currentPort] = steps
			}
			steps += 1
		}
	}
	return ports
}

func (p port) nextPort(direction direction) port {
	switch direction {
	case Up:
		return port{x: p.x, y: p.y + 1}
	case Down:
		return port{x: p.x, y: p.y - 1}
	case Right:
		return port{x: p.x + 1, y: p.y}
	case Left:
		return port{x: p.x - 1, y: p.y}
	}
	return p
}

func (c circuit) findMinimumDistance() (distance, error) {
	costFunction := func(port port) uint {
		return uint(port.distanceToCentralPort())
	}
	res, err := c.minimizeCostFunction(costFunction)
	return distance(res), err
}

func (c circuit) findMinimumSignalDelay() (signalDelay, error) {
	costFunction := func(port port) uint {
		return uint(c.signalDelay(port))
	}
	res, err := c.minimizeCostFunction(costFunction)
	return signalDelay(res), err
}

type costFunction func(port port) uint

func (c circuit) minimizeCostFunction(costFunction costFunction) (uint, error) {
	intersections := c.findIntersections()
	if len(intersections) == 0 {
		return 0, fmt.Errorf("there are no intersections between the two wires")
	}

	minimumCost := costFunction(intersections[0])
	for _, intersection := range intersections {
		currentCost := costFunction(intersection)
		if currentCost < minimumCost {
			minimumCost = currentCost
		}
	}
	return minimumCost, nil
}

func (c circuit) findIntersections() []port {
	var intersections []port
	for port := range c.ports1 {
		if _, ok := c.ports2[port]; ok {
			intersections = append(intersections, port)
		}
	}
	return intersections
}

func (p port) distanceToCentralPort() distance {
	return distance(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (c circuit) signalDelay(port port) signalDelay {
	return signalDelay(c.ports1[port] + c.ports2[port])
}
