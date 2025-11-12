package models

import (
  "gorm.io/gorm"
)

type Player struct {
	gorm.Model
	DiscordId string
}
