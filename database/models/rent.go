package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Rent struct {
	gorm.Model
	ID            uint
	BookID        uint
	Book          Book `gorm:"foreignKey:BookID"`
	ReaderID      uint
	Reader        Reader `gorm:"foreignKey:ReaderID"`
	RentalTime    time.Time
	RentalPeriod  time.Time
	AmountPenalty uint
}

func (receiver *Rent) GetById(db *gorm.DB, id int) {
	err := db.Model(&Rent{}).Preload("Books").Preload("Readers").Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln(err)
	}
}
