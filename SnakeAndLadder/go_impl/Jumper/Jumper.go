package jumper


type Jumper struct{
	StartPoint   int
	EndPoint     int

}



func NewJumper(st int, ed int) *Jumper{	
	return &Jumper{
		StartPoint : st,
		EndPoint   : ed  

	}



}
