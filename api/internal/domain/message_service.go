package domain

import (
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/date_utils"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
)

// MessageService ...
type MessageService interface {
	CreateMessage(message Message) (*Response, error)
	GetMessage(string) (*Message, error)
	GetMessagesByAuthor(string) (*[]Message, rest.Err)
	GetAllMessages() (*[]Message, rest.Err)
	UpdateMessage(*Message) (*Response, rest.Err)
	UpdateReaction(MessageReaction) rest.Err
	DeleteMessage(string) rest.Err
	RemoveReaction(MessageReaction) rest.Err
	GetFeatured(string) (*Pin, rest.Err)
	UpdateLatestPin(*Pin) rest.Err
}

type messageService struct {
	messageRepo MessageRepository
}

// NewMessageService creates a new message service
func NewMessageService(messageRepo MessageRepository) MessageService {
	return &messageService{messageRepo: messageRepo}
}

// CreateMessage saves a message
func (s *messageService) CreateMessage(message Message) (*Response, error) {

	result, err := s.messageRepo.Save(message)
	if err != nil {
		return nil, err
	}
	return result, nil
	// return &Response{Success: "your message was successfully posted. link (http://...)"}, nil
}

// GetMessage gets a specified message by the message id
func (s *messageService) GetMessage(messageID string) (*Message, error) {

	result, err := s.messageRepo.GetByID(messageID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMessagesByAuthor gets all messages created by author
func (s *messageService) GetMessagesByAuthor(authorID string) (*[]Message, rest.Err) {

	result, err := s.messageRepo.GetByAuthor(authorID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllMessages gets all messages
func (s *messageService) GetAllMessages() (*[]Message, rest.Err) {

	messages, err := s.messageRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// UpdateMessage ...
func (s *messageService) UpdateMessage(message *Message) (*Response, rest.Err) {

	// update timestamp
	message.EditedTimestamp = date_utils.GetNowUnix()
	resp, err := s.messageRepo.Update(message)
	if err != nil {
		// log error
		return nil, err
	}
	return resp, nil
}

// DeleteMessage deletes a message using the id
func (s *messageService) DeleteMessage(messageID string) rest.Err {
	err := s.messageRepo.Delete(messageID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMessageReactions updates reactions for a message
func (s *messageService) UpdateReaction(r MessageReaction) rest.Err {
	err := s.messageRepo.AddReaction(r)
	if err != nil {
		return err
	}
	return nil
}

func (s *messageService) RemoveReaction(r MessageReaction) rest.Err {

	err := s.messageRepo.DeleteReaction(r)
	if err != nil {
		return err
	}

	return nil
}

// GetFeatured
func (s *messageService) GetFeatured(channelID string) (*Pin, rest.Err) {

	result, err := s.messageRepo.GetLatestPinned(channelID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *messageService) UpdateLatestPin(pin *Pin) rest.Err {
	// update pin timestamp
	t := date_utils.GetNowUnix()
	pin.PinnedAt = t

	err := s.messageRepo.InsertLatestPinned(pin)
	if err != nil {
		return err
	}
	return nil
}
