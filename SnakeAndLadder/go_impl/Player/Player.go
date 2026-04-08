package player



type Player struct{
	PlayerName 	string
	ID 		int


} 




func NewPlayer(id int ,playerN string) *Player{
	return &Player{
		ID:id,
		PlayerName: playerN,

	}

}
