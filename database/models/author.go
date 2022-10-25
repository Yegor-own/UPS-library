package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Author struct {
	gorm.Model
	ID             uint
	Name           string
	PhotoLink      string
	DOB            time.Time
	DOD            time.Time
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
