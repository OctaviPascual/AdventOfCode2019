package model

type Day interface {
	SolvePartOne() Answer
	SolvePartTwo() Answer
}

type Answer interface {
	String() string
}
