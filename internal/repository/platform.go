package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Platform struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name" validate:"required,alphanum,min=8,max=32"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type PlatformRepositoryInterface interface {
	Get(id uint64) (*Platform, error)
	GetAll() ([]*Platform, error)
	Create(p *Platform) error
	Update(platform *Platform, id uint64) error
	Delete(id uint64) error
	Restore(id uint64) error
}

type PlatformRepository struct {
	db *sql.DB
}

func NewPlatformRepository(db *sql.DB) *PlatformRepository {
	return &PlatformRepository{
		db: db,
	}
}

func (r *PlatformRepository) Get(id uint64) (*Platform, error) {
	var platform Platform
	query := "SELECT id, name FROM platforms WHERE id=$1 AND deleted_at IS NULL"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&platform.ID, &platform.Name)

	if err != nil {
		return nil, err
	}

	return &platform, nil
}

func (r *PlatformRepository) GetAll() ([]*Platform, error) {
	var platforms []*Platform
	query := "SELECT id, name FROM platforms WHERE deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var platform Platform
		err = rows.Scan(&platform.ID, &platform.Name)
		if err != nil {
			return nil, err
		}
		platforms = append(platforms, &platform)
	}

	return platforms, nil
}

func (r *PlatformRepository) Create(platform *Platform) error {
	query := "INSERT INTO platforms (name) VALUES ($1)"
	result, err := r.db.Exec(query, platform.Name)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("platform was not created")
	}

	return nil
}

func (r *PlatformRepository) Update(platform *Platform, id uint64) error {
	query := "UPDATE platforms SET name = $1 WHERE id = $2 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, platform.Name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("platform was not updated")
	}

	return nil
}

func (r *PlatformRepository) Delete(id uint64) error {
	query := "UPDATE platforms SET deleted_at=now() WHERE id = $1 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("platform was not deleted")
	}

	return nil
}

func (r *PlatformRepository) Restore(id uint64) error {
	query := "UPDATE platforms SET deleted_at=NULL WHERE id=$1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("platforms was not restored")
	}

	return nil
}
