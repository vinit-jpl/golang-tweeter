package main

import (
	"fmt"
	"go-tweets/internal/config"
	commentHandler "go-tweets/internal/handler/comment"
	postHandler "go-tweets/internal/handler/post"
	userHandler "go-tweets/internal/handler/user"
	commentRepo "go-tweets/internal/repository/comment"
	postRepo "go-tweets/internal/repository/post"
	userRepo "go-tweets/internal/repository/user"
	commentService "go-tweets/internal/service/comment"
	postService "go-tweets/internal/service/post"
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
	postRepo := postRepo.NewPostRepository(db)
	commentRepo := commentRepo.NewCommentRepository(db)

	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)
	commentService := commentService.NewCommentService(cfg, commentRepo, postRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)
	commentHandler := commentHandler.NewHandler(r, validate, commentService)

	userHandler.RouteList(cfg.SecretJwt)
	postHandler.RouterList(cfg.SecretJwt)
	commentHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
