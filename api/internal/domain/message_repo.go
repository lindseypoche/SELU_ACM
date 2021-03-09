package domain

import "github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"

// MessageRepository is an interface for blog repositories
type MessageRepository interface {
	Save(Message) (*Response, error)
	GetByID(messageID string) (*Message, error)
	GetAll() (*[]Message, rest.Err)
	GetByAuthor(string) (*[]Message, rest.Err)
	Update(*Message) (*Response, rest.Err)
	AddReaction(MessageReaction) rest.Err
	Delete(string) rest.Err
	DeleteReaction(MessageReaction) rest.Err
}
