package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/3qual/zentral-back-go/internal/accesstoken"
	"github.com/3qual/zentral-back-go/internal/auth"
	"github.com/3qual/zentral-back-go/internal/collaborator"
	"github.com/3qual/zentral-back-go/internal/folder"
	"github.com/3qual/zentral-back-go/internal/foldertransaction"
	"github.com/3qual/zentral-back-go/internal/image"
	"github.com/3qual/zentral-back-go/internal/refreshtoken"
	"github.com/3qual/zentral-back-go/internal/session"
	"github.com/3qual/zentral-back-go/internal/transaction"
	"github.com/3qual/zentral-back-go/internal/user"
)

// NewRouter инициализирует маршруты для всех сущностей и возвращает роутер
func NewRouter(
	userHandler *user.UserHandler, // Изменено на импортированный UserHandler
	transactionHandler *transaction.TransactionHandler,
	folderHandler *folder.FolderHandler,
	folderTransactionHandler *foldertransaction.FolderTransactionHandler,
	collaboratorHandler *collaborator.CollaboratorHandler,
	accessTokenHandler *accesstoken.AccessTokenHandler,
	refreshTokenHandler *refreshtoken.RefreshTokenHandler,
	authHandler *auth.AuthHandler,
	imageHandler *image.ImageHandler,
	sessionHandler *session.SessionHandler,
) chi.Router {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Recoverer) // Восстановление от паники
	r.Use(middleware.Logger)    // Логирование запросов

	// CORS настройка
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}
	r.Use(cors.Handler(corsOptions))

	// Настройка маршрутов API
	r.Route("/api", func(r chi.Router) {
		r.Mount("/user", user.UserRouter(userHandler)) // Изменено на маршруты для User
		r.Mount("/transaction", transaction.TransactionRouter(transactionHandler))
		r.Mount("/folder", folder.FolderRouter(folderHandler))
		r.Mount("/foldertransaction", foldertransaction.FolderTransactionRouter(folderTransactionHandler))
		r.Mount("/collaborator", collaborator.CollaboratorRouter(collaboratorHandler))
		r.Mount("/accesstoken", accesstoken.AccessTokenRouter(accessTokenHandler))
		r.Mount("/refreshtoken", refreshtoken.RefreshTokenRouter(refreshTokenHandler))
		r.Mount("/auth", auth.AuthRouter(authHandler))
		r.Mount("/image", image.ImageRouter(imageHandler))
		r.Mount("/session", session.SessionRouter(sessionHandler))
	})

	return r
}
