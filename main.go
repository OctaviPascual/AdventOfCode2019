package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/day01"
	"github.com/OctaviPascual/AdventOfCode2019/day02"
	"github.com/OctaviPascual/AdventOfCode2019/day03"
	"github.com/OctaviPascual/AdventOfCode2019/model"
)

var days = []struct {
	filename    string
	constructor func(input string) (model.Day, error)
}{
	{
		filename:    "./day01/day01.txt",
		constructor: day01.NewDay,
	},
	{
		filename:    "./day02/day02.txt",
		constructor: day02.NewDay,
	},
	{
		filename:    "./day03/day03.txt",
		constructor: day03.NewDay,
	},
}

func main() {
	for i, day := range days {
		fmt.Printf("\nRunning day %d\n", i+1)
		bytes, err := ioutil.ReadFile(day.filename)
		if err != nil {
			log.Fatalf("could not read file %s: %v", day.filename, err)
		}
		input := string(bytes)
		input = strings.TrimSuffix(input, "\n")

		day, err := day.constructor(input)
		if err != nil {
			log.Fatalf("could not create day %d: %v", i+1, err)
		}

		answer, err := day.SolvePartOne()
		if err != nil {
			log.Fatalf("could not solve part one for day %d: %v", i+1, err)
		}
		fmt.Printf("Part One: %s\n", answer)

		answer, err = day.SolvePartTwo()
		if err != nil {
			log.Fatalf("could not solve part two for day %d: %v", i+1, err)
		}
		fmt.Printf("Part Two: %s\n", answer)
	}
}
