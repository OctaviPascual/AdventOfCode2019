package day06

import (
	"errors"
	"fmt"
	"strings"
)

const (
	youID   = "YOU"
	santaID = "SAN"
)

// orbitalRelationship reads as "satellite is in orbit around parent" or "satellite orbits parent"
type orbitalRelationship struct {
	parent    string
	satellite string
}

type orbitMap struct {
	orbitalRelationships []orbitalRelationship
}

type planet struct {
	id           string
	orbitsAround *planet
}

// Day holds the data needed to solve part one and part two
type Day struct {
	orbitMap orbitMap
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	orbitMapString := strings.Split(input, "\n")

	orbitMap, err := parseOrbitMap(orbitMapString)
	if err != nil {
		return nil, fmt.Errorf("invalid orbit map %s: %w", orbitMapString, err)
	}

	return &Day{
		orbitMap: *orbitMap,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	galaxy, err := createGalaxy(d.orbitMap)
	if err != nil {
		return "", fmt.Errorf("could not create galaxy: %w", err)
	}

	totalOrbits := getTotalOrbits(galaxy)
	return fmt.Sprintf("%d", totalOrbits), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	galaxy, err := createGalaxy(d.orbitMap)
	if err != nil {
		return "", fmt.Errorf("could not create galaxy: %w", err)
	}

	minimumOrbitalTransfers, err := getMinimumOrbitalTransfers(galaxy)
	return fmt.Sprintf("%d", minimumOrbitalTransfers), nil
}

func parseOrbitMap(orbitMapString []string) (*orbitMap, error) {
	orbitalRelationships := make([]orbitalRelationship, 0, len(orbitMapString))

	for _, orbitalRelationshipString := range orbitMapString {

		orbitalRelationship, err := parseOrbitalRelationship(orbitalRelationshipString)
		if err != nil {
			return nil, fmt.Errorf("invalid orbital relationship %s: %w", orbitalRelationshipString, err)
		}

		orbitalRelationships = append(orbitalRelationships, *orbitalRelationship)
	}

	return &orbitMap{
		orbitalRelationships: orbitalRelationships,
	}, nil
}

func parseOrbitalRelationship(orbitalRelationshipString string) (*orbitalRelationship, error) {
	relationship := strings.Split(orbitalRelationshipString, ")")

	if len(relationship) != 2 {
		return nil, fmt.Errorf("invalid format: %s", orbitalRelationshipString)
	}

	return &orbitalRelationship{
		parent:    relationship[0],
		satellite: relationship[1],
	}, nil
}

func createGalaxy(orbitMap orbitMap) (map[string]*planet, error) {
	galaxy := make(map[string]*planet)

	for _, orbitalRelationship := range orbitMap.orbitalRelationships {

		parentPlanet := getOrCreatePlanet(galaxy, orbitalRelationship.parent)
		satellitePlanet := getOrCreatePlanet(galaxy, orbitalRelationship.satellite)

		if satellitePlanet.orbitsAround != nil {
			return nil, fmt.Errorf("a satellite can orbit around one planet at most: %v", satellitePlanet)
		}

		satellitePlanet.orbitsAround = parentPlanet
	}

	return galaxy, nil
}

func getOrCreatePlanet(galaxy map[string]*planet, plantedID string) *planet {
	if _, ok := galaxy[plantedID]; !ok {
		galaxy[plantedID] = &planet{
			id: plantedID,
		}
	}
	return galaxy[plantedID]
}

func getTotalOrbits(galaxy map[string]*planet) int {
	totalOrbits := 0
	for _, planet := range galaxy {
		totalOrbits += planet.directAndIndirectOrbits()
	}
	return totalOrbits
}

func (p *planet) directAndIndirectOrbits() int {
	var path []string
	p.pathToGalaxyRoot(&path)
	return len(path) - 1
}

func (p *planet) isGalaxyRoot() bool {
	return p.orbitsAround == nil
}

func getMinimumOrbitalTransfers(galaxy map[string]*planet) (int, error) {
	you, ok := galaxy[youID]
	if !ok {
		return 0, errors.New("you are not in the galaxy")
	}

	santa, ok := galaxy[santaID]
	if !ok {
		return 0, errors.New("santa is not in the galaxy")
	}

	var path1 []string
	you.pathToGalaxyRoot(&path1)

	var path2 []string
	santa.pathToGalaxyRoot(&path2)

	ancestor, err := lowestCommonAncestor(path1, path2)
	if err != nil {
		return 0, fmt.Errorf("santa in not reachable by you: %w", err)
	}

	return orbitalTransfers(ancestor, path1) + orbitalTransfers(ancestor, path2), nil
}

func (p *planet) pathToGalaxyRoot(path *[]string) {
	*path = append(*path, p.id)

	if !p.isGalaxyRoot() {
		p.orbitsAround.pathToGalaxyRoot(path)
	}
}

func lowestCommonAncestor(path1, path2 []string) (string, error) {
	if len(path1) == 0 || len(path2) == 0 {
		return "", fmt.Errorf("no common ancestor between %v and %v", path1, path2)
	}

	i := len(path1) - 1
	j := len(path2) - 1

	if path1[i] != path2[j] {
		return "", fmt.Errorf("no common ancestor between %v and %v", path1, path2)
	}

	for {
		if path1[i-1] != path2[j-1] {
			return path1[i], nil
		}
		i--
		j--
		if i == 0 {
			return path1[0], nil
		}
		if j == 0 {
			return path2[0], nil
		}
	}
}

func orbitalTransfers(ancestor string, path []string) int {
	orbitalTransfers := 0
	// we don't have to count the first orbital transfer
	for _, planetID := range path[1:] {
		if planetID == ancestor {
			break
		}
		orbitalTransfers++
	}
	return orbitalTransfers
}
