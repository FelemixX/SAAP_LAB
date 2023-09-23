package main

import (
	"fmt"
	"time"
)

// Player represents a chess player.
type Player struct {
	name string
}

// Game represents a chess game.
type Game struct {
	// players is a list of Players.
	players []Player

	// moves is a list of Moves.
	moves []Move

	// startTime is the time the game started.
	startTime time.Time

	// location is the location of the game.
	location string
}

// Move represents a move in a chess game.
type Move struct {
	// from is the square where the piece was captured.
	from string

	// to is the square the piece was moved to.
	to string

	// piece is the type of piece that was moved.
	piece string
}

// NewGame creates a new Game object.
func NewGame() *Game {
	return &Game{
		players:   []Player{{"White"}, {"Black"}},
		moves:     []Move{},
		startTime: time.Now(),
		location:  "home",
	}
}

// AddMove adds a move to a game.
func (g *Game) AddMove(from, to, piece string) {
	g.moves = append(g.moves, Move{from, to, piece})
}

// String returns a string representation of a Game.
func (g Game) String() string {
	var movesStr string

	for _, move := range g.moves {
		movesStr += fmt.Sprintf("%s %s %s\n", move.from, move.to, move.piece)
	}

	return fmt.Sprintf("Game started at %v in %s.\n%s\n%s", g.startTime, g.location, g.players[0].name, movesStr)
}

func main() {
	// Create a new game.
	game := NewGame()

	// Add some moves.
	game.AddMove("a1", "b2", "pawn")
	game.AddMove("c3", "d4", "rook")
}
