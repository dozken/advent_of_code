package day2_test

import (
	"github.com/dozken/aoc23/pkg/day2"
	"github.com/dozken/aoc23/pkg/util"
	"testing"
)

func TestP1_sample(t *testing.T) {
	got := util.SumFromFile("p1.sample", day2.ParseLinePart1)
	expected := 8
	if got != expected {
		t.Errorf("P1() = %v, want %v", got, expected)
	}
}

func TestP1_data(t *testing.T) {
    got := util.SumFromFile("p1.data", day2.ParseLinePart1)
    expected := 8
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}

// func TestP2_testData(t *testing.T) {
// 	actual := util.SumFromText(testDataP2(), day1.ParseLinePart2)
// 	expected := 281
// 	if actual != expected {
// 		t.Errorf("P2() = %v, want %v", actual, expected)
// 	}
// }
//
// func TestP2FromFile_testData(t *testing.T) {
// 	got := util.SumFromFile("p1.input", day1.ParseLinePart2)
// 	expected := 54239
// 	if got != expected {
// 		t.Errorf("P1FromFile() = %v, want %v", got, expected)
// 	}
// }
