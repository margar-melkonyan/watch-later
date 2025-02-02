package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type MediaProperty map[string]interface{}

type Media struct {
	ID                uint64
	ModelID           uint64
	ModelType         string
	Name              string
	Filename          string
	MimeType          string
	DefaultProperties MediaProperty
	CustomProperties  MediaProperty
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

type MediaRepositoryInterface interface {
	Get(id uint64) (*Media, error)
	GetAll(modelID uint64, modelType string) (*[]Media, error)
	Create(media *Media) error
	Update(media *Media, modelID uint64, modelType string) error
	Delete(mediaID uint64) error
}

type MediaRepository struct {
	db *sql.DB
}

func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (properties *MediaProperty) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, properties)
}

func (r *MediaRepository) Get(id uint64) (*Media, error) {
	var media Media
	query := "SELECT * FROM media WHERE id = $1 AND deleted_at IS NULL"
	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&media.ID,
		&media.ModelID,
		&media.ModelType,
		&media.Name,
		&media.Filename,
		&media.MimeType,
		&media.DefaultProperties,
		&media.CustomProperties,
		&media.CreatedAt,
		&media.UpdatedAt,
		&media.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &media, nil
}

func (r *MediaRepository) GetAll(modelID uint64, modelType string) (*[]Media, error) {
	var medias []Media
	query := "SELECT * FROM media WHERE model_id = $1 AND model_type = $2 AND deleted_at IS NULL"
	rows, err := r.db.Query(query, modelID, modelType)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var media Media
		err = rows.Scan(
			&media.ID,
			&media.ModelID,
			&media.ModelType,
			&media.Name,
			&media.Filename,
			&media.MimeType,
			&media.DefaultProperties,
			&media.CustomProperties,
			&media.CreatedAt,
			&media.UpdatedAt,
			&media.DeletedAt,
		)

		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}

	return &medias, nil
}

func (r *MediaRepository) Create(media *Media) error {
	return nil
}

func (r *MediaRepository) Update(media *Media, modelID uint64, modelType string) error {
	return nil
}

func (r *MediaRepository) Delete(mediaID uint64) error {
	return nil
}
