package gameBoard

import(
	dice    "SnLd/Dice"
	player  "SnLd/Player"
	jumper  "SnLd/Jumper"
	
	

)


type GameBoard struct{
	Dice dice.Dice
	NextTurn []player.Player
	Snakes []jumper.Jumper
	Ladders []jumper.Jumper
	BoardSize int
	Map map[string]int 

}

func NewGameboard(
	d dice.Dice,
	nextTurn []player.Player,
	snakes []jumper.Jumper,
	ladders []jumper.Jumper,
	boardSize int,
	m map[string]int,
) *GameBoard {

	return &GameBoard{
		Dice:      d,
		NextTurn:  nextTurn,
		Snakes:    snakes,
		Ladders:   ladders,
		BoardSize: boardSize,
		Map:       m,
	}
}


func (gb * GameBoard)StartGame(){

//check whose turn is it 

//roll a dice 

// Check if > board size-> dont move 

// check for collision -> move position accordingly

// end -> win



// change turn 




//repeat loop 



//exit in case any one won


}

