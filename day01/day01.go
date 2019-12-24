package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/model"
)

type fuel int

type module struct {
	mass int
}

type spacecraft struct {
	modules []module
}

type day struct {
	spacecraft spacecraft
}

func NewDay(input string) model.Day {
	lines := strings.Split(input, "\n")
	modules := make([]module, 0, len(input))
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		module := module{
			mass: mass,
		}
		modules = append(modules, module)
	}
	return &day{
		spacecraft: spacecraft{
			modules: modules,
		},
	}
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

func (f fuel) String() string {
	return fmt.Sprintf("%d", f)
}

func (d day) SolvePartOne() model.Answer {
	return d.spacecraft.totalFuelRequired()
}

func (d day) SolvePartTwo() model.Answer {
	return d.spacecraft.totalFuelRequiredWithAddedFuel()
}
