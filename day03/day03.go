package day03

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type direction rune

const (
	up    direction = 'U'
	right direction = 'R'
	down  direction = 'D'
	left  direction = 'L'
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

// Day holds the data needed to solve part one and part two
type Day struct {
	wire1 wire
	wire2 wire
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
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

	return &Day{
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
	case up, left, right, down:
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

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	circuit := circuit{
		ports1: getPortsFromWire(d.wire1),
		ports2: getPortsFromWire(d.wire2),
	}
	distance, err := circuit.findMinimumDistance()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", distance), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	circuit := circuit{
		ports1: getPortsFromWire(d.wire1),
		ports2: getPortsFromWire(d.wire2),
	}
	signalDelay, err := circuit.findMinimumSignalDelay()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", signalDelay), nil
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
			steps++
		}
	}
	return ports
}

func (p port) nextPort(direction direction) port {
	switch direction {
	case up:
		return port{x: p.x, y: p.y + 1}
	case down:
		return port{x: p.x, y: p.y - 1}
	case right:
		return port{x: p.x + 1, y: p.y}
	case left:
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
