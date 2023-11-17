package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.ExecContext(ctx, "insert into user(username,email,password) values(?,?,?)", user.Username, user.Email, user.Password)
	return err
}
