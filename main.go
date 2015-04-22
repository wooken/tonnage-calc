package main

func main() {
	var filename string
	var data []string
	filename = GetArgs()
	data = ParseFile(filename)
	PrintTonnage(data)
}
