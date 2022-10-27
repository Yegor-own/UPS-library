package models

import (
	"gorm.io/gorm"
	"log"
)

type Book struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	Title           string
	MaxRentalPeriod string
	CoverLink       string
	AuthorID        int
	Author          Author `gorm:"foreignKey:AuthorID"`
}

func (receiver *Book) GetById(db *gorm.DB, id int) {
	err := db.Model(&Book{}).Preload("Author").Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln(err)
	}
}
