package pokeapi

type Pokedex struct {
	Caught map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Caught: make(map[string]Pokemon),
	}
}
