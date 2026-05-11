# Pokédex CLI (Go)

A command-line interface application built with Go that allows users to explore the Pokémon world, catch Pokémon, and build their own Pokédex using the [PokéAPI](https://pokeapi.co/).

## Features

- **Interactive REPL**: Command-line interface using `bufio.Scanner` for seamless user interaction
- **Location Navigation**: Browse through Pokémon locations with forward/backward pagination
- **Exploration**: Discover wild Pokémon in different areas
- **Catch Mechanics**: Attempt to catch Pokémon with probability-based success rates
- **Pokédex Management**: Store and inspect caught Pokémon with detailed stats
- **Smart Caching**: Built-in cache system with automatic reaping to optimize API calls
- **Concurrency-Safe**: Cache uses `sync.Mutex` for safe concurrent access from a background reap loop
- **Idiomatic Go**: Organized into `internal/` packages with clear separation of concerns

## Tech Stack

- **Language**: Go 1.26+
- **API**: [PokéAPI v2](https://pokeapi.co/api/v2)
- **Standard Library Only**: `net/http`, `encoding/json`, `bufio`, `sync`, `time`, `math/rand/v2`
- **Testing**: Go's built-in `testing` package

## Project Structure

```
.
├── main.go                    # Application entry point
├── repl.go                    # REPL loop, command registry, input cleaning
├── repl_test.go               # REPL unit tests
├── config.go                  # Shared application config/state
├── command_help.go            # `help` command
├── command_exit.go            # `exit` command
├── command_map.go             # `map` / `mapb` commands (pagination)
├── command_explore.go         # `explore` command
├── command_catch.go           # `catch` command
├── command_inspect.go         # `inspect` command
├── command_pokedex.go         # `pokedex` command
└── internal/
    ├── pokeapi/               # PokéAPI HTTP client + typed models
    │   ├── client.go
    │   ├── http.go
    │   ├── location_area.go
    │   ├── pokemon.go
    │   ├── pokeapi.go
    │   └── models.go
    ├── pokecache/             # In-memory cache with TTL reap loop
    │   ├── pokecache.go
    │   ├── models.go
    │   └── pokecache_test.go
    └── pokedex/               # In-memory Pokédex storage
        ├── pokedex.go
        └── models.go
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/farulivan/pokedex-cli-go.git
   cd pokedex-cli-go
   ```

2. (Optional) Build a binary:
   ```bash
   go build -o pokedex-cli-go
   ```

## Usage

Run directly with `go run`:
```bash
go run .
```

Or run the compiled binary:
```bash
./pokedex-cli-go
```

### Available Commands

| Command   | Usage                       | Description                                  |
| --------- | --------------------------- | -------------------------------------------- |
| `help`    | `help`                      | Display all available commands               |
| `map`     | `map`                       | Show the next page of location areas         |
| `mapb`    | `mapb`                      | Show the previous page of location areas     |
| `explore` | `explore <location_name>`   | List all Pokémon in a specific location      |
| `catch`   | `catch <pokemon_name>`      | Attempt to catch a Pokémon                   |
| `inspect` | `inspect <pokemon_name>`    | View details of a caught Pokémon             |
| `pokedex` | `pokedex`                   | List all caught Pokémon                      |
| `exit`    | `exit`                      | Exit the application                         |

### Example Session

```
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
explore <location_name>: Explore the list of pokemon in a location
catch <pokemon_name>: Catch a Pokemon and add it to your Pokedex
inspect <pokemon_name>: Inspect a caught Pokemon's details
pokedex: List all Pokemon in your Pokedex
exit: Exit the Pokedex

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
  -hp: 40
  -attack: 40
  -defense: 35
  -special-attack: 50
  -special-defense: 100
  -speed: 70
Types:
  - water
  - poison

Pokedex > pokedex
Your Pokedex:
 - tentacool
```

## Architecture Highlights

### Configuration & State
A single `config` struct (`config.go`) is threaded through every command callback, holding:
- The `pokeapi.Client`
- Pagination cursors (`nextLocationsURL`, `prevLocationsURL`)
- The user's `*pokedex.Pokedex`

### Command Registry
Commands are registered in `getCommands()` in `repl.go` as a `map[string]cliCommand`, where each entry stores a name, description, and callback of signature `func(cfg *config, args ...string) error`.

### Caching System
The `pokecache.Cache` package implements:
- Time-based expiration with a configurable interval
- A background goroutine (`reapLoop`) that periodically evicts expired entries
- `sync.Mutex`-guarded access for concurrency safety

The HTTP client in `internal/pokeapi` consults the cache before issuing network requests, falling back to PokéAPI on a miss and storing the response body for next time.

### Catch Mechanics
Catch success is determined by a probabilistic check against the Pokémon's `base_experience`:

```go
if rand.IntN(pokemon.BaseExperience) < catchThreshold {
    // caught
}
```

A higher `base_experience` means a wider RNG range, making rarer Pokémon harder to catch. The `catchThreshold` constant lives in `command_catch.go`.

## Testing

Run the full test suite:
```bash
go test ./...
```

Run with verbose output:
```bash
go test -v ./...
```

Current tests cover:
- REPL input cleaning (`repl_test.go`)
- Cache add/get and expiration (`internal/pokecache/pokecache_test.go`)

## Future Enhancements

- **Battle System**: Turn-based battles with type advantages and move sets
- **Expanded Test Coverage**: Unit tests for every command and edge cases
- **Pokémon Party System**: A party of up to 6 Pokémon with level-up mechanics
- **Evolution System**: Allow caught Pokémon to evolve after meeting requirements
- **Persistent Storage**: Save Pokédex progress to disk (JSON / SQLite / BoltDB)
- **Enhanced Navigation**: Replace location names with directional choices (left/right/forward)
- **Random Encounters**: Probability-based wild Pokémon encounters during exploration
- **Poké Ball Variety**: Different ball types (Great Ball, Ultra Ball, Master Ball) with varying catch rates
- **Move Learning**: Allow Pokémon to learn and use moves in battles
- **Trainer Profile**: Track player statistics, badges, and achievements
- **Graceful Shutdown**: Handle `Ctrl+C` to flush state and stop the reap goroutine cleanly

## Development

### Common Commands

- `go run .` — Run the application
- `go build -o pokedex-cli-go` — Build a binary
- `go test ./...` — Run all tests
- `go vet ./...` — Static analysis
- `gofmt -w .` — Format all Go files

## License

[ISC](./LICENSE)

## Repository

[https://github.com/farulivan/pokedex-cli-go](https://github.com/farulivan/pokedex-cli-go)

---

Built with Go and PokéAPI.
