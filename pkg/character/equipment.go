package character

type Equipment struct {
	Gear       []Gear
	Ammunition int
	Cash       int
}

type Gear struct {
	Name string
	Note string
}
