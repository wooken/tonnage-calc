package main

func main() {
	filename := GetArgs()
	data := ParseFile(filename)
	var e Exercises
	e.WriteExerciseData(data)
	e.OutputData()
	e.OutputTonnage()
}
