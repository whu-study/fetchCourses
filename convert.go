package main

import (
	"encoding/json"
	"os"
)

func convert() {
	file, _ := os.ReadFile("general_details.json")
	var oneJson []OneJson
	json.Unmarshal(file, &oneJson)

	println(len(oneJson))
}
