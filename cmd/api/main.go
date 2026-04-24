package main

import (
	"fmt"
	"go-tweets/internal/config"
	userHandler "go-tweets/internal/handler/user"
	userRepo "go-tweets/internal/repository/user"
	userService "go-tweets/internal/service/user"
	internalsql "go-tweets/pkg/internalSQL"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "app is healthy",
		})
	})

	userRepo := userRepo.NewRepository(db)
	userService := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, validate, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
