package model

type Day interface {
	SolvePartOne() (Answer, error)
	SolvePartTwo() (Answer, error)
}

type Answer interface {
	String() string
}
