package aoc02

import (
	"testing"
)

func TestGetPositionAndDepthSample(t *testing.T) {
	position, depth := getPositionAndDepth("sample")

	expectedPosition := 15
	if position != expectedPosition {
		t.Fatalf(`position = %d, want "%d"`, position, expectedPosition)
	}

	expectedDepth := 10
	if depth != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, depth, expectedDepth)
	}

	multipled := position * depth
	expectedMult := 150
	if multipled != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, multipled, expectedMult)
	}
}

func TestGetPositionAndDepthInput(t *testing.T) {
	position, depth := getPositionAndDepth("input")

	expectedPosition := 1905
	if position != expectedPosition {
		t.Fatalf(`position = %d, want "%d"`, position, expectedPosition)
	}

	expectedDepth := 907
	if depth != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, depth, expectedDepth)
	}

	multipled := position * depth
	expectedMult := 1727835
	if multipled != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, multipled, expectedMult)
	}
}
