package main

import "fmt"

func main() {
	sourceMap := make(map[string]int, 8)
	sourceMap["张1"] = 90
	sourceMap["明1"] = 100
	sourceMap["张2"] = 90
	sourceMap["明2"] = 100
	sourceMap["张3"] = 90
	sourceMap["明3"] = 100
	sourceMap["张4"] = 90
	sourceMap["明4"] = 100
	sourceMap["张5"] = 90
	sourceMap["明5"] = 100
	sourceMap["张6"] = 90
	sourceMap["明6"] = 100
	sourceMap["张7"] = 90
	sourceMap["明7"] = 1000000000000000000
	fmt.Println(sourceMap)
	fmt.Println(sourceMap["明1"])
	fmt.Printf("type of a:%T\n", sourceMap)
}
