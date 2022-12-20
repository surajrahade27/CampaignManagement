package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ID           int     `gorm:"primaryKey" json:"id"`
	Descripstion string  `json:"Descripstion"`
	Sku_number   uint32  `json:"Sku_number"`
	Name         string  `json:"Name"`
	Category     string  `json:"Category"`
	Unit_price   float32 `json:"Unit_price"`
	Lead_time    int     `json:"Lead_time"`
}
