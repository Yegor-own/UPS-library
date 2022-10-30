package graph

import (
	"github.com/Yegor-own/ghqllibrary/database"
	"gorm.io/gorm"
	"log"
)

var DBLib *gorm.DB

func init() {
	DBLib = database.ConnectDB("host=localhost user=postgres password=root dbname=library port=5432 sslmode=disable")
	
	database.MigrateDB(DBLib)

	log.Println("Generating BooksAndAuthors")
	authors := database.GenerateBooksAndAuthors()
	for _, author := range authors {
		DBLib.Create(&author)
	}
	log.Println("Generate successful")

	log.Println("Generating ReadersAndLeases")
	readers := database.GenerateReadersAndLeases(10, DBLib)
	for _, reader := range readers {
		DBLib.Create(&reader)
	}
	log.Println("Generate successful")

}
