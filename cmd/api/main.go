package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gen-ai-workshop-4-be/internal/domain"
	"gen-ai-workshop-4-be/internal/handlers"
	"gen-ai-workshop-4-be/internal/infra"
	"gen-ai-workshop-4-be/internal/middleware"
	"gen-ai-workshop-4-be/internal/service"
)

func main() {
	r := gin.Default()

	// db
	db := infra.NewSQLiteDB("app.db")
	db.AutoMigrate(&domain.User{})

	// repo + service
	repo := infra.NewGormUserRepo(db)
	svc := service.NewAuthService(repo)

	// routes
	r.POST("/register", handlers.Register(svc))
	r.POST("/login", handlers.Login(svc))
	r.GET("/me", middleware.JWTMiddleware(), handlers.Me(svc))

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	r.StaticFile("/docs/openapi.yaml", "docs/openapi.yaml")
	r.StaticFile("/swagger-ui", "docs/swagger-ui.html")

	// root
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Hello World"}) })

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
