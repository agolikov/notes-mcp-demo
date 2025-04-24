// Ensure this file properly uses the injected controllers
package routes

import (
	"log"
	"mcp_demo/controllers"
	"mcp_demo/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	noteController *controllers.NoteController
}

func NewRouter(
	noteController *controllers.NoteController,
) *Router {
	// Add validation to ensure controllers are not nil
	if noteController == nil {
		log.Fatal("noteController is nil in NewRouter")
	}

	return &Router{
		noteController: noteController,
	}
}

func (r *Router) SetupRoutes(router *gin.Engine) {
	// Add validation to ensure router is not nil
	if r == nil {
		log.Fatal("Router instance is nil in SetupRoutes")
	}

	router.Use(middleware.CorsMiddleware())

	// Note routes
	SetupNoteRoutes(router, r.noteController)
}
