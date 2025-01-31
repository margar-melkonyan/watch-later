package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Payload map[string]interface{}

type Notification struct {
	ID        uint64
	UserID    uint64
	Payload   Payload
	CreatedAt time.Time
	UpdatedAt time.Time
	ReadAt    *time.Time
}

type NotificationRepositoryInterface interface {
	GetUnreadNotifications(userID uint64) (*[]Notification, error)
	MarkAsReadNotification(id uint64) error
	CreateNotification(notification *Notification) error
}

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (payload *Payload) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, payload)
}

func (r *NotificationRepository) GetUnreadNotifications(userID uint64) (*[]Notification, error) {
	var notifications []Notification
	query := `SELECT id, user_id, payload, created_at FROM notifications 
              WHERE user_id=$1 AND read_at IS NULL ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var notification Notification
		err = rows.Scan(&notification.ID, &notification.UserID, &notification.Payload, &notification.CreatedAt)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return &notifications, nil
}

func (r *NotificationRepository) MarkAsReadNotification(id uint64) error {
	query := "UPDATE notifications SET read_at=now() WHERE id=$1 AND read_at IS NULL"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cannot update the read notification")
	}

	return nil
}

func (r *NotificationRepository) CreateNotification(notification *Notification) error {
	query := "INSERT INTO notifications (user_id, payload) VALUES ($1, $2)"
	marshaledPayload, err := json.Marshal(notification.Payload)
	if err != nil {
		return err
	}
	result, err := r.db.Exec(query, notification.UserID, marshaledPayload)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cannot create the notification")
	}

	return nil
}
