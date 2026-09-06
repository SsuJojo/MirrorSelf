package main

import (
	"strings"
	"testing"
)

func TestValidateMealTrimsWhitespace(t *testing.T) {
	got, err := validateMeal("  炒饭  ")
	if err != nil {
		t.Fatalf("validateMeal returned an unexpected error: %v", err)
	}
	if got != "炒饭" {
		t.Fatalf("validateMeal returned %q, want %q", got, "炒饭")
	}
}

func TestValidateMealRejectsEmptyInput(t *testing.T) {
	if _, err := validateMeal(" \n\t "); err == nil {
		t.Fatal("validateMeal accepted empty input")
	}
}

func TestValidateMealRejectsOversizedInput(t *testing.T) {
	if _, err := validateMeal(strings.Repeat("菜", maxMealLength+1)); err == nil {
		t.Fatalf("validateMeal accepted input longer than %d characters", maxMealLength)
	}
}

func TestValidateMealCountsUnicodeCharacters(t *testing.T) {
	meal := strings.Repeat("饭", maxMealLength)
	if _, err := validateMeal(meal); err != nil {
		t.Fatalf("validateMeal rejected %d-character Unicode input: %v", maxMealLength, err)
	}
}
