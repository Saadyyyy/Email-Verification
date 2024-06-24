package route

import (
	"email.v1/api/users/handler"
	"email.v1/api/users/repository"
	"email.v1/api/users/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(c *gin.Engine, db *gorm.DB) {
	emailRepository := repository.NewRepositoryEmail(db)
	emailService := service.NewRepositoryEmail(emailRepository, db)
	emailHandler := handler.NewHandler(emailService)

	c.POST("register", emailHandler.CreateUser)
	c.POST("/login", emailHandler.Login)
}
