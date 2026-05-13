// Package items defines structs and methods for creating and using items in the users Bag
package items

type Ball struct {
	Name       string
	Multiplier float64
}

func GetAvailableBalls() map[string]Ball {
	return map[string]Ball{
		"pokeball": {
			Name:       "pokeball",
			Multiplier: 1.0,
		},
		"greatball": {
			Name:       "pgreatball",
			Multiplier: 1.5,
		},
		"ultraball": {
			Name:       "ultraball",
			Multiplier: 2.0,
		},
		"masterball": {
			Name:       "masterball",
			Multiplier: 100.0,
		},
	}
}
