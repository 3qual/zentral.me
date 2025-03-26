package migrations

import (
	"log"

	"gorm.io/gorm"

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

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&accesstoken.AccessToken{},
		&collaborator.Collaborator{},
		&folder.Folder{},
		&foldertransaction.FolderTransaction{},
		&image.Image{},
		&refreshtoken.RefreshToken{},
		&session.Session{},
		&transaction.Transaction{},
		&user.User{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migrations has been successfully applied")
}
