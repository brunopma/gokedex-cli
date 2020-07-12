# gokedex-cli : pokedex CLI written in Golang fetching from pokeapi.co. GOtta Catch 'Em All!

## Overview [![GoDoc](https://godoc.org/github.com/brunopma/gokedex-cli?status.svg)](https://godoc.org/github.com/brunopma/gokedex-cli)

## Install

```
go get github.com/brunopma/gokedex-cli
```

## Build

After installing the package, just go to the project root and write:
```
go build .
```

## Download
Pre-built binary file for UNIX of this project can be downloaded from [Releases page](https://github.com/brunopma/gokedex-cli/releases)

## Usage:

* Fetching pokemon by name:
```
$ ./gokedex-cli -n charmander

#4
Name: Charmander
Height: 0.6m
Weight: 8.5Kg
Types: fire
Abilities: blaze / solar-power
```
* Fetching pokemon by id:
```
$ ./gokedex-cli -i 91

#91
Name: Cloyster
Height: 1.5m
Weight: 132.5Kg
Types: water / ice
Abilities: shell-armor / skill-link / overcoat
```
* Displaying result as JSON:
```
$ ./gokedex-cli -n bulbasaur -o json

{"abilities":[{"ability":{"name":"overgrow","url":"https://pokeapi.co/api/v2/ability/65/"},"is_hidden":false,"slot":1},{"ability":{"name":"chlorophyll","url":"https://pokeapi.co/api/v2/ability/34/"},"is_hidden":true,"slot":3}],"height":7,"id":1,"name":"bulbasaur","types":[{"slot":1,"type":{"name":"grass","url":"https://pokeapi.co/api/v2/type/12/"}},{"slot":2,"type":{"name":"poison","url":"https://pokeapi.co/api/v2/type/4/"}}],"weight":69}
```
* Asking for Help:
```
$ ./gokedex-cli --help

Usage of ./gokedex-cli:
  -i int
        Pokemon ID to fetch, currently there's 807 available pokemon!
  -n string
        Pokemon ID to fetch (default "text")
  -o string
        Print output in format: text/json (default "text")
  -t int
        Client timeout in seconds (default 30)
```