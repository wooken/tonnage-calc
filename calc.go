package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	ExerciseNameRegexp = `^[A-Z0-9]+\.\s[A-Za-z\s]+$`
	SetDataRegexp      = `^[ivx]+\.\s[0-9]+x[0-9]+\s@[0-9]+$`
	SplitSetDataRegexp = `\s`
)

type ExerciseData struct {
	Name    string
	Data    []string
	Tonnage int
}

type Exercises []ExerciseData

func (e *Exercises) WriteExerciseData(data []string) {
	for i := range data {
		if isExerciseName(data[i]) {
			(*e) = expandExerciseSlice(*e)
			index := len(*e) - 1
			(*e)[index].Name = data[i]
			for j := i + 1; j < len(data); j++ {
				if isExerciseName(data[j]) {
					break
				}
				dataString := fmt.Sprintf("%s", data[j])
				(*e)[index].Data = append((*e)[index].Data, dataString)
				dataInt := extractIntegers(dataString)
				(*e)[index].Tonnage = (*e)[index].Tonnage + calcTonnage(dataInt)
			}
		}
	}
}

func (e *Exercises) OutputData() {
	for i := range *e {
		fmt.Println((*e)[i].Name)
		for j := range (*e)[i].Data {
			t := splitSetData((*e)[i].Data[j])
			fmt.Printf("\t%s %s\n", t[1], t[2])
		}
	}
}

func (e *Exercises) OutputTonnage() {
	fmt.Println()
	fmt.Println("Tonnage:")
	totalTonnage := 0
	for i := range *e {
		name := regexp.MustCompile(`\s`).Split((*e)[i].Name, 2)
		tonnage := (*e)[i].Tonnage
		fmt.Printf("%s: %d\n", name[1], tonnage)
		totalTonnage = totalTonnage + tonnage
	}
	fmt.Printf("Total: %d\n", totalTonnage)
}

func calcTonnage(vals []int) int {
	tonnage := 1
	for i := range vals {
		tonnage = tonnage * vals[i]
	}
	return tonnage
}

func extractIntegers(s string) []int {
	// refactor this, completely unreadable

	vals := make([]int, 3)
	t := splitSetData(s)
	t1 := regexp.MustCompile("x").Split(t[1], -1)
	t2 := regexp.MustCompile("@").Split(t[2], -1)
	vals[0], _ = strconv.Atoi(t1[0])
	vals[1], _ = strconv.Atoi(t1[1])
	vals[2], _ = strconv.Atoi(t2[1])
	return vals
}

func isExerciseName(s string) bool {
	return regexp.MustCompile(ExerciseNameRegexp).MatchString(s)
}

func isSetData(s string) bool {
	return regexp.MustCompile(SetDataRegexp).MatchString(s)
}

func splitSetData(s string) []string {
	return regexp.MustCompile(SplitSetDataRegexp).Split(s, -1)
}

func expandExerciseSlice(e []ExerciseData) []ExerciseData {
	t := make([]ExerciseData, len(e)+1)
	copy(t, e)
	return t
}
