package database_config

import (
	"alterra-agmc-dynamic-crud/app"
	"alterra-agmc-dynamic-crud/models/database"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	application := app.Init()
	maxIdleConn := application.Config.DBMaxIdleConnections
	maxConn := application.Config.DBMaxConnections
	maxLifetimeConn := application.Config.DBMaxLifetimeConnections
	db_user := application.Config.DBUsername
	db_pass := application.Config.DBPassword
	db_host := application.Config.DBHost
	db_port := application.Config.DBPort
	db_database := application.Config.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	InitCreateTable(db)
	log.Println("database connect success")
	return db

}

func InitCreateTable(db *gorm.DB) {

	err := db.Debug().AutoMigrate(
		&database.Book{},
		&database.User{},
	)
	if err != nil {
		log.Fatal("Failed Auto Migrate Database")
	}

}
