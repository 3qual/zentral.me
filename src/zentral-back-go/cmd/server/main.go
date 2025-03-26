package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/3qual/zentral-back-go/common/db"
)

func main() {
	// 1. Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found or can't be loaded: %v", err)
	}

	// 2. Инициализируем соединение с БД
	dbConn := db.ConnectDB()
	// Если в пакете db используется глобальная переменная + sync.Once,
	// важно вызвать ConnectDB() ровно один раз.

	// 3. Закрываем соединение при завершении
	defer db.CloseDB()

	// 4. Получаем порт из окружения или берем 8080
	port := os.Getenv("GO_APP_INTERNAL_PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	// 5. Создаем роутер и передаём в него dbConn
	r := newRouter(dbConn)

	// 6. Создаем http-сервер
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 7. Запускаем сервер в горутине
	go func() {
		log.Printf("Server is running on %s\n", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	// 8. Ожидаем сигнал для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server exited properly")
}

// newRouter инициализирует chi.Router и принимает gorm.DB для работы в хендлерах
func newRouter(dbConn *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World! Fuck this..."))
	})

	// Пример эндпоинта, проверяющего доступность БД
	r.Get("/ping-db", func(w http.ResponseWriter, r *http.Request) {
		sqlDB, err := dbConn.DB()
		if err != nil {
			http.Error(w, "Error get DB instance", http.StatusInternalServerError)
			return
		}
		if err := sqlDB.Ping(); err != nil {
			http.Error(w, "Database not reachable", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Database is alive!"))
	})

	return r
}
