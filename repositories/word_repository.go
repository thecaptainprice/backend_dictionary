// repositories/word_repository.go

package repositories

import (
	"database/sql"

	"github.com/thecaptainprice/dictionary-app/backend/models"
)

type WordRepository struct {
	DB *sql.DB
}

func NewWordRepository(db *sql.DB) *WordRepository {
	return &WordRepository{DB: db}
}

// GetWords retrieves all words from the database
func (r *WordRepository) GetWords() ([]models.Word, error) {
	query := "SELECT id, word, meaning FROM word"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := []models.Word{}
	for rows.Next() {
		var word models.Word
		err := rows.Scan(&word.ID, &word.Word, &word.Meaning)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

// GetWordByID retrieves a word from the database by its ID
func (r *WordRepository) GetWordByID(id uint64) (*models.Word, error) {
	query := "SELECT id, word, meaning FROM word WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var word models.Word
	err := row.Scan(&word.ID, &word.Word, &word.Meaning)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &word, nil
}

// CreateWord creates a new word in the database
func (r *WordRepository) CreateWord(word *models.Word) error {
	query := "INSERT INTO word (word, meaning) VALUES (?, ?)"
	result, err := r.DB.Exec(query, word.Word, word.Meaning)
	if err != nil {
		return err
	}

	wordID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	word.ID = uint64(wordID)

	return nil
}

// UpdateWord updates an existing word in the database
func (r *WordRepository) UpdateWord(word *models.Word) error {
    query := "UPDATE word SET word=?, meaning=? WHERE id=?"
    _, err := r.DB.Exec(query, word.Word, word.Meaning, word.ID)
    if err != nil {
        return err
    }
    return nil
}

// DeleteWord deletes an existing word from the database
func (r *WordRepository) DeleteWord(id uint64) error {
	query := "DELETE FROM word WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
