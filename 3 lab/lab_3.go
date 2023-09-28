package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const savePath = `C:\Users\bythe\GolandProjects\SAAP_LAB\3 lab\dump.json`
const testPath = `C:\Users\bythe\GolandProjects\SAAP_LAB\3 lab\dump2.json`

// Player represents a chess player.
type Player struct {
	Name string `json:"name"`
}

// Game represents a chess game.
type Game struct {
	// players is a list of Players.
	Players []Player `json:"players"`

	// moves is a list of Moves.
	Moves []Move `json:"moves"`

	// startTime is the time the game started.
	StartTime time.Time `json:"startTime"`

	// location is the location of the game.
	Location string `json:"location"`
}

// Move represents a move in a chess game.
type Move struct {
	// from is the square where the piece was captured.
	From string `json:"from"`

	// to is the square the piece was moved to.
	To string `json:"to"`

	// piece is the type of piece that was moved.
	Piece string `json:"piece"`

	Player Player `json:"player"`
}

// NewGame creates a new Game object.
func NewGame() *Game {
	return &Game{
		Moves:     []Move{},
		StartTime: time.Now(),
		Location:  "home",
	}
}

// AddMove adds a move to a game.
func (g *Game) AddMove(player Player, from, to, piece string) {
	g.Moves = append(g.Moves, Move{from, to, piece, player})
}

// String returns a string representation of a Game.
func (g *Game) String() string {
	var movesStr string

	for _, move := range g.Moves {
		movesStr += fmt.Sprintf("%s %s %s\n", move.From, move.To, move.Piece)
	}

	return fmt.Sprintf("Game started at %v in %s.\n%s\n%s", g.StartTime, g.Location, g.Players[0].Name, movesStr)
}

func (g *Game) SaveToFile(path string) {
	data, _ := json.MarshalIndent(g, "", " ")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	if _, err := f.Write(data); err != nil {
		log.Println(err)
	}
}

func LoadFromFile(path string) *Game {
	game := NewGame()
	data, _ := os.ReadFile(path)

	err := json.Unmarshal(data, game)

	if err != nil {
		log.Println(err)
	}

	return game
}

func main() {
	// Create a new game.
	game := NewGame() //1

	// Add some moves.
	game.AddMove(Player{`White`}, "a1", "b2", "pawn")
	game.AddMove(Player{`Black`}, "c3", "d4", "rook")

	game.SaveToFile(savePath)

	loadedGame := LoadFromFile(savePath)
	loadedGame.SaveToFile(testPath)
}
