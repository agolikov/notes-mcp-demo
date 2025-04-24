package routes

import (
	"mcp_demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupNoteRoutes(router *gin.Engine, controller *controllers.NoteController) {
	noteGroup := router.Group("/notes")
	{
		noteGroup.GET("", controller.GetAll)
		noteGroup.POST("", controller.Create)
		noteGroup.GET("/:id", controller.GetByID)
		noteGroup.PUT("/:id", controller.Update)
		noteGroup.DELETE("/:id", controller.Delete)
	}
}
