package main

import (
	"cpr-interactive/pkg/character/roles"
	"cpr-interactive/pkg/misc"
	"fmt"
)

func main() {
	var role = roles.MedTech
	fmt.Println(role.GetDescription(misc.Russian))
}
