package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Test function to load and print monster data
func TestLoadMonsters(t *testing.T) {
	file, err := os.Open("5e-SRD-Monsters.json")
	if err != nil {
		t.Fatalf("Failed to open monsters.json: %v", err)
	}
	defer file.Close()

	// Decode the JSON into a slice of maps
	var rawMonsters []map[string]interface{}
	if err := json.NewDecoder(file).Decode(&rawMonsters); err != nil {
		panic(err)
	}

	// Convert each monster to the flat Monster struct
	monsters := make([]Monster, len(rawMonsters))
	for i, data := range rawMonsters {
		monsters[i] = parseMonster(data)
	}

	// Check if we loaded at least one monster
	if len(monsters) == 0 {
		t.Fatal("No monsters found in the JSON file")
	}

	// Print the first monster’s name as a test
	fmt.Println("First monster:", monsters[0].Name)
}
