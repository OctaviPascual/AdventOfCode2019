package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/day01"
	"github.com/OctaviPascual/AdventOfCode2019/model"
)

var days = []struct {
	filename    string
	constructor func(input string) model.Day
}{
	{
		filename:    "./day01/day01.txt",
		constructor: day01.NewDay,
	},
}

func main() {
	for _, day := range days {
		bytes, err := ioutil.ReadFile(day.filename)
		if err != nil {
			panic(err)
		}
		input := string(bytes)
		input = strings.TrimSuffix(input, "\n")

		day := day.constructor(input)
		fmt.Println(day.SolvePartOne())
		fmt.Println(day.SolvePartTwo())
	}
}
