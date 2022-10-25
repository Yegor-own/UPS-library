package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Book struct {
	gorm.Model
	ID              uint
	Title           string
	MaxRentalPeriod time.Time
	CoverLink       string
	AuthorID        uint
	Author          Author `gorm:"foreignKey:AuthorID"`
}

func (receiver *Book) GetById(db *gorm.DB, id int) {
	err := db.Model(&Book{}).Preload("Author").Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln(err)
	}
}
