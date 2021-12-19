package aoc01

import (
	"testing"
)

func TestCountRangeIncreasesSample(t *testing.T) {
	increases := countRangeIncreases("sample", 3)

	if increases != 5 {
		t.Fatalf(`countRangeIncreases = %d, want "5"`, increases)
	}
}

func TestCountRangeIncreasesInput(t *testing.T) {
	increases := countRangeIncreases("input", 3)

	if increases != 1748 {
		t.Fatalf(`countRangeIncreases = %d, want "1748"`, increases)
	}
}
