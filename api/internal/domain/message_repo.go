package domain

// MessageRepository is an interface for blog repositories
type MessageRepository interface {
	Save(message Message) (*Response, error)
	GetByID(messageID string) (*Message, error)
	GetAll() (*[]Message, error)
	GetByAuthor() error
	Update() error
	Delete() error
}
