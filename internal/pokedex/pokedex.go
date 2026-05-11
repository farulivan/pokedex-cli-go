package pokedex

import "github.com/farulivan/pokedex-cli-go/internal/pokeapi"

func NewPokedex() *Pokedex {
	return &Pokedex{
		entries: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.entries[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, bool) {
	pokemon, ok := p.entries[name]
	return pokemon, ok
}
