package dao

import (
	"consultation-service/models"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DB_PATH = "./data/db/test.db"
)

var (
	Db *gorm.DB
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Init() {
	fmt.Printf(os.Getwd())
	dsn := fmt.Sprintf(DB_PATH + "?_pragma_cipher_page_size=4096")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database:%s", DB_PATH))
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.PatientInfo{})
	db.AutoMigrate(&models.Prescription{})

	// db.Migrator().MigrateColumn(&models.PatientInfo{},)

	Db = db
}
