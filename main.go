package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/day01"
)

func main() {
	bytes, _ := ioutil.ReadFile("./day01/day01.txt")
	input := string(bytes)
	input = strings.TrimSuffix(input, "\n")
	day := day01.NewDay(input)
	fmt.Println(day.SolvePartOne().Format())
	fmt.Println(day.SolvePartTwo().Format())
}
