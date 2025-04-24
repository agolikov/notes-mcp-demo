package repositories

import (
	"mcp_demo/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) GetAll() ([]entities.Note, error) {
	var notes []entities.Note
	result := r.db.Preload("Tags").Find(&notes)
	return notes, result.Error
}

func (r *NoteRepository) GetByID(id uuid.UUID) (*entities.Note, error) {
	var note entities.Note
	result := r.db.Preload("Tags").First(&note, "id = ?", id)
	return &note, result.Error
}

func (r *NoteRepository) Add(note *entities.Note) error {
	// Handle tags
	var tags []entities.Tag
	for _, tag := range note.Tags {
		var existingTag entities.Tag
		result := r.db.First(&existingTag, "name = ?", tag.Name)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}
		if result.Error == gorm.ErrRecordNotFound {
			if err := r.db.Create(&tag).Error; err != nil {
				return err
			}
			tags = append(tags, tag)
		} else {
			tags = append(tags, existingTag)
		}
	}
	note.Tags = tags

	return r.db.Create(note).Error
}

func (r *NoteRepository) Edit(note *entities.Note) error {
	// Handle tags
	var tags []entities.Tag
	for _, tag := range note.Tags {
		var existingTag entities.Tag
		result := r.db.First(&existingTag, "name = ?", tag.Name)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}
		if result.Error == gorm.ErrRecordNotFound {
			if err := r.db.Create(&tag).Error; err != nil {
				return err
			}
			tags = append(tags, tag)
		} else {
			tags = append(tags, existingTag)
		}
	}
	note.Tags = tags

	return r.db.Save(note).Error
}

func (r *NoteRepository) Remove(id uuid.UUID) error {
	return r.db.Delete(&entities.Note{}, "id = ?", id).Error
}
