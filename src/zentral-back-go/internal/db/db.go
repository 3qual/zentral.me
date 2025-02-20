package db

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	once     sync.Once
)

// ConnectDB инициализирует соединение с БД (ленивая инициализация).
func ConnectDB() *gorm.DB {
	once.Do(func() {
		connectionString := os.Getenv("DATABASE_URL")
		if connectionString == "" {
			log.Fatal("DATABASE_URL is not set")
		}

		// Настройка уровня логирования
		var logLevel logger.LogLevel
		switch os.Getenv("DB_LOG_LEVEL") {
		case "silent":
			logLevel = logger.Silent
		case "error":
			logLevel = logger.Error
		case "warn":
			logLevel = logger.Warn
		default:
			logLevel = logger.Info
		}

		// Подключение к базе данных
		var err error
		database, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// Настройки пула соединений
		sqlDB, err := database.DB()
		if err != nil {
			log.Fatalf("Failed to access underlying SQL DB: %v", err)
		}

		sqlDB.SetMaxOpenConns(25)                 // Максимальное количество открытых соединений
		sqlDB.SetMaxIdleConns(10)                 // Максимальное количество соединений в простое
		sqlDB.SetConnMaxLifetime(5 * time.Minute) // Время жизни соединения

		log.Println("Database successfully connected")
	})

	return database
}

// CloseDB закрывает соединение с базой данных
func CloseDB() {
	if database != nil {
		sqlDB, err := database.DB()
		if err != nil {
			log.Printf("Error getting DB instance: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
