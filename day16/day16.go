package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	signal signal
}

type digit int

type signal struct {
	digits []digit
}

type pattern struct {
	digits  []digit
	current int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	signal, err := parseSignal(input)
	if err != nil {
		return nil, fmt.Errorf("invalid signal: %w", err)
	}

	return &Day{
		signal: signal,
	}, nil
}

const (
	digitsToShow          = 8
	phases                = 100
	realSignalRepetitions = 10000
	offsetLength          = 7
)

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	finalSignal := runFFT(d.signal, phases)

	var sb strings.Builder
	sb.Grow(digitsToShow)

	for _, d := range finalSignal.digits[:digitsToShow] {
		sb.WriteString(strconv.Itoa(int(d)))
	}

	return sb.String(), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	realSignal := createRealSignal(d.signal)
	offset := realSignal.offset()

	finalRealSignal, err := runFFTWithOffset(realSignal, phases, offset)
	if err != nil {
		return "", fmt.Errorf("could not compute final signal: %w", err)
	}

	var sb strings.Builder
	sb.Grow(digitsToShow)

	for _, d := range finalRealSignal.digits[offset:(offset + digitsToShow)] {
		sb.WriteString(strconv.Itoa(int(d)))
	}

	return sb.String(), nil
}

func parseSignal(input string) (signal, error) {
	digits := make([]digit, 0, len(input))
	for _, r := range input {
		digit, err := newDigitFromRune(r)
		if err != nil {
			return signal{}, fmt.Errorf("invalid digit: %w", err)
		}
		digits = append(digits, digit)
	}
	return signal{
		digits: digits,
	}, nil
}

func newDigitFromRune(r rune) (digit, error) {
	n := int(r - '0')
	if n < 0 || n > 9 {
		return 0, fmt.Errorf("rune must represent a digit: %v", r)
	}
	return digit(n), nil
}

func newDigitFromInt(n int) (digit, error) {
	if n < 0 || n > 9 {
		return 0, fmt.Errorf("int must represent a digit: %v", n)
	}
	return digit(n), nil
}

func (d digit) multiply(digit digit) int {
	return int(d) * int(digit)
}

func runFFT(signal signal, phases int) signal {
	for i := 0; i < phases; i++ {
		signal = signal.next()
	}
	return signal
}

func createRealSignal(originalSignal signal) signal {
	digits := make([]digit, 0, originalSignal.length()*realSignalRepetitions)
	for i := 0; i < realSignalRepetitions; i++ {
		digits = append(digits, originalSignal.digits...)
	}
	return signal{
		digits: digits,
	}
}

func runFFTWithOffset(signal signal, phases, offset int) (signal, error) {
	if offset < signal.length()/2 {
		return signal, fmt.Errorf(
			"offset (%v) must be greater than half of the lentgh of signal (%v)", offset, signal.length(),
		)
	}
	for i := 0; i < phases; i++ {
		signal = signal.nextOnlySecondHalf()
	}
	return signal, nil
}

var basePattern = []digit{0, 1, 0, -1}

func newPattern(position int) *pattern {
	digits := make([]digit, 0, len(basePattern)*position)

	for _, d := range basePattern {
		for i := 0; i < position; i++ {
			digits = append(digits, d)
		}
	}

	return &pattern{
		digits: digits,
	}
}

func (p *pattern) next() digit {
	d := p.digits[p.current]
	p.current = (p.current + 1) % len(p.digits)
	return d
}

func (s signal) next() signal {
	digits := make([]digit, 0, s.length())
	for position := 1; position <= s.length(); position++ {
		pattern := newPattern(position)
		_ = pattern.next()

		sum := 0
		for _, d := range s.digits {
			sum += d.multiply(pattern.next())
		}
		digit, _ := newDigitFromInt(util.Abs(sum) % 10)
		digits = append(digits, digit)
	}
	return signal{
		digits: digits,
	}
}

func (s signal) length() int {
	return len(s.digits)
}

func (s signal) offset() int {
	offset := 0
	for _, d := range s.digits[:offsetLength] {
		offset *= 10
		offset += int(d)
	}

	return offset
}

// We can compute the second half of the signal much faster.
// Given a signal of size N, from the second half of the signal to the end, at position i the pattern will be
// i-1 zeroes followed by N-i+1 ones. For instance, if N=6 then for position i=4 the pattern will be 000111.
// With this pattern it's possible to directly derive the next signal s' from the current signal s:
//	- position i=N-1: pattern is 0..01 so s'[N-1] = s[N-1]
//	- position i=N-2: pattern is 0..011 so s'[N-2] = (s[N-2] + s[N-1]) % 10 = (s[N-2] + s'[N-1]) % 10
//	- position i=N-3: pattern is 0..0111 so s'[N-3] = (s[N-3] + s[N-2] + s[N-1]) % 10 = (s[N-3] + s'[N-2]) % 10
func (s signal) nextOnlySecondHalf() signal {
	digits := make([]digit, s.length())
	copy(digits, s.digits)

	for i := s.length() - 2; i >= 0; i-- {
		digits[i] = (digits[i+1] + s.digits[i]) % 10
	}

	return signal{
		digits: digits,
	}
}
