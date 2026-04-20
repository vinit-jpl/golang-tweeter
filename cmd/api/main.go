package main

import (
	"fmt"
	"go-tweets/internal/config"
	internalsql "go-tweets/pkg/internalSQL"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	_, err = internalsql.ConnectMySQL(cfg)
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

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
