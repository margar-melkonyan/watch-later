package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Category struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id" validate:"required,numeric"`
	Name      string     `json:"name" validate:"required,alphanum"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type CategoryRepositoryInterface interface {
	Get(id uint64) (*Category, error)
	GetAll() ([]*Category, error)
	GetUserCategoriesList(userId uint64) ([]*Category, error)
	Create(category *Category) error
	Update(category *Category, id uint64) error
	Delete(id uint64) error
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Get(id uint64) (*Category, error) {
	var category Category
	query := "SELECT id, name, user_id, created_at FROM categories WHERE id=$1 AND deleted_at IS NULL"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetAll() ([]*Category, error) {
	var categories []*Category
	query := "SELECT id, name, user_id, created_at FROM categories WHERE deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.UserID,
			&category.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *CategoryRepository) GetUserCategoriesList(userId uint64) ([]*Category, error) {
	var categories []*Category
	query := "SELECT id, name, user_id, created_at FROM categories WHERE user_id=$1 AND deleted_at IS NULL"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *CategoryRepository) Create(category *Category) error {
	query := "INSERT INTO categories (name, user_id) VALUES ($1, $2)"
	result, err := r.db.Exec(query, category.Name, category.UserID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("categories was not created")
	}

	return nil
}

func (r *CategoryRepository) Update(category *Category, id uint64) error {
	query := "UPDATE categories SET name=$1, updated_at=now() WHERE id=$2 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, category.Name, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("categories was not updated")
	}

	return nil
}

func (r *CategoryRepository) Delete(id uint64) error {
	query := "UPDATE categories SET  updated_at=now(), deleted_at = now() WHERE id=$1 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("categories was not deleted")
	}

	return nil
}

func (r *CategoryRepository) Restore(id uint64) error {
	query := "UPDATE categories SET deleted_at=NULL WHERE id=$1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("categories was not restored")
	}

	return nil
}
