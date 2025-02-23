package pokeapi

import (
	"testing"
)

func TesGetLocationInfo(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "pastoria-city-area",
			expected: []string{"tentacool", "tentacruel", "magikarp", "gyarados", "remoraid", "octillery", "wingull", "pelipper", "shellos", "gastrodon"},
		},
		{
			input:    "canalave-city-area",
			expected: []string{"tentacool", "tentacruel", "staryu", "magikarp", "gyarados", "wingull", "pelipper", "shellos", "gastrodon", "finneon", "lumineon"},
		},
	}

	for _, c := range cases {
		location, _ := GetLocationInfo(c.input)

		var actual []string

		for i, pe := range location.PokemonEncounters {
			actual[i] = pe.Pokemon.Name
		}

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d pokemon in result, got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			pokemon := actual[i]
			expectedPokemon := c.expected[i]
			if pokemon != expectedPokemon {
				t.Errorf("Expected %s, got %s", expectedPokemon, pokemon)
			}
		}
	}
}