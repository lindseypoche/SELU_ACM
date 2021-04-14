package blogging

import (
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/date_utils"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

var err rest.Err

// Cache provides access to redis cache
type Cache interface {
	Set(message *listing.Message) rest.Err
	Delete(channelID string) rest.Err
}

// Repository provides access to the message storage.
type Repository interface {
	SaveMessage(*Message) rest.Err

	// EditMessage edits the message being passed in.
	// If that message was already in cache, then
	// return the listing.Message value and add to cache.
	EditMessage(*Message) (*listing.Message, rest.Err)
	DeleteMessage(string) rest.Err
	UpdatePinMessage(string, bool) (*listing.Message, rest.Err)
}

// Service provides message listing operations.
type Service interface {
	AddMessage(*Message) rest.Err
	EditMessage(*Message) rest.Err

	// DeleteMessage takes in (message id, channel id) where
	// the message id is used by Repository types
	// and channel id is used by Cache types.
	DeleteMessage(string, string) rest.Err
	UpdatePinMessage(string, string, bool) rest.Err
}

type service struct {
	r Repository
	c Cache
}

// NewMessageService creates a new message service
func NewService(repo Repository, cache Cache) Service {
	return &service{r: repo, c: cache}
}

// AddMessage adds a message by the message id
func (s *service) AddMessage(message *Message) rest.Err {

	if err = s.r.SaveMessage(message); err != nil {
		return err
	}

	return nil
}

// EditMessage ...
func (s *service) EditMessage(message *Message) rest.Err {
	message.EditedTimestamp = date_utils.GetNowUnix()

	// edit message in mongo db
	edited, err := s.r.EditMessage(message)
	if err != nil {
		return err
	}

	// if the message was successfully edited in mongo,
	// store the edited message into cache if the
	// key already exists in cache.
	if err = s.c.Set(edited); err != nil {
		return err
	}

	return nil
}

// DeleteMessage ...
func (s *service) DeleteMessage(id, channelID string) rest.Err {
	// delete message from mongo db
	if err = s.r.DeleteMessage(id); err != nil {
		return err
	}

	// delete message from cache if exists
	if err = s.c.Delete(channelID); err != nil {
		return err
	}

	return nil
}

// UpdatePinMessage ...
// Updates the message IsPinned
func (s *service) UpdatePinMessage(messageID, channelID string, isPinned bool) rest.Err {
	// update main db (mongo)
	updated, err := s.r.UpdatePinMessage(messageID, isPinned)
	if err != nil {
		return err
	}

	// if is_pinned was set to true, add it to cache
	if isPinned {
		err = s.c.Set(updated)
		if err != nil {
			return err
		}
	} else if !isPinned {
		// else if the opposite is true, remove it from cache
		err = s.c.Delete(channelID)
		if err != nil {
			return err
		}
	}

	return nil
}
