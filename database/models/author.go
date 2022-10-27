package models

import (
	"gorm.io/gorm"
	"log"
)

type Author struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Name           string
	PhotoLink      string
	DOB            string
	DOD            string
	Books          []Book
	AvailableBooks int
}

func (receiver *Author) GetById(db *gorm.DB, id int) {
	err := db.Model(&Author{}).Preload("Books").Where("id = ?", id).First(&receiver).Error
	if err != nil {
		log.Fatalln(err)
	}
}

//func (receiver *Author) GetAll(db *gorm.DB)  {
//	var authors []Author
//	err := db.Model(&Author{}).Preload("Books").Find(&authors).Error
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
