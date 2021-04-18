package commenting

import (
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/date_utils"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

var err rest.Err

// Repository provides access to the comment storage.
type Repository interface {
	SaveComment(*Comment) rest.Err
	EditComment(*Comment) rest.Err
	DeleteComment(string) rest.Err
}

// Service provides comment listing operations.
type Service interface {
	AddComment(*Comment) rest.Err
	EditComment(*Comment) rest.Err
	DeleteComment(string) rest.Err
}

type service struct {
	r Repository
}

// NewMessageService creates a new message service
func NewService(repo Repository) Service {
	return &service{r: repo}
}

// AddComment adds a comment by the message id
func (s *service) AddComment(comment *Comment) rest.Err {

	err = s.r.SaveComment(comment)
	if err != nil {
		return err
	}

	return nil
}

// EditComment ...
func (s *service) EditComment(comment *Comment) rest.Err {
	comment.EditedTimestamp = date_utils.GetNowUnix()

	// edit comment in mongo db
	err = s.r.EditComment(comment)
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment ...
func (s *service) DeleteComment(id string) rest.Err {
	// delete comment from mongo db
	err = s.r.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
