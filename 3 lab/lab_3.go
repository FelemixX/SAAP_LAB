package main

import (
	"fmt"
	"time"
)

const savePath = `D:\Study\Magistracy\System Administration and Programming\Labs\1\3 lab\`

// Player represents a chess player.
type Player struct {
	name string
}

// Game represents a chess game.
type Game struct {
	// players is a list of Players.
	players []Player `json:"players"`

	// moves is a list of Moves.
	moves []Move `json:"moves"`

	// startTime is the time the game started.
	startTime time.Time `json:"startTime"`

	// location is the location of the game.
	location string `json:"location"`
}

// Move represents a move in a chess game.
type Move struct {
	// from is the square where the piece was captured.
	from string `json:"from"`

	// to is the square the piece was moved to.
	to string `json:"to"`

	// piece is the type of piece that was moved.
	piece string `json:"piece"`

	player Player `json:"player"`
}

// NewGame creates a new Game object.
func NewGame() *Game {
	return &Game{
		moves:     []Move{},
		startTime: time.Now(),
		location:  "home",
	}
}

// AddMove adds a move to a game.
func (g *Game) AddMove(player Player, from, to, piece string) {
	g.moves = append(g.moves, Move{from, to, piece, player})
}

// String returns a string representation of a Game.
func (g Game) String() string {
	var movesStr string

	for _, move := range g.moves {
		movesStr += fmt.Sprintf("%s %s %s\n", move.from, move.to, move.piece)
	}

	return fmt.Sprintf("Game started at %v in %s.\n%s\n%s", g.startTime, g.location, g.players[0].name, movesStr)
}

// String returns a string representation of a Game.
func (g *Game) GetMovesJson() string {
	return fmt.Sprint("%+v", g)
}

func (g Game) SaveToFile() {
	if len(g.moves) == 0 {
		return
	}
	println(g.GetMovesJson())
	/*gameState := g.GetMovesJson()
	fileName := savePath + time.Now().Format("01-02-2006") + "_output.txt"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = io.WriteString(file, gameState)
	if err != nil {
		log.Fatal(err)
	}*/
}

func LoadFromFile(fileName string) {

}

func main() {
	// Create a new game.
	game := NewGame() //1

	// Add some moves.
	game.AddMove(Player{`White`}, "a1", "b2", "pawn")
	game.AddMove(Player{`Black`}, "c3", "d4", "rook")

	game.SaveToFile()
}
