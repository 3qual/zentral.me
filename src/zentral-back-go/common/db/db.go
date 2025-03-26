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
		connectionString := os.Getenv("POSTGRES_DATABASEURL")
		if connectionString == "" {
			log.Fatal("POSTGRES_DATABASEURL is not set")
		}

		// Подключение к базе данных
		var err error
		database, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn), // Фиксированный уровень логирования
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
