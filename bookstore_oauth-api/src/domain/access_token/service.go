package access_token

import "github.com/esboych/microservices-book/bookstore_oauth-api/utils/errors"

type Repository interface {
	GetById(string) (*AccesToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccesToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccesToken, *errors.RestErr){
	accessToken, err :=  s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}