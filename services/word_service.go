// services/word_service.go

package services

import (
	"github.com/thecaptainprice/dictionary-app/backend/models"
	"github.com/thecaptainprice/dictionary-app/backend/repositories"
)

type WordService struct {
	WordRepo *repositories.WordRepository
}

func NewWordService(wordRepo *repositories.WordRepository) *WordService {
	return &WordService{WordRepo: wordRepo}
}

// GetWords retrieves all words from the database
func (s *WordService) GetWords() ([]models.Word, error) {
	return s.WordRepo.GetWords()
}

// GetWordByID retrieves a word from the database by its ID
func (s *WordService) GetWordByID(id uint64) (*models.Word, error) {
	return s.WordRepo.GetWordByID(id)
}

// CreateWord creates a new word in the database
func (s *WordService) CreateWord(word *models.Word) error {
	return s.WordRepo.CreateWord(word)
}

// UpdateWord updates an existing word in the database
func (s *WordService) UpdateWord(word *models.Word) error {
	return s.WordRepo.UpdateWord(word)
}

// DeleteWord deletes an existing word from the database
func (s *WordService) DeleteWord(id uint64) error {
	return s.WordRepo.DeleteWord(id)
}
