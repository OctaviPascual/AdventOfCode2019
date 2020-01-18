package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type passwordRange struct {
	from, to int
}

type password int

type criterion func(password password) bool

// Day holds the data needed to solve part one and part two
type Day struct {
	passwordRange passwordRange
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	passwordRange, err := parsePasswordRange(input)
	if err != nil {
		return nil, fmt.Errorf("invalid password range %s: %w", input, err)
	}

	return &Day{
		passwordRange: passwordRange,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	criteria := []criterion{
		password.isSixDigitNumber,
		password.hasTwoAdjacentDigits,
		password.hasNeverDecreasingDigits,
	}
	meetingCriteriaPasswords := d.meetingCriteriaPasswords(criteria)
	return fmt.Sprintf("%d", len(meetingCriteriaPasswords)), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	criteria := []criterion{
		password.isSixDigitNumber,
		password.hasNeverDecreasingDigits,
		password.hasExactlyTwoAdjacentDigits,
	}
	meetingCriteriaPasswords := d.meetingCriteriaPasswords(criteria)
	return fmt.Sprintf("%d", len(meetingCriteriaPasswords)), nil
}

func parsePasswordRange(passwordRangeString string) (passwordRange, error) {
	ranges := strings.Split(passwordRangeString, "-")

	if len(ranges) != 2 {
		return passwordRange{}, fmt.Errorf("invalid number of ranges %d", len(ranges))
	}

	from, err := strconv.Atoi(ranges[0])
	if err != nil {
		return passwordRange{}, fmt.Errorf("invalid from %s: %w", ranges[0], err)
	}

	to, err := strconv.Atoi(ranges[1])
	if err != nil {
		return passwordRange{}, fmt.Errorf("invalid to %s: %w", ranges[1], err)
	}

	if from > to {
		return passwordRange{}, fmt.Errorf("from (%d) can't be greater than to (%d)", from, to)
	}

	return passwordRange{
		from: from,
		to:   to,
	}, nil

}

func (d Day) meetingCriteriaPasswords(criteria []criterion) []password {
	var meetingCriteriaPasswords []password
	for i := d.passwordRange.from; i <= d.passwordRange.to; i++ {
		password := password(i)
		if password.meetsCriteria(criteria) {
			meetingCriteriaPasswords = append(meetingCriteriaPasswords, password)
		}
	}
	return meetingCriteriaPasswords
}

func (p password) meetsCriteria(criteria []criterion) bool {
	for _, criterion := range criteria {
		if !criterion(p) {
			return false
		}
	}
	return true
}

func (p password) isSixDigitNumber() bool {
	return p >= 100000 && p < 1000000
}

func (p password) hasTwoAdjacentDigits() bool {
	return twoAdjacentDigits(int(p))
}

func twoAdjacentDigits(n int) bool {
	if n < 10 {
		return false
	}

	lastDigit := n % 10
	penultimateDigit := (n / 10) % 10
	return penultimateDigit == lastDigit || twoAdjacentDigits(n/10)
}

func (p password) hasNeverDecreasingDigits() bool {
	return neverDecreasingDigits(int(p))
}

func neverDecreasingDigits(n int) bool {
	if n < 10 {
		return true
	}

	lastDigit := n % 10
	penultimateDigit := (n / 10) % 10
	return lastDigit >= penultimateDigit && neverDecreasingDigits(n/10)
}

func (p password) hasExactlyTwoAdjacentDigits() bool {
	return exactlyTwoAdjacentDigits(int(p))
}

func exactlyTwoAdjacentDigits(n int) bool {
	if n < 100 {
		return twoAdjacentDigits(n)
	}

	lastDigit := n % 10
	penultimateDigit := (n / 10) % 10
	antePenultimateDigit := (n / 100) % 10

	if lastDigit == penultimateDigit && penultimateDigit != antePenultimateDigit {
		return true
	}
	return exactlyTwoAdjacentDigits(removeLastEqualDigits(n, lastDigit))
}

func removeLastEqualDigits(n, digit int) int {
	lastDigit := n % 10
	if lastDigit == digit {
		return removeLastEqualDigits(n/10, digit)
	}
	return n
}
