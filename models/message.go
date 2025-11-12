package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	PlayerID uint
	Player Player
	CharacterID uint
	Character Character
	Message string
}
