package graph

import (
	"github.com/Yegor-own/ghqllibrary/database"
	"gorm.io/gorm"
)

var DBLib *gorm.DB

func init() {
	DBLib = database.ConnectDB("host=localhost user=postgres password=root dbname=library port=5432 sslmode=disable")
	database.MigrateDB(DBLib)
}
