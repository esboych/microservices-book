package db

import (
	"github.com/esboych/microservices-book/bookstore_oauth-api/src/domain/access_token"
	"github.com/esboych/microservices-book/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccesToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccesToken, *errors.RestErr){
	return nil, errors.NewInternalServerError("db conn is not implemented yet")
}

