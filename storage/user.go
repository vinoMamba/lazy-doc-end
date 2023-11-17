package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/models"
)

func CreateUser(ctx context.Context, user *models.User) (int64, error) {

	r, err := db.ExecContext(ctx, "insert into user(username,email,password) values(?,?,?)", user.Username, user.Email, user.Password)
	id, err2 := r.LastInsertId()
	if err2 != nil {
		return (0), err2
	}
	return id, err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	err := db.GetContext(ctx, &u, "select * from user where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func DeleteUser(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}
