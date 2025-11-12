package bastion

import "de.nilskrau.dndbot/calendar"

type SpecialFacility struct {
	Level        uint
	Name         string
	Prerequisite string
	Order        string
}

type Bastion struct {
	SpecialFacilities []SpecialFacility
	StartDate         calendar.Date
}
