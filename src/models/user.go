package models

import "comiditapp/api/src/enums"

type User struct {
	Role           enums.Role `json:"role"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	HashedPassword string     `json:"hashedPassword"`
	Phone          int        `json:"phone"`
	Address        []Address  `json:"address"`
	Menu           []Product  `json:"menu,omitempty"`
}
