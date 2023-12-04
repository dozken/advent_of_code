package day3_test

import (
	"testing"
	"github.com/dozken/aoc23/pkg/day3"
)

func TestP1_sample(t *testing.T) {
	got := day3.SumFromFile("p1.sample")
	expected := 4361
	if got != expected {
		t.Errorf("P1() = %v, want %v", got, expected)
	}
}

func TestP1_data(t *testing.T) {
    got := day3.SumFromFile("p1.data")
    expected := 2105
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}
//
// func TestP2_sample(t *testing.T) {
// 	got := util.SumFromFile("p1.sample", day2.ParseLinePart2)
// 	expected := 2286
// 	if got != expected {
// 		t.Errorf("P2() = %v, want %v", got, expected)
// 	}
// }
//
// func TestP2_data(t *testing.T) {
//     got := util.SumFromFile("p1.data", day2.ParseLinePart2)
//     expected := 2105
//     if got != expected {
//         t.Errorf("P2() = %v, want %v", got, expected)
//     }
// }
