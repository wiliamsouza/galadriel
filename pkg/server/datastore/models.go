package datastore

import (
	"gorm.io/gorm"
)

//TODO - Remember that the DeletedAt is removed in Spire

type Member struct {
	gorm.Model
	BridgeID uint
	SpiffeID string `gorm:"unique_index"`
}

type Bridge struct {
	gorm.Model
	OrganizationID uint
	Description    string `gorm:"unique_index"`
	Active         bool
	Members        []Member
}

type Organization struct {
	gorm.Model
	Name    string `json:"name" gorm:"unique_index"`
	Bridges []Bridge
}