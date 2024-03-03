package repositories

import (
	"bookstore/internals/models"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}

func (b *BookRepo) FindAll() ([]models.BookModel, error) {
	query := "select * from books"
	result := []models.BookModel{}
	if err := b.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil

}
func (b *BookRepo) SaveBook(body models.BookModel) error {
	query := "insert into books(title, description, author) values (?,?,?)"
	if _, err := b.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil

}

// Menggunakan parameter id
func (b *BookRepo) FindOneBook(id int) ([]models.BookModel, error) {
	query := "select * from books where id=?"
	result := []models.BookModel{}
	if err := b.Select(&result, query, id); err != nil {
		return nil, err
	}
	return result, nil
}
func (b *BookRepo) DeleteBookById(id int) error {
	query := "delete from books where id=?"
	if _, err := b.Exec(query, id); err != nil {
		return err
	}
	return nil
}
