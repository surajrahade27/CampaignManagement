package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model

	ID           int    `gorm:"primaryKey" json:"Id"`
	Name         string `json:"Name"`
	Type         string `json:"Type"`
	Location     string `json:"Location"`
	Is_pre_order bool   `json:"Is_pre_order"`
}
