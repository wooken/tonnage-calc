package main

import (
	"testing"
)

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
