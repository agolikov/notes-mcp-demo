package main

import (
	"context"
	"log"
	"mcp_demo/config"
	_ "mcp_demo/controllers"
	_ "mcp_demo/docs"
	_ "mcp_demo/repositories"
	_ "mcp_demo/routes"
	_ "mcp_demo/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gorm.io/gorm"
)

// @title           Entry Point API
// @version         1.0
// @description     API Server for Entry Point Application
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	// Initialize environment and database
	config.LoadEnvVariables()
	db := config.ConnectDatabase()

	// Setup database cleanup before any potential exit points
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	r := gin.Default()

	router, err := InitializeRouter(db)
	if err != nil {
		log.Fatal("Failed to initialize router:", err)
	}

	// Update Swagger configuration
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.DocExpansion("none"),
	))

	r.GET("/openapi", func(c *gin.Context) {
		c.File("docs/swagger.yaml")
	})

	router.SetupRoutes(r)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a server with the gin handler
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

//func InitializeRouter(db *gorm.DB) (*routes.Router, error)
