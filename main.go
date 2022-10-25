package main

import (
	"github.com/Yegor-own/ghqllibrary/database"
	"gorm.io/gorm"
	"log"
)

var (
	DBLib *gorm.DB
)

func init() {
	DBLib = database.ConnectDB("host=localhost user=postgres password=root dbname=library port=5432 sslmode=disable")
	database.MigrateDB(DBLib)
}

func main() {

	readers := database.GenerateReadersAndLeases(10, DBLib)
	for _, reader := range readers {
		res := DBLib.Create(&readers)
		log.Println(reader.ID, res.RowsAffected)
	}
}
