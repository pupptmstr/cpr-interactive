package roles

import (
	"cpr-interactive/pkg/misc"
	"encoding/json"
	"fmt"
	"strings"
)

type Role string

const (
	RockerBoy Role = "RockerBoy"
	Solo      Role = "Solo"
	Netrunner Role = "NetRunner"
	Techie    Role = "Techie"
	MedTech   Role = "MedTech"
	Media     Role = "Media"
	Fixer     Role = "Fixer"
	Nomad     Role = "Nomad"
	Cop       Role = "Cop"
	Corporate Role = "Corporate"
)

type RoleDescription struct {
	Role        string `json:"role"`
	Description string `json:"description"`
}

func (role *Role) GetDescription(language misc.Language) string {
	fullPath := fmt.Sprintf("/descriptions/roles-%s.json", language)
	fmt.Println(fullPath)
	content, err := misc.ReadAssetAsBytes(fullPath)
	if err != nil {
		return misc.StandardDescriptionError
	}
	var data []RoleDescription
	err = json.Unmarshal(content, &data)
	if err != nil {
		return misc.StandardDescriptionError
	}

	var description string
	for _, value := range data {
		if strings.ToLower(value.Role) == strings.ToLower(string(*role)) {
			description = value.Description
		}
	}

	if description == "" {
		if language == misc.English {
			return misc.StandardDescriptionError
		} else {
			return role.GetDescription(misc.English)
		}
	} else {
		return description
	}
}

func (role *Role) GetRoleSkill() {
	//TODO("сделать возврат способности класса")
}
