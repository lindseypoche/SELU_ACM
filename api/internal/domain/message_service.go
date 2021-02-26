package domain

// var (
// 	MessageSrvc MessageService = &messageService{}
// )

// MessageService ...
type MessageService interface {
	CreateMessage(message Message) (*Response, error)
	GetMessage(string) (*Message, error)
	GetAllMessages() (*[]Message, error)
	UpdateMessage(int64)
	DeleteMessage(int64)
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

// GetMessage gets a specified message from the blog id
func (s *messageService) GetMessage(messageID string) (*Message, error) {

	result, err := s.messageRepo.GetByID(messageID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllMessages gets all messages
func (s *messageService) GetAllMessages() (*[]Message, error) {

	messages, err := s.messageRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// UpdateMessage ...
func (s *messageService) UpdateMessage(int64) {

}

// DeleteMessage ...
func (s *messageService) DeleteMessage(int64) {

}
