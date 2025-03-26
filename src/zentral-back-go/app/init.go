package app

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/3qual/zentral-back-go/common/auth"
	"github.com/3qual/zentral-back-go/common/db"
	"github.com/3qual/zentral-back-go/common/migrations"
	"github.com/3qual/zentral-back-go/internal/accesstoken"
	"github.com/3qual/zentral-back-go/internal/collaborator"
	"github.com/3qual/zentral-back-go/internal/folder"
	"github.com/3qual/zentral-back-go/internal/foldertransaction"
	"github.com/3qual/zentral-back-go/internal/image"
	"github.com/3qual/zentral-back-go/internal/refreshtoken"
	"github.com/3qual/zentral-back-go/internal/session"
	"github.com/3qual/zentral-back-go/internal/transaction"
	"github.com/3qual/zentral-back-go/internal/user"
)

// InitializeApp инициализирует компоненты и возвращает все обработчики
func InitializeApp() (
	*user.UserHandler,
	*transaction.TransactionHandler,
	*folder.FolderHandler,
	*foldertransaction.FolderTransactionHandler,
	*collaborator.CollaboratorHandler,
	*accesstoken.AccessTokenHandler,
	*refreshtoken.RefreshTokenHandler,
	*auth.AuthHandler,
	*image.ImageHandler,
	*session.SessionHandler,
) {
	// Загружаем конфигурацию из .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Подключение к базе данных
	database := db.ConnectDB()

	// Миграция
	migrations.Migrate(database)

	// Инициализация репозиториев
	userRepo := user.NewUserRepository(database)
	transactionRepo := transaction.NewTransactionRepository(database)
	folderRepo := folder.NewFolderRepository(database)
	folderTransactionRepo := foldertransaction.NewFolderTransactionRepository(database)
	collaboratorRepo := collaborator.NewCollaboratorRepository(database)
	accessTokenRepo := accesstoken.NewAccessTokenRepository(database)
	refreshTokenRepo := refreshtoken.NewRefreshTokenRepository(database)
	imageRepo := image.NewImageRepository(database)
	sessionRepo := session.NewSessionRepository(database)
	authRepo := auth.NewAuthRepository(database)

	// Инициализация сервисов
	userService := user.NewUserService(userRepo)
	transactionService := transaction.NewTransactionService(transactionRepo)
	folderService := folder.NewFolderService(folderRepo)
	folderTransactionService := foldertransaction.NewFolderTransactionService(folderTransactionRepo)
	collaboratorService := collaborator.NewCollaboratorService(collaboratorRepo)
	accessTokenService := accesstoken.NewAccessTokenService(accessTokenRepo)
	refreshTokenService := refreshtoken.NewRefreshTokenService(refreshTokenRepo)
	imageService := image.NewImageService(imageRepo)
	sessionService := session.NewSessionService(sessionRepo)
	authService := auth.NewAuthService(authRepo)

	// Инициализация обработчиков
	userHandler := user.NewUserHandler(userService)
	transactionHandler := transaction.NewTransactionHandler(transactionService)
	folderHandler := folder.NewFolderHandler(folderService)
	folderTransactionHandler := foldertransaction.NewFolderTransactionHandler(folderTransactionService)
	collaboratorHandler := collaborator.NewCollaboratorHandler(collaboratorService)
	accessTokenHandler := accesstoken.NewAccessTokenHandler(accessTokenService)
	refreshTokenHandler := refreshtoken.NewRefreshTokenHandler(refreshTokenService)
	authHandler := auth.NewAuthHandler(authService)
	imageHandler := image.NewImageHandler(imageService)
	sessionHandler := session.NewSessionHandler(sessionService)

	return userHandler, transactionHandler, folderHandler, folderTransactionHandler, collaboratorHandler, accessTokenHandler, refreshTokenHandler, authHandler, imageHandler, sessionHandler
}
