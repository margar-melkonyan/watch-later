package notifications

type NotificationService struct {
	repository NotificationRepository,
}

func NewNotificationService(repository NotificationRepository) {
	return &NotificationService{
		repository: repository,
	}
}

func (s *NotificationService) GetUnreadNotifications() {}

func (s *NotificationService) MarkAsReadNotification() {}

func (s *NotificationService) CreateNotification() {}
