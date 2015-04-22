package main

import (
	"fmt"
	"regexp"
)

const (
	ExerciseNameRegexp = `^[A-Z]+\.\s[A-Za-z\s]+$`
	SetDataRegexp      = `^[ivx]+\.\s[0-9]+x[0-9]+\s@[0-9]+$`
	SplitSetDataRegexp = `\s`
)

func PrintTonnage(data []string) {
	var output []string
	for i := range data {
		if isExerciseName(data[i]) {
			output = append(output, fmt.Sprintf("%s", data[i]))
		} else if isSetData(data[i]) {
			setParts := splitSetData(data[i])
			output = append(output, fmt.Sprintf("\t%s %s", setParts[1], setParts[2]))
		}
	}

	output = append(output, "\nTonnage:")
	for i := range data {
		if isExerciseName(data[i]) {
			// split string and print exercise name
		}
	}

	for i := range output {
		fmt.Println(output[i])
	}
}

func isExerciseName(s string) bool {
	re := regexp.MustCompile(ExerciseNameRegexp)
	return re.MatchString(s)
}

func isSetData(s string) bool {
	re := regexp.MustCompile(SetDataRegexp)
	return re.MatchString(s)
}

func splitSetData(s string) []string {
	return regexp.MustCompile(SplitSetDataRegexp).Split(s, -1)
}
