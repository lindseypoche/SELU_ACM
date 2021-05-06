package architecting

import "github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"

var err rest.Err

type Repository interface {
	SaveRole(*Role) rest.Err
	SaveChannel(*Channel) rest.Err
	UpdateRole(role *Role) rest.Err
	DeleteRole(id string) rest.Err
}

type Service interface {
	CreateRole(*Role) rest.Err
	CreateChannel(*Channel) rest.Err
	UpdateRole(role *Role) rest.Err
	DeleteRole(id string) rest.Err
}

type service struct {
	r Repository
}

func NewService(repo Repository) Service {
	return &service{r: repo}
}

func (s *service) CreateRole(role *Role) rest.Err {
	err = s.r.SaveRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateChannel(channel *Channel) rest.Err {
	err = s.r.SaveChannel(channel)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateRole(role *Role) rest.Err {
	err = s.r.UpdateRole(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteRole(id string) rest.Err {
	err = s.r.DeleteRole(id)
	if err != nil {
		return err
	}
	return nil
}
