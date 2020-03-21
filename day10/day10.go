package day10

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

const (
	asteroid = '#'
	empty    = '.'
)

// Day holds the data needed to solve part one and part two
type Day struct {
	asteroidMap [][]bool
}

type position struct {
	// instead of working with the coordinates suggested in the statement, we use regular matrix coordinates
	// (that is, (i=0, j=0) is the top left position, (i=0, j=1) is the position to its right)
	i, j int
}

type monitoringStation struct {
	position    position
	asteroidMap [][]bool

	positionsGroupedByAngle [][]position
	vaporizedAsteroids      []position
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	asteroidMap, err := parseAsteroidMap(input)
	if err != nil {
		return nil, err
	}

	return &Day{
		asteroidMap: asteroidMap,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	bestMonitoringStation := findBestMonitoringStationLocation(d.asteroidMap)
	return fmt.Sprintf("%d", bestMonitoringStation.asteroidsDetected()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	bestMonitoringStation := findBestMonitoringStationLocation(d.asteroidMap)
	bestMonitoringStation.computePositionsGroupedByAngle()
	bestMonitoringStation.runLaser()
	twoHundredthVaporizedAsteroid := bestMonitoringStation.vaporizedAsteroids[200-1]
	return fmt.Sprintf("%d", twoHundredthVaporizedAsteroid.i+twoHundredthVaporizedAsteroid.j*100), nil
}

func parseAsteroidMap(input string) ([][]bool, error) {
	lines := strings.Split(input, "\n")

	asteroidMap := make([][]bool, len(lines))
	for i, line := range lines {
		asteroidMap[i] = make([]bool, len(line))
		for j, position := range line {
			switch position {
			case empty:
				asteroidMap[i][j] = false
			case asteroid:
				asteroidMap[i][j] = true
			default:
				return nil, fmt.Errorf("invalid position %c", position)
			}
		}
	}
	return asteroidMap, nil
}

func findBestMonitoringStationLocation(asteroidMap [][]bool) monitoringStation {
	rows := len(asteroidMap)
	cols := len(asteroidMap[0])
	maxAsteroidsDetected := 0
	var bestMonitoringStation monitoringStation
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if asteroidMap[i][j] {
				monitoringStation := monitoringStation{
					position:    position{i, j},
					asteroidMap: asteroidMap,
				}
				asteroidsDetected := monitoringStation.asteroidsDetected()
				if asteroidsDetected > maxAsteroidsDetected {
					maxAsteroidsDetected = asteroidsDetected
					bestMonitoringStation = monitoringStation
				}
			}
		}
	}
	return bestMonitoringStation
}

func (ms monitoringStation) asteroidsDetected() int {
	rows := len(ms.asteroidMap)
	cols := len(ms.asteroidMap[0])
	asteroidsDetected := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ms.asteroidMap[i][j] && ms.hasDirectLineOfSight(position{i, j}) {
				asteroidsDetected++
			}
		}
	}
	return asteroidsDetected
}

func (ms monitoringStation) hasDirectLineOfSight(p position) bool {
	if p == ms.position {
		return false
	}

	// (di, dj) is a difference vector between the position p and the monitoring station ms
	di := p.i - ms.position.i
	dj := p.j - ms.position.j

	var iStep, jStep, steps int
	gcd := gcd(abs(di), abs(dj))
	if gcd == 0 {
		if di == 0 {
			iStep = 0
			jStep = 1
			steps = dj
		} else {
			iStep = 1
			jStep = 0
			steps = di
		}
	} else {
		iStep = di / gcd
		jStep = dj / gcd
		steps = gcd
	}

	// check all positions in direct line of sight between monitoring station ms and position p
	for k := 1; k < steps; k++ {
		if ms.asteroidMap[ms.position.i+k*iStep][ms.position.j+k*jStep] {
			return false
		}
	}
	return true
}

func (ms *monitoringStation) computePositionsGroupedByAngle() {
	rows := len(ms.asteroidMap)
	cols := len(ms.asteroidMap[0])

	type enrichedPosition struct {
		position    position
		pseudoAngle *big.Rat
	}

	enrichedPositions := make([]enrichedPosition, 0, rows*cols-1)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			position := position{i, j}
			if position != ms.position {
				di := position.i - ms.position.i
				dj := position.j - ms.position.j
				pseudoAngle := pseudoAngle(di, dj)
				enrichedPosition := enrichedPosition{
					position:    position,
					pseudoAngle: pseudoAngle,
				}
				enrichedPositions = append(enrichedPositions, enrichedPosition)
			}
		}
	}
	sort.Slice(enrichedPositions, func(i, j int) bool {
		return enrichedPositions[i].pseudoAngle.Cmp(enrichedPositions[j].pseudoAngle) < 0
	})

	// group positions that have the same pseudo angle in same bucket
	i := 0
	j := 0
	for i < len(enrichedPositions) {
		ms.positionsGroupedByAngle = append(ms.positionsGroupedByAngle, make([]position, 0))
		ms.positionsGroupedByAngle[j] = append(ms.positionsGroupedByAngle[j], enrichedPositions[i].position)
		pseudoAngle := enrichedPositions[i].pseudoAngle
		i++
		for i < len(enrichedPositions) && enrichedPositions[i].pseudoAngle.Cmp(pseudoAngle) == 0 {
			ms.positionsGroupedByAngle[j] = append(ms.positionsGroupedByAngle[j], enrichedPositions[i].position)
			i++
		}
		j++
	}
}

func (ms *monitoringStation) runLaser() []position {
	totalAsteroids := ms.asteroidsDetected()
	i := 0
	for totalAsteroids > 0 {
		for _, position := range ms.positionsGroupedByAngle[i] {
			if ms.asteroidMap[position.i][position.j] && ms.hasDirectLineOfSight(position) {
				ms.asteroidMap[position.i][position.j] = false
				ms.vaporizedAsteroids = append(ms.vaporizedAsteroids, position)
				totalAsteroids--
				break
			}
		}
		i = (i + 1) % len(ms.positionsGroupedByAngle)
	}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// pseudoAngle returns a number from the range [-2 .. 2] which is monotonic in the angle the vector (di, dj)
// makes against the x axis
// https://stackoverflow.com/q/16542042
func pseudoAngle(di, dj int) *big.Rat {
	r := big.NewRat(int64(di), int64(abs(di)+abs(dj)))
	if dj < 0 {
		return r.Sub(big.NewRat(1, 1), r)
	}
	return r.Sub(r, big.NewRat(1, 1))
}
