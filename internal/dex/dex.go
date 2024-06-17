package dex

type Pokedex struct {
	Dex map[string]Pokemon
}

type Pokemon struct {
	Name string
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Dex: make(map[string]Pokemon),
	}
}
