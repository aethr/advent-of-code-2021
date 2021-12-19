package aoc02

import (
	"testing"
)

func TestGetPositionAndDepthWithAimSample(t *testing.T) {
	position, depth := getPositionAndDepthUsingAim("sample")

	expectedPosition := 15
	if position != expectedPosition {
		t.Fatalf(`position = %d, want "%d"`, position, expectedPosition)
	}

	expectedDepth := 60
	if depth != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, depth, expectedDepth)
	}

	multipled := position * depth
	expectedMult := 900
	if multipled != expectedDepth {
		t.Fatalf(`multipled = %d, want "%d"`, multipled, expectedMult)
	}
}

func TestGetPositionAndDepthWithAimInput(t *testing.T) {
	position, depth := getPositionAndDepthUsingAim("input")

	expectedPosition := 1905
	if position != expectedPosition {
		t.Fatalf(`position = %d, want "%d"`, position, expectedPosition)
	}

	expectedDepth := 810499
	if depth != expectedDepth {
		t.Fatalf(`depth = %d, want "%d"`, depth, expectedDepth)
	}

	multipled := position * depth
	expectedMult := 1544000595
	if multipled != expectedDepth {
		t.Fatalf(`multipled = %d, want "%d"`, multipled, expectedMult)
	}
}
