package service

import (
	"context"

	"github.com/margar-melkonyan/watch-later.git/internal/repository"
)

type NotificationService struct {
	notificationRepository *repository.NotificationRepository
	userRepository         *repository.UserRepository
}

func NewNotificationService(
	notificationRepository *repository.NotificationRepository,
	userRepository *repository.UserRepository,
) *NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
		userRepository:         userRepository,
	}
}

func (s *NotificationService) GetUnreadNotifications(ctx context.Context) ([]repository.Notification, error) {
	email := ctx.Value("user_email").(string)
	user, err := s.userRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}
	notifications, err := s.notificationRepository.GetUnreadNotifications(user.ID)

	return notifications, err
}

func (s *NotificationService) MarkAsReadNotification(id uint64) error {
	return s.notificationRepository.MarkAsReadNotification(id)
}

func (s *NotificationService) MultipleMarkAsRead(ids []uint64) {
	for _, id := range ids {
		s.notificationRepository.MarkAsReadNotification(id)
	}
}

func (s *NotificationService) CreateNotification(notification repository.Notification) error {
	return s.notificationRepository.CreateNotification(&notification)
}
