package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database structure
type DbInstance struct {
	Db *gorm.DB
}

var (
	_ = godotenv.Load(".env")
	//DB variable
	DB DbInstance
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	var err error
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
		//fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	// db.DB().SetMaxIdleConns(10)
	// DB.LogMode(true)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	DB = DbInstance{Db: db}

	// return DB

}
