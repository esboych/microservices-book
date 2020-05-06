package http

import (
	"github.com/esboych/microservices-book/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	//GetById(string) (*access_token.AccesToken, *errors.RestErr)
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler {
		service: service,
	}
}

func (handler *accessTokenHandler )GetByID(c *gin.Context){
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

