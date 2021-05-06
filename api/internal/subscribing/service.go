package subscribing

import "github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"

var err rest.Err

// Repository provides access to the message storage.
type Repository interface {
	SaveMember(*Member) rest.Err
	UpdateMember(*Member) rest.Err
	DeleteMember(string) rest.Err
	UpdateOfficerContent(*Member, string) rest.Err
	DeleteOfficerContent(string) rest.Err
}

// Service provides message listing operations.
type Service interface {
	SaveMember(*Member) rest.Err
	UpdateMember(*Member) rest.Err
	DeleteMember(string) rest.Err
	UpdateOfficerContent(*Member, string) rest.Err
	DeleteOfficerContent(string) rest.Err
}

type service struct {
	r Repository
}

// NewMessageService creates a new message service
func NewService(repo Repository) Service {
	return &service{r: repo}
}

func (s *service) SaveMember(member *Member) rest.Err {

	err = s.r.SaveMember(member)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateMember(member *Member) rest.Err {
	err = s.r.UpdateMember(member)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteMember(mid string) rest.Err {
	err = s.r.DeleteMember(mid)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateOfficerContent(member *Member, id string) rest.Err {
	err = s.r.UpdateOfficerContent(member, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteOfficerContent(id string) rest.Err {
	err = s.r.DeleteOfficerContent(id)
	if err != nil {
		return err
	}

	return nil
}
