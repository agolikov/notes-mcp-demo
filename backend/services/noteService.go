package services

import (
	"mcp_demo/models/entities"
	"mcp_demo/models/requests"
	"mcp_demo/models/responses"
	"mcp_demo/repositories"

	"github.com/google/uuid"
)

type NoteService struct {
	repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetAll() ([]responses.NoteResponse, error) {
	notes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var noteResponses []responses.NoteResponse
	for _, note := range notes {
		tagNames := make([]string, len(note.Tags))
		for i, tag := range note.Tags {
			tagNames[i] = tag.Name
		}

		noteResponses = append(noteResponses, responses.NoteResponse{
			ID:        note.ID.String(),
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
			Title:     note.Title,
			Content:   note.Content,
			Tags:      tagNames,
		})
	}

	return noteResponses, nil
}

func (s *NoteService) GetByID(id uuid.UUID) (*responses.NoteResponse, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	tagNames := make([]string, len(note.Tags))
	for i, tag := range note.Tags {
		tagNames[i] = tag.Name
	}

	response := &responses.NoteResponse{
		ID:        note.ID.String(),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		Title:     note.Title,
		Content:   note.Content,
		Tags:      tagNames,
	}

	return response, nil
}

func (s *NoteService) CreateNote(req requests.CreateNoteRequest) (*responses.NoteResponse, error) {
	tags := make([]entities.Tag, len(req.Tags))
	for i, tagName := range req.Tags {
		tags[i] = entities.Tag{Name: tagName}
	}

	note := &entities.Note{
		Title:   req.Title,
		Content: req.Content,
		Tags:    tags,
	}

	if err := s.repo.Add(note); err != nil {
		return nil, err
	}

	tagNames := make([]string, len(note.Tags))
	for i, tag := range note.Tags {
		tagNames[i] = tag.Name
	}

	response := &responses.NoteResponse{
		ID:        note.ID.String(),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		Title:     note.Title,
		Content:   note.Content,
		Tags:      tagNames,
	}

	return response, nil
}

func (s *NoteService) UpdateNote(id uuid.UUID, req requests.UpdateNoteRequest) (*responses.NoteResponse, error) {
	existingNote, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	tags := make([]entities.Tag, len(req.Tags))
	for i, tagName := range req.Tags {
		tags[i] = entities.Tag{Name: tagName}
	}

	note := &entities.Note{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
		Tags:    tags,
	}

	if err := s.repo.Edit(note); err != nil {
		return nil, err
	}

	tagNames := make([]string, len(note.Tags))
	for i, tag := range note.Tags {
		tagNames[i] = tag.Name
	}

	response := &responses.NoteResponse{
		ID:        note.ID.String(),
		CreatedAt: existingNote.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		Title:     note.Title,
		Content:   note.Content,
		Tags:      tagNames,
	}

	return response, nil
}

func (s *NoteService) Remove(id uuid.UUID) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Remove(id)
}
