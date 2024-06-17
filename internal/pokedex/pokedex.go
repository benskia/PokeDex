package pokedex

type Pokedex struct {
	dex map[string]Pokemon
}

type Pokemon struct {
	Name string
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		dex: make(map[string]Pokemon),
	}
}
