package listing

import (
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

// Cache provides access to local cache
type Cache interface {
	// Get a message from cache with the Channel ID
	Get(channelID string) (*Message, rest.Err)
}

// Repository provides access to the message storage.
type Repository interface {
	GetByID(string) (*Message, rest.Err)
	GetAll() (*[]Message, rest.Err)
	GetByUsername(string) (*[]Message, rest.Err)
	GetLatestPinned(string) (*Message, rest.Err)
	GetPinnedMessageByID(string) (*Message, rest.Err)
	GetAllPinnedMessages() (*[]Message, rest.Err)
}

// Service provides message listing operations.
type Service interface {
	GetMessage(string) (*Message, rest.Err)
	GetMessagesByUsername(string) (*[]Message, rest.Err)
	GetAllMessages() (*[]Message, rest.Err)

	// Get the pinned message from cache with the Channel ID
	GetLatestPinned(string) (*Message, rest.Err)

	// Get all latest pins stored in cache
	GetAllLatestPinned() (map[string]*Message, rest.Err)

	// Get any pinned message with Message ID
	GetPinnedMessage(string) (*Message, rest.Err)
	GetAllPinnedMessages() (*[]Message, rest.Err)
}

type service struct {
	r Repository
	c Cache
}

// NewMessageService creates a new message service
func NewService(repo Repository, cache Cache) Service {
	return &service{r: repo, c: cache}
}

// GetMessage returns a specified message by the message id
func (s *service) GetMessage(messageID string) (*Message, rest.Err) {

	result, err := s.r.GetByID(messageID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMessagesByAuthor returns all messages created by an author
func (s *service) GetMessagesByUsername(username string) (*[]Message, rest.Err) {

	result, err := s.r.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllMessages returns all messages
func (s *service) GetAllMessages() (*[]Message, rest.Err) {

	messages, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// GetLatestPinned returns the latest pin from a given channel
func (s *service) GetLatestPinned(channelID string) (*Message, rest.Err) {

	// attempt to get pinned message by channel id from cache
	result, err := s.c.Get(channelID)
	if err != nil {
		// if cant get from cache, get from mongo db
		result, err = s.r.GetLatestPinned(channelID)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *service) GetAllLatestPinned() (map[string]*Message, rest.Err) {
	// result, err := s.c.GetMessages()
	// if err != nil {
	// 	return nil, err
	// }

	// return result, nil
	return nil, rest.NewStatusNotImplemented("GetAllLatestPinned not implemented")
}

// GetFeatured returns the featured message by message id from mongo db
func (s *service) GetPinnedMessage(messageID string) (*Message, rest.Err) {

	result, err := s.r.GetPinnedMessageByID(messageID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) GetAllPinnedMessages() (*[]Message, rest.Err) {

	result, err := s.r.GetAllPinnedMessages()
	if err != nil {
		return nil, err
	}
	return result, nil
}
