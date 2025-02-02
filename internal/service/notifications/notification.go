package service

import "github.com/margar-melkonyan/watch-later.git/internal/repository"

type NotificationService struct {
	repository *repository.NotificationRepository
}

func NewNotificationService(repository *repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		repository: repository,
	}
}

func (s *NotificationService) GetUnreadNotifications() {}

func (s *NotificationService) MarkAsReadNotification() {}

func (s *NotificationService) CreateNotification() {}
