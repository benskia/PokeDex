package dex

import "fmt"

type Pokedex struct {
	Dex map[string]Pokemon
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  struct {
		HP             int
		Attack         int
		Defense        int
		SpecialAttack  int
		SpecialDefense int
		Speed          int
	}
	Types []string
}

func (p Pokemon) PrintDetails() {
	fmt.Println()
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Height: %v decimeters\n", p.Height)
	fmt.Printf("Weight: %v hectograms\n", p.Weight)
	fmt.Print("Stats:\n")
	fmt.Printf("\t-hp: %v\n", p.Stats.HP)
	fmt.Printf("\t-attack: %v\n", p.Stats.Attack)
	fmt.Printf("\t-defense: %v\n", p.Stats.Defense)
	fmt.Printf("\t-special-attack: %v\n", p.Stats.SpecialAttack)
	fmt.Printf("\t-special-defense: %v\n", p.Stats.SpecialDefense)
	fmt.Printf("\t-speed: %v\n", p.Stats.Speed)
	fmt.Print("Types:\n")
	for _, t := range p.Types {
		fmt.Printf("\t- %v\n", t)
	}
	fmt.Println()
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Dex: make(map[string]Pokemon),
	}
}
