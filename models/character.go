package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model

	ServerID string

	PlayerID uint
	Player Player

	Name string
	Image string
	Handle string
	Default bool
}

