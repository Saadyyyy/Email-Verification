package main

import (
	"fmt"

	"email.v1/config"
	"email.v1/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// Mode Gin, ganti sesuai kebutuhan. ubah ke ReleaseMode jika sudah selesai development
	// gin.SetMode(gin.ReleaseMode)
}

func main() {
	cfg := config.InitConfig()
	g := gin.Default()

	db := database.InitDBPostgres(cfg)
	database.DBMigration(db)

	g.Use(cors.Default())

	// Route
	// <Route function disini >

	g.Run(fmt.Sprintf(":%d", cfg.SERVERPORT))
}
