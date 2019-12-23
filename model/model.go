package model

type Day interface {
	SolvePartOne() Answer
}

type Answer interface {
	Format() string
}
