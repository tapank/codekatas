package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	m := score - 10
	if m%2 != 0 {
		m--
	}
	return m / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(16) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Ability(),
		Ability(),
		constitution,
		Ability(),
		Ability(),
		Ability(),
		Modifier(constitution) + 10,
	}
}
