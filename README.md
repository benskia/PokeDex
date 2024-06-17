# PokeDex

A CLI project for exploring Pokemon locations, capturing monsters, and inspecting their stats.

## How to Install

With Go installed, either ```go build``` within './cmd/pokedex/' or ```go build -o bin/pokedex cmd/pokedex/main.go``` from within the projects root directory.

Place and run the 'pokedex' executable wherever you like.

## Commands

* help : Displays a list of available commands.
```
❯ cd cmd/pokedex
❯ go build && ./pokedex

Welcome to the PokeDex!

PokeDex > help

Available commands:

 - Name: explore
 - Description: Display which Pokemon can be found at a location.

 - Name: catch
 - Description: Attempt to catch a Pokemon, by name.

 - Name: inspect
 - Description: Display a Pokemon's details, by name.

 - Name: pokedex
 - Description: Display the names of all the Pokemon you've caught.

 - Name: help
 - Description: Display a list of available commands.

 - Name: exit
 - Description: Exit the Pokedex CLI.

 - Name: map
 - Description: Display the next 20 Pokemon locations.

 - Name: mapb
 - Description: Display the previous 20 Pokemon locations.

PokeDex >
```
* exit : Ends the program.
```
PokeDex > exit
Shutting down PokeDex...
```
* map : Displays the next 20 locations.
* mapb : Displays the previous 20 locations.
```
PokeDex >  map

canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior

PokeDex >
```
* explore area-name : Displays the Pokemon that can be encountered at area-name.
```

PokeDex > explore canalave-city-area

tentacool
tentacruel
staryu
magikarp
gyarados
wingull
pelipper
shellos
gastrodon
finneon
lumineon

PokeDex >
```
* catch pokemon-name: Attempts to catch the Pokemon, by pokemon-name.
```
PokeDex > catch tentacool

Throwing a Pokeball at tentacool...

Failed to capture tentacool!

PokeDex >
```
* inspect pokemon-name : Displays details about the Pokemon, by pokemon-name.
```
PokeDex > inspect rattata

Name:  rattata
Height:  3 decimeters
Weight:  35 hectograms
Stats:
  -hp:  30
  -attack:  56
  -defense:  35
  -special-attack:  25
  -special-defense:  35
  -speed:  72
Types:
  - normal

PokeDex >
```
* pokedex : Displays a list of all Pokemon caught so far.
```
PokeDex > pokedex

Your Pokedex:
  - rattata
  - tentacool
  - noctowl
  - wooper

PokeDex >
```
