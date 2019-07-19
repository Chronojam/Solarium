package namegenerator

import "math/rand"

var (
	Prefixes = []string{
		"Talons",
		"Scarabs",
		"Winds",
		"Dooming",
		"Pipedream",
		"Lemons",
		"Charming",
		"Foobaz",
	}
	Suffixes = []string{
		"Hi'reek",
		"Salnoth Saar",
		"Winter",
		"Chaos",
		"Nyarlathotep",
		"Drelnoch",
		"Aphid",
		"Chrono",
	}
)

func GenerateNew() string {
	Prefix := Prefixes[rand.Intn(len(Prefixes)-1)]
	Suffix := Suffixes[rand.Intn(len(Suffixes)-1)]

	return "The " + Prefix + " Of " + Suffix
}
