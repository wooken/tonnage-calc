package main

import (
	"testing"
)

func TestWriteExerciseData(t *testing.T) {
	var e Exercises
	d := []string{"A. Push Press", "i. 4x5 @95", "i. 4x5 @95", "i. 4x5 @95"}
	e.WriteExerciseData(d)
	if len(e) != 1 {
		t.Fail()
	}
	if len(e[0].Data) != 3 {
		t.Fail()
	}
}

func TestCalcTonnage(t *testing.T) {
	vals := []int{5, 3, 115}
	if calcTonnage(vals) != 1725 {
		t.Fail()
	}
}

func TestExtractIntegers(t *testing.T) {
	s := "i. 5x3 @115"
	vals := extractIntegers(s)
	if vals[0] != 5 ||
		vals[1] != 3 ||
		vals[2] != 115 {
		t.Fail()
	}
}

func TestIsExerciseName(t *testing.T) {
	name := "A. Push Press"
	data := "i. 4x5 @95"
	if isExerciseName(name) == false {
		t.Fail()
	}
	if isExerciseName(data) == true {
		t.Fail()
	}
}

func TestIsSetData(t *testing.T) {
	name := "A. Push Press"
	data := "i. 4x5 @95"
	if isSetData(name) == true {
		t.Fail()
	}
	if isSetData(data) == false {
		t.Fail()
	}
}
