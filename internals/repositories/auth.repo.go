package repositories

import (
	"bookstore/internals/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	*sqlx.DB
}

func InitAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (a *AuthRepo) FindByEMail(body models.AuthModel) ([]models.AuthModel, error) {
	query := "select * from users where email = ?"
	result := []models.AuthModel{}
	if err := a.Select(&result, query, body.Email); err != nil {
		return nil, err
	}
	return result, nil
}
func (a *AuthRepo) SaveUser(body models.AuthModel) error {
	query := "insert into users(email, password) values(?,?)"
	if _, err := a.Exec(query, body.Email, body.Password); err != nil {
		return err
	}
	return nil
}
