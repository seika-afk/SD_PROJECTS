package dice

import "math/rand"

type Dice struct {
    NumberOfDice int
}

func NewDice(num int) Dice {
    return Dice{
        NumberOfDice: num,
    }
}


func (d Dice) RollDice() int{
	total:=0
	for i:=0 ;i<d.NumberOfDice; i++{
		total+=rand.Intn(6)+1
		
	}
	return total
}
 
