package models

import (
	"gorm.io/gorm"
	"log"
)

type Reader struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Name    string
	Email   string
	Leases  []Rent
	LateFee int
}

func (receiver *Reader) GetById(db *gorm.DB, id int) {
	err := db.Model(&Reader{}).Preload("Rents").Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln(err)
	}

}
