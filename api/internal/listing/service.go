package listing

import (
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

var err rest.Err

// Cache provides access to local cache
type Cache interface {
	// Get a message from cache with the Channel ID
	Get(channelID string) (*Message, rest.Err)
}

// Repository provides access to the message storage.
type Repository interface {
	// Message repo methods
	GetByID(string) (*Message, rest.Err)
	GetAllMessages() (*[]Message, rest.Err)
	GetByUsername(string) (*[]Message, rest.Err)
	GetLatestPinned(string) (*Message, rest.Err)
	GetPinnedMessageByID(string) (*Message, rest.Err)
	GetAllPinnedMessages() (*[]Message, rest.Err)

	// GetByStartTime sorts messages by the start_time field.
	// start_time is sorted from soonest to expire to the latest.
	GetByStartTime(int) (*[]Message, rest.Err)

	// Comment repo methods
	GetAllComments(string) (*[]Comment, rest.Err)

	// Officers
	GetAllOfficers() (*[]Member, rest.Err)
	GetActiveOfficers() (*[]Member, rest.Err)
	// GetOfficerByUsername()
	// GetOfficerByID()

	// Roles
	GetUserRoles(string) (*[]Role, rest.Err)
	GetAllRoles() (*[]Role, rest.Err)

	// Channel
	GetChannel(id string) (*Channel, rest.Err)
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

	// Get non-expired events by the time they start
	GetByStartTime() (*[]Message, rest.Err)

	// Get comments by the MessageReference ID
	GetComments(string) (*[]Comment, rest.Err)

	// Get officers
	GetAllOfficers() (*[]Member, rest.Err)
	GetActiveOfficers() (*[]Member, rest.Err)
	// GetOfficerByUsername()
	// GetOfficerByID()

	// Get roles
	GetUserRoles(string) (*[]Role, rest.Err)
	GetAllRoles() (*[]Role, rest.Err)

	// Get channel
	GetChannel(id string) (*Channel, rest.Err)
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

	messages, err := s.r.GetAllMessages()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// GetLatestPinned returns the latest pin from a given channel
func (s *service) GetLatestPinned(channelID string) (*Message, rest.Err) {

	// attempt to get pinned message by channel id from cache
	// result, err := s.c.Get(channelID)
	// if err != nil {
	// if cant get from cache, get from mongo db
	result, err := s.r.GetLatestPinned(channelID)
	if err != nil {
		return nil, err
	}
	// }

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

func (s *service) GetByStartTime() (*[]Message, rest.Err) {

	result, err := s.r.GetByStartTime(int(time.Now().Local().AddDate(0, 0, -1).Unix()))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetComments(refID string) (*[]Comment, rest.Err) {

	result, err := s.r.GetAllComments(refID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetAllOfficers() (*[]Member, rest.Err) {
	result, err := s.r.GetAllOfficers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetActiveOfficers() (*[]Member, rest.Err) {
	result, err := s.r.GetActiveOfficers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetUserRoles(id string) (*[]Role, rest.Err) {
	result, err := s.r.GetUserRoles(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetAllRoles() (*[]Role, rest.Err) {
	result, err := s.r.GetAllRoles()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetChannel(id string) (*Channel, rest.Err) {
	result, err := s.r.GetChannel(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
