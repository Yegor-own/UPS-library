package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(dsn string) *gorm.DB {

	log.Println("Connecting to db")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("ConnectDB > ", err)
	}
	log.Println("Db connected successful")

	return db
}

func MigrateDB(db *gorm.DB) {
	//modelsList := models.Models{
	//	&models.Author{},
	//	&models.Book{},
	//	&models.Rent{},
	//	&models.Reader{},
	//}
	modelsList := []interface{}{
		&Author{},
		&Book{},
		&Rent{},
		&Reader{},
	}
	log.Println("Migrating...")

	for _, model := range modelsList {
		//log.Println("Migrating ", model)
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalln("MigrateDB > ", err)
		}
		//log.Println(model, " migrated successful")
	}
	log.Println("Migration successful")
}
