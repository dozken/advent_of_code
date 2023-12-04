package day1_test

import (
	"github.com/dozken/aoc23/pkg/day1"
	"github.com/dozken/aoc23/pkg/util"
	"testing"
)

func testDataP1() string {
	return `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchetzu`
}

func testDataP2() string {
	return `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
}

func TestP1_testData(t *testing.T) {
	got := util.SumFromText(testDataP1(), day1.ParseLinePart1)
	expected := 142
	if got != expected {
		t.Errorf("P1() = %v, want %v", got, expected)
	}
}

func TestP1FromFile_testData(t *testing.T) {
	got := util.SumFromFile("p1.input", day1.ParseLinePart1)
	expected := 54239
	if got != expected {
		t.Errorf("P1FromFile() = %v, want %v", got, expected)
	}
}

func TestP2_testData(t *testing.T) {
	actual := util.SumFromText(testDataP2(), day1.ParseLinePart2)
	expected := 281
	if actual != expected {
		t.Errorf("P2() = %v, want %v", actual, expected)
	}
}

func TestP2FromFile_testData(t *testing.T) {
	got := util.SumFromFile("p1.input", day1.ParseLinePart2)
	expected := 55343
	if got != expected {
		t.Errorf("P1FromFile() = %v, want %v", got, expected)
	}
}
