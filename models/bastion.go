package models

import (
	"de.nilskrau.dndbot/bastion"
	"gorm.io/gorm"
)

type Bastion struct {
	gorm.Model

	SpecialFacilities []*bastion.SpecialFacility
}
