//go:build wireinject

package main

import (
	"mcp_demo/controllers"
	"mcp_demo/repositories"
	"mcp_demo/routes"
	"mcp_demo/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeRouter(db *gorm.DB) (*routes.Router, error) {
	wire.Build(
		// Repositories
		repositories.NewNoteRepository,

		// Services
		services.NewNoteService,

		// Controllers
		controllers.NewNoteController,

		// Router
		routes.NewRouter,
	)
	return nil, nil
}
