package main

import (
	"fmt"
	"io/ioutil"

	"github.com/OctaviPascual/AdventOfCode2019/day01"
)

func main() {
	bytes, _ := ioutil.ReadFile("./day01/day01.txt")
	day := day01.NewDay(string(bytes))
	fmt.Println(day.SolvePartOne().Format())
	fmt.Println(day.SolvePartTwo().Format())
}
