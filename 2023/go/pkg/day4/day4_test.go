package day4_test

import (
	"testing"

	"github.com/dozken/aoc23/pkg/day4"
	"github.com/dozken/aoc23/pkg/util"
)

func TestP1_sample(t *testing.T) {
	got := util.SumFromFile("p1.sample", day4.Part1)
	expected := 13
	if got != expected {
		t.Errorf("P1() = %v, want %v", got, expected)
	}
}

func TestP1_data(t *testing.T) {
    got := util.SumFromFile("p1.data", day4.Part1)
    expected := 21088
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}

func TestP2_sample(t *testing.T) {
	got := day4.SumFromFile("p1.sample")
	expected := 30
	if got != expected {
		t.Errorf("P2() = %v, want %v", got, expected)
	}
}

func TestP2_data(t *testing.T) {
    got := day4.SumFromFile("p1.data")
    expected := 6874754
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}
