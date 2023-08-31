package character

import "cpr-interactive/pkg/character/roles"

type Character struct {
	Name     string
	Role     roles.Role
	Portrait string //file name in assets/portraits
	Stats    Stats
	Humanity int
	Health   Health
	//Skills    []Skill
	Armor   Armor
	Weapons []Weapon
	//CyberWare []CyberWare
	Equipment Equipment
	//LifePath  LifePath
	Notes string
}
