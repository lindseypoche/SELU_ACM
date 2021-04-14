package reacting

import (
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

// Repository provides access to the message storage.
type Repository interface {
	SaveReaction(MessageReaction) rest.Err
	// GetReactions gets the reactions from a message in the database.
	// MessageReactions from the package reacting, should be returned to the client.
	GetReactions(string) (*MessageReactions, rest.Err)
	DeleteReaction(MessageReaction) rest.Err
}

// Service provides message listing operations.
type Service interface {
	AddReaction(MessageReaction) rest.Err
	GetAllReactions(string) (*MessageReactions, rest.Err)
	RemoveReaction(MessageReaction) rest.Err
}

type service struct {
	r Repository
}

// NewService creates a new message reaction service
func NewService(repo Repository) Service {
	return &service{r: repo}
}

func (s *service) AddReaction(mr MessageReaction) rest.Err {
	err := s.r.SaveReaction(mr)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveReaction(mr MessageReaction) rest.Err {
	err := s.r.DeleteReaction(mr)
	if err != nil {
		return err
	}
	return nil
}

// GetAllReactions ❌ (postponed)
func (s *service) GetAllReactions(channelID string) (*MessageReactions, rest.Err) {
	mrs, err := s.r.GetReactions(channelID)
	if err != nil {
		return nil, err
	}
	return mrs, nil
}
