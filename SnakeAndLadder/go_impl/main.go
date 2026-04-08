package main

import (
	"SnLd/Dice"
	"SnLd/GameBoard"
	"SnLd/Jumper"
	"SnLd/Player"
	"fmt"
	"time"
)

//rand.Seed(time.Now().UnixNano())


func main(){
	// import dice -> create Dice 
		d := dice.NewDice(1)	
	//import player-> create player 1 and 2 
		p1 := player.NewPlayer(1,"Player1")
		p2 := player.NewPlayer(2,"Player2")
		
		players := []*player.Player{p1,p2}
		// snakes and ladders 
		s1 := jumper.NewJumper(4,1)
		s2 := jumper.NewJumper(8,3)
		snakes := []*jumper.Jumper{s1,s2}

		l1 := jumper.NewJumper(5,10)	
		ladders := []*jumper.Jumper{l1}

	// current posisiton of players 
		mp := make(map[int]int)
		mp[1]= 0
		mp[2]=0
	
// SLIGHT DELAY ------------------------

fmt.Println("######################################################################")
fmt.Println("                         STARTING GAME ...")
fmt.Println("######################################################################")
time.Sleep(1*time.Second)	

// start gameboard	
		gB := gameBoard.NewGameboard(d,players,snakes,ladders,10,mp,1)
		gB.StartGame()

}
