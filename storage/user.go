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

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, "SELECT * FROM user WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}
