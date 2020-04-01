package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/day01"
	"github.com/OctaviPascual/AdventOfCode2019/day02"
	"github.com/OctaviPascual/AdventOfCode2019/day03"
	"github.com/OctaviPascual/AdventOfCode2019/day04"
	"github.com/OctaviPascual/AdventOfCode2019/day05"
	"github.com/OctaviPascual/AdventOfCode2019/day06"
	"github.com/OctaviPascual/AdventOfCode2019/day07"
	"github.com/OctaviPascual/AdventOfCode2019/day08"
	"github.com/OctaviPascual/AdventOfCode2019/day09"
	"github.com/OctaviPascual/AdventOfCode2019/day10"
	"github.com/OctaviPascual/AdventOfCode2019/day11"
	"github.com/OctaviPascual/AdventOfCode2019/day12"
	"github.com/OctaviPascual/AdventOfCode2019/day13"
)

// Day is the interface that wraps SolvePartOne and SolvePartTwo methods
type Day interface {
	SolvePartOne() (string, error)
	SolvePartTwo() (string, error)
}

var days = []struct {
	filename    string
	constructor func(input string) (Day, error)
}{
	{
		filename: "./day01/day01.txt",
		constructor: func(input string) (Day, error) {
			return day01.NewDay(input)
		},
	},
	{
		filename: "./day02/day02.txt",
		constructor: func(input string) (Day, error) {
			return day02.NewDay(input)
		},
	},
	{
		filename: "./day03/day03.txt",
		constructor: func(input string) (Day, error) {
			return day03.NewDay(input)
		},
	},
	{
		filename: "./day04/day04.txt",
		constructor: func(input string) (Day, error) {
			return day04.NewDay(input)
		},
	},
	{
		filename: "./day05/day05.txt",
		constructor: func(input string) (Day, error) {
			return day05.NewDay(input)
		},
	},
	{
		filename: "./day06/day06.txt",
		constructor: func(input string) (Day, error) {
			return day06.NewDay(input)
		},
	},
	{
		filename: "./day07/day07.txt",
		constructor: func(input string) (Day, error) {
			return day07.NewDay(input)
		},
	},
	{
		filename: "./day08/day08.txt",
		constructor: func(input string) (Day, error) {
			return day08.NewDay(input)
		},
	},
	{
		filename: "./day09/day09.txt",
		constructor: func(input string) (Day, error) {
			return day09.NewDay(input)
		},
	},
	{
		filename: "./day10/day10.txt",
		constructor: func(input string) (Day, error) {
			return day10.NewDay(input)
		},
	},
	{
		filename: "./day11/day11.txt",
		constructor: func(input string) (Day, error) {
			return day11.NewDay(input)
		},
	},
	{
		filename: "./day12/day12.txt",
		constructor: func(input string) (Day, error) {
			return day12.NewDay(input)
		},
	},
	{
		filename: "./day13/day13.txt",
		constructor: func(input string) (Day, error) {
			return day13.NewDay(input)
		},
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
