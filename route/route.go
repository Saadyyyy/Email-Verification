package route

import (
	"email.v1/api/handler"
	"email.v1/api/repository"
	"email.v1/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(c *gin.Engine,  db *gorm.DB){
	emailRepository := repository.NewRepositoryEmail(db)
	emailService := service.NewRepositoryEmail(emailRepository)
	emailHandler := handler.NewHandler(emailService)

	c.POST( "/", emailHandler.CreateUser)
	//c.GET("/token")
}