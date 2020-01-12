package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type fuel int

type module struct {
	mass int
}

type spacecraft struct {
	modules []module
}

type Day struct {
	spacecraft spacecraft
}

func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")
	modules := make([]module, 0, len(input))
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid mass %s: %v", line, err)
		}
		module := module{
			mass: mass,
		}
		modules = append(modules, module)
	}
	return &Day{
		spacecraft: spacecraft{
			modules: modules,
		},
	}, nil
}

func (m module) fuelRequired() fuel {
	return fuel(m.mass/3 - 2)
}

func (f fuel) fuelRequired() fuel {
	fuelRequired := f/3 - 2
	if fuelRequired <= 0 {
		return 0
	}
	return fuelRequired + fuelRequired.fuelRequired()
}

func (s spacecraft) totalFuelRequired() fuel {
	totalFuelRequired := fuel(0)
	for _, module := range s.modules {
		totalFuelRequired += module.fuelRequired()
	}
	return totalFuelRequired
}

func (s spacecraft) totalFuelRequiredWithAddedFuel() fuel {
	totalFuelRequired := fuel(0)
	for _, module := range s.modules {
		moduleFuelRequired := module.fuelRequired()
		addedFuelRequired := moduleFuelRequired.fuelRequired()

		totalFuelRequired += moduleFuelRequired + addedFuelRequired
	}
	return totalFuelRequired
}

func (d Day) SolvePartOne() (string, error) {
	fuel := d.spacecraft.totalFuelRequired()
	return fmt.Sprintf("%d", fuel), nil
}

func (d Day) SolvePartTwo() (string, error) {
	fuel := d.spacecraft.totalFuelRequiredWithAddedFuel()
	return fmt.Sprintf("%d", fuel), nil
}
