package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Monster struct definition with all fields as strings
type Monster struct {
	Index                 string `json:"index"`
	Name                  string `json:"name"`
	Size                  string `json:"size"`
	Type                  string `json:"type"`
	Alignment             string `json:"alignment"`
	ArmorClass            string `json:"armor_class"`
	HitPoints             string `json:"hit_points"`
	HitDice               string `json:"hit_dice"`
	HitPointsRoll         string `json:"hit_points_roll"`
	Speed                 string `json:"speed"`
	Actions               string `json:"actions"`
	Strength              string `json:"strength"`
	Dexterity             string `json:"dexterity"`
	Constitution          string `json:"constitution"`
	Intelligence          string `json:"intelligence"`
	Wisdom                string `json:"wisdom"`
	Charisma              string `json:"charisma"`
	DamageVulnerabilities string `json:"damage_vulnerabilities"`
	DamageResistances     string `json:"damage_resistances"`
	DamageImmunities      string `json:"damage_immunities"`
	ConditionImmunities   string `json:"condition_immunities"`
	Senses                string `json:"senses"`
	Languages             string `json:"languages"`
	ChallengeRating       string `json:"challenge_rating"`
	XP                    string `json:"xp"`
	SpecialAbilities      string `json:"special_abilities"`
	LegendaryActions      string `json:"legendary_actions"`
	Image                 string `json:"image"`
	URL                   string `json:"url"`
}

// Function to convert any value to a JSON string representation
func convertToString(value interface{}) string {
	// If the value is already a string, return it directly
	if str, ok := value.(string); ok {
		return str
	}

	// Use json.Marshal for non-string values
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%v", value) // Fallback to default string conversion
	}

	// Convert JSON bytes to string
	result := string(bytes)

	// Remove surrounding quotes if present
	if strings.HasPrefix(result, "\"") && strings.HasSuffix(result, "\"") {
		return result[1 : len(result)-1]
	}

	return result
}

// Parse a monster from a map using reflection to assign values.
func parseMonster(data map[string]interface{}) Monster {
	monster := Monster{}
	val := reflect.ValueOf(&monster).Elem()

	// Iterate over the struct fields and assign values from the map.
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		// Get the JSON key name
		fieldName := field.Tag.Get("json")

		// If the field exists in the map, convert and assign it.
		if value, ok := data[fieldName]; ok {
			val.Field(i).SetString(convertToString(value))
		}
	}

	return monster
}
