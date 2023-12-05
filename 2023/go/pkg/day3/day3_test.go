package day3_test

import (
	"testing"
	"github.com/dozken/aoc23/pkg/day3"
)

func TestP1_sample(t *testing.T) {
	got := day3.SumFromFile("p1.sample", day3.Part1)
	expected := 4361
	if got != expected {
		t.Errorf("P1() = %v, want %v", got, expected)
	}
}

func TestP1_data(t *testing.T) {
    got := day3.SumFromFile("p1.data", day3.Part1)
    expected := 520019
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}

func TestP2_sample(t *testing.T) {
	got := day3.SumFromFile("p1.sample", day3.Part2)
	expected := 467835
	if got != expected {
		t.Errorf("P2() = %v, want %v", got, expected)
	}
}

func TestP2_data(t *testing.T) {
    got := day3.SumFromFile("p1.data", day3.Part2)
    expected := 75519888
    if got != expected {
        t.Errorf("P1() = %v, want %v", got, expected)
    }
}
