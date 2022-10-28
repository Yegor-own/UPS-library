package models

import (
	"gorm.io/gorm"
	"log"
)

type Rent struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	BookID        int
	Book          Book `gorm:"foreignKey:BookID"`
	ReaderID      int
	Reader        Reader `gorm:"foreignKey:ReaderID"`
	RentalTime    string
	RentalPeriod  string
	AmountPenalty int
}

func (receiver *Rent) GetById(db *gorm.DB, id int) {
	err := db.Model(&Rent{}).Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln("rent > ", err)
	}
}
