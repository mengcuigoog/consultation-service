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

	db.Migrator().DropColumn(&models.Prescription{}, "created")

	// // Create
	// db.Debug().Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// // db.

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)

	// p := make([]Product, 0)
	// ret := db.Model(&product).Limit(10).Find(&p)
	// fmt.Printf("%+v, err :%v\n", p, ret.Error)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
	// fmt.Printf(">>>>>>>>>>>> db init success")

	Db = db
}
