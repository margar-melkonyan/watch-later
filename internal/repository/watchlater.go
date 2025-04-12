package repository

import (
	"database/sql"
	"errors"
	"time"
)

type WatchLater struct {
	ID         uint64     `json:"id"`
	UserID     uint64     `json:"user_id" validate:"required,min=1,numeric"`
	CategoryID uint64     `json:"category_id" validate:"required,min=1,numeric"`
	PlatformID uint64     `json:"platform_id" validate:"required,min=1,numeric"`
	Name       string     `json:"name" validate:"required,min=4"`
	Text       string     `json:"text" validate:"required,min=4"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

type WatchLaterRepositoryInterface interface {
	Get(id uint64) (*WatchLater, error)
	GetByUser(userID uint64) ([]*WatchLater, error)
	GetByCategory(categoryID uint64) ([]*WatchLater, error)
	GetByPlatform(platformID uint64) ([]*WatchLater, error)
	GetAll() ([]*WatchLater, error)
	Create(watch *WatchLater) error
	Update(watch *WatchLater, watchLaterId uint64) error
	Delete(id uint64) error
}

type WatchLaterRepository struct {
	db *sql.DB
}

func NewWatchLaterRepository(db *sql.DB) *WatchLaterRepository {
	return &WatchLaterRepository{
		db: db,
	}
}

func (r *WatchLaterRepository) Get(id uint64) (*WatchLater, error) {
	var watch WatchLater
	query := `SELECT id, user_id, category_id, platform_id, name, text, created_at 
			  FROM watch_laters WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&watch.ID,
		&watch.UserID,
		&watch.CategoryID,
		&watch.PlatformID,
		&watch.Name,
		&watch.Text,
		&watch.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &watch, nil
}

func (r *WatchLaterRepository) getWatchLaters(query string, arg interface{}) ([]*WatchLater, error) {
	var watchLaters []*WatchLater
	var rows *sql.Rows
	var err error

	if arg != nil {
		rows, err = r.db.Query(query, arg)
	} else {
		rows, err = r.db.Query(query)
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var watchLater WatchLater
		err = rows.Scan(
			&watchLater.ID,
			&watchLater.UserID,
			&watchLater.CategoryID,
			&watchLater.PlatformID,
			&watchLater.Name,
			&watchLater.Text,
			&watchLater.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		watchLaters = append(watchLaters, &watchLater)
	}

	return watchLaters, nil
}

func (r *WatchLaterRepository) GetByUser(userID uint64) ([]*WatchLater, error) {
	return r.getWatchLaters(
		`SELECT id, user_id, category_id, platform_id, name, text, created_at  
			   FROM watch_laters 
			   WHERE user_id = $1 AND deleted_at IS NULL`,
		userID,
	)
}

func (r *WatchLaterRepository) GetByCategory(categoryID uint64) ([]*WatchLater, error) {
	return r.getWatchLaters(
		`SELECT id, user_id, category_id, platform_id, name, text, created_at  
			   FROM watch_laters 
			   WHERE category_id = $1 AND deleted_at IS NULL`,
		categoryID,
	)
}

func (r *WatchLaterRepository) GetByPlatform(platformID uint64) ([]*WatchLater, error) {
	return r.getWatchLaters(
		`SELECT id, user_id, category_id, platform_id, name, text, created_at  
			FROM watch_laters 
			WHERE platform_id = $1 AND deleted_at IS NULL`,
		platformID,
	)
}

func (r *WatchLaterRepository) GetAll() ([]*WatchLater, error) {
	return r.getWatchLaters(
		`SELECT id, user_id, category_id, platform_id, name, text, created_at 
			   FROM watch_laters 
			   WHERE deleted_at IS NULL`,
		nil)
}

func (r *WatchLaterRepository) Create(watch *WatchLater) error {
	query := `INSERT INTO watch_laters (name, text, user_id, category_id, platform_id) 
			  VALUES ($1, $2, $3, $4, $5)`
	result, err := r.db.Exec(
		query,
		watch.Name,
		watch.Text,
		watch.UserID,
		watch.CategoryID,
		watch.PlatformID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("watch laters was not created")
	}

	return nil
}

func (r *WatchLaterRepository) Update(watch *WatchLater, watchLaterId uint64) error {
	query := `UPDATE watch_laters 
			  SET name = $1, text = $2, user_id = $3, category_id = $4, platform_id = $5 
			  WHERE id = $6 AND deleted_at IS NULL`

	result, err := r.db.Exec(query,
		watch.Name,
		watch.Text,
		watch.UserID,
		watch.CategoryID,
		watch.PlatformID,
		watchLaterId,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("watch laters was not updated")
	}

	return nil
}

func (r *WatchLaterRepository) Delete(id uint64) error {
	query := "UPDATE watch_laters SET deleted_at=now() WHERE id = $1 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("watch laters was not deleted")
	}

	return nil
}
