package controllers

import (
	"mcp_demo/models/requests"
	"mcp_demo/models/responses"
	"mcp_demo/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NoteController struct {
	service *services.NoteService
}

func NewNoteController(service *services.NoteService) *NoteController {
	return &NoteController{service: service}
}

// GetAll godoc
// @Summary Get all notes
// @Description Retrieve all notes
// @Tags notes
// @Produce json
// @Success 200 {array} responses.NoteResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /notes [get]
func (c *NoteController) GetAll(ctx *gin.Context) {
	notes, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(200, notes)
}

// GetByID godoc
// @Summary Get a note by ID
// @Description Retrieve a specific note by its ID
// @Tags notes
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} responses.NoteResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /notes/{id} [get]
func (c *NoteController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	note, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(200, note)
}

// Create godoc
// @Summary Create a new note
// @Description Create a new journal note
// @Tags notes
// @Accept json
// @Produce json
// @Param request body requests.CreateNoteRequest true "Note data"
// @Success 201 {object} responses.NoteResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /notes [post]
func (c *NoteController) Create(ctx *gin.Context) {
	var req requests.CreateNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Error: "Invalid request format"})
		return
	}

	note, err := c.service.CreateNote(req)
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(201, note)
}

// Update godoc
// @Summary Update a note
// @Description Update an existing note
// @Tags notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param request body requests.UpdateNoteRequest true "Updated note data"
// @Success 200 {object} responses.NoteResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /notes/{id} [put]
func (c *NoteController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var req requests.UpdateNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Error: "Invalid request format"})
		return
	}

	note, err := c.service.UpdateNote(id, req)
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(200, note)
}

// Delete godoc
// @Summary Delete a note
// @Description Delete an existing note
// @Tags notes
// @Param id path string true "Note ID"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /notes/{id} [delete]
func (c *NoteController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	if err := c.service.Remove(id); err != nil {
		ctx.JSON(500, responses.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(204, nil)
}
