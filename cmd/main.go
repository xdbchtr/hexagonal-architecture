package main

import (
	"context"
	"library-app/internal/adapter/http"
	"library-app/internal/adapter/repository"
	"library-app/internal/core/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:4dmin2024!@db-dev.wewearseia.com/?authSource=admin&replicaSet=rs0"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize repositories
	bookCollection := client.Database("library").Collection("books")
	userCollection := client.Database("library").Collection("users")

	bookRepo := repository.NewMongoBookRepository(bookCollection)
	userRepo := repository.NewMongoUserRepository(userCollection)
	authRepo := repository.NewMongoUserRepository(userCollection)

	// Initialize services
	bookService := service.NewBookService(bookRepo)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo)

	// Initialize Gin
	r := gin.Default()

	// Initialize HTTP handlers
	bookHandler := http.NewBookHandler(bookService)
	userHandler := http.NewUserHandler(userService)
	authHandler := http.NewAuthHandler(authService)

	// Public routes (no authentication required)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Protected routes (require authentication)
	authGroup := r.Group("/").Use(http.AuthMiddleware())
	{
		authGroup.POST("/books", bookHandler.CreateBook)
		authGroup.GET("/books/:id", bookHandler.GetBook)
		authGroup.GET("/books", bookHandler.GetAllBooks)
		authGroup.PUT("/books/:id", bookHandler.UpdateBook)
		authGroup.DELETE("/books/:id", bookHandler.DeleteBook)

		authGroup.POST("/users", userHandler.CreateUser)
		authGroup.GET("/users/:id", userHandler.GetUser)
		authGroup.GET("/users", userHandler.GetAllUsers)
		authGroup.PUT("/users/:id", userHandler.UpdateUser)
		authGroup.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
