package app

import (
	"github.com/esboych/microservices-book/bookstore_oauth-api/http"
	"github.com/esboych/microservices-book/bookstore_oauth-api/reposiroty/db"
	"github.com/esboych/microservices-book/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.Run(":8080")
}
