package character

type Character struct {
	Name       string
	Role       Role
	Portrait   string //file name in assets/portraits
	Parameters Parameters
	Health     int
	Skills     []Skill
	Armor      Armor
	Weapons    []Weapon
	CyberWare  []CyberWare
	Equipment  Equipment
	LifePath   LifePath
	Notes      string
}
