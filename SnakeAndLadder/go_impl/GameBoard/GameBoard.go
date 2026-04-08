package gameBoard

import (
	dice "SnLd/Dice"
	jumper "SnLd/Jumper"
	player "SnLd/Player"
	"fmt"
	"os"
)


type GameBoard struct{
	Dice dice.Dice
	NextTurn []*player.Player
	Snakes []*jumper.Jumper
	Ladders []*jumper.Jumper
	BoardSize int
	Map map[int]int 
	steps int

}

func NewGameboard(
	d dice.Dice,
	nextTurn []*player.Player,
	snakes []*jumper.Jumper,
	ladders []*jumper.Jumper,
	boardSize int,
	m map[int]int,
	s int,
) *GameBoard {

	return &GameBoard{
		Dice:      d,
		NextTurn:  nextTurn,
		Snakes:    snakes,
		Ladders:   ladders,
		BoardSize: boardSize,
		Map:       m,
		steps:     s,
	}
}




///////////////////FOR GAMEBOARD



func (gb *GameBoard) ShowBoard() {
    cellWidth := 6

    // build lookup maps for quick check
    snakeMap := make(map[int]int)
    for _, s := range gb.Snakes {
        snakeMap[s.StartPoint] = s.EndPoint
    }

    ladderMap := make(map[int]int)
    for _, l := range gb.Ladders {
        ladderMap[l.StartPoint] = l.EndPoint
    }

    // print board
    for i := 1; i <= gb.BoardSize; i++ {
        cellContent := fmt.Sprintf("[%d]", i)

        // check snake / ladder
        if _, ok := snakeMap[i]; ok {
            cellContent = fmt.Sprintf("[%dS]", i)
        }
        if _, ok := ladderMap[i]; ok {
            cellContent = fmt.Sprintf("[%dL]", i)
        }

        // check player
        for _, p := range gb.NextTurn {
            if gb.Map[p.ID] == i {
                cellContent = fmt.Sprintf("[P%d]", p.ID)
                break
            }
        }

        fmt.Printf("%-*s", cellWidth, cellContent)
    }

    fmt.Println()
    fmt.Println("------")
}

/////////////////////////////////////// MAIN GAME LOOP

func (gb * GameBoard)StartGame(){
//repeat loop 

	for{
//check whose turn is it 

player_ := gb.steps%len(gb.NextTurn)
curr_player := gb.NextTurn[player_]
currPos := gb.Map[curr_player.ID]
//it have .ID .PlayerName
//roll a dice 
num := gb.Dice.RollDice()

//next position 
nextPos := currPos+num 

// Check if > board size-> dont move 
if nextPos > gb.BoardSize{

nextPos = currPos
}else{

gb.Map[curr_player.ID]=nextPos
}


fmt.Printf("%s rolled  %d and moved to %d\n", curr_player.PlayerName, num, gb.Map[curr_player.ID])


// check for collision -> move position accordingly
//check for snakes and ladders 

//for snake 
for _,item := range  gb.Snakes{
	if item.StartPoint  == gb.Map[curr_player.ID]{
	oldPos := gb.Map[curr_player.ID]

gb.Map[curr_player.ID] = item.EndPoint

fmt.Printf("Player %s fell from %d to %d\n",
    curr_player.PlayerName,
    oldPos,
    item.EndPoint,
)
break 
	}

}



//for ladder  
for _,item := range gb.Ladders{

	if item.StartPoint == gb.Map[curr_player.ID]{
		fmt.Println("Player ",curr_player.PlayerName," entered a Ladder")
	fmt.Println("Player",curr_player.PlayerName," Jumped from ",nextPos ," to ",item.EndPoint)
		gb.Map[curr_player.ID]= item.EndPoint
		break 
	}
}

// end -> win

if gb.Map[curr_player.ID]== gb.BoardSize{
	fmt.Printf("PLAYER : %s WON \n",curr_player.PlayerName)
	
//exit in case any one won
gb.ShowBoard()


	os.Exit(0)
}
gb.ShowBoard()




gb.steps+=1




}
}





