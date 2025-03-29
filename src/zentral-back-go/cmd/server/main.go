package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/3qual/zentral-back-go/app"
	"github.com/3qual/zentral-back-go/common/db"
)

func main() {
	// Инициализируем приложение (включая коннект к БД и миграции)
	userHandler, transactionHandler, folderHandler, folderTransactionHandler, collaboratorHandler, accessTokenHandler, refreshTokenHandler, authHandler, imageHandler, sessionHandler := app.InitializeApp()

	// Закрываем базу данных по завершении
	defer db.CloseDB()

	// Создаем основной роутер, подключая все обработчики
	r := app.NewRouter(
		userHandler,
		transactionHandler,
		folderHandler,
		folderTransactionHandler,
		collaboratorHandler,
		accessTokenHandler,
		refreshTokenHandler,
		authHandler,
		imageHandler,
		sessionHandler,
	)

	// Пример минимального хендлера
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Zentral!"))
	})

	// Читаем порт из переменной окружения или берем 8080
	port := os.Getenv("GO_APP_INTERNAL_PORT")
	if port == "" {
		port = "8080"
	}

	// Настраиваем HTTP-сервер
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("Server is running on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Ожидаем сигнала для корректной остановки
	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	// Канал для сигналов
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server exited properly")
}
