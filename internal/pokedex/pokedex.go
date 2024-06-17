package pokedex

type Pokemon struct {
	Name string
}

func NewPokedex() *map[string]Pokemon {
	return new(map[string]Pokemon)
}
