package character

type Armor struct {
	Head   ArmorField
	Body   ArmorField
	Shield ArmorField
}

type ArmorField struct {
	SP      int
	Penalty int
}
