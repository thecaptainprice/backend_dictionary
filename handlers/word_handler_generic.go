package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/thecaptainprice/dictionary-app/backend/models"
	"github.com/thecaptainprice/dictionary-app/backend/request"
)

// GetWordsHandler retrieves all words from the database
func (h *WordHandler) GetWordsHandlerGeneric(r *request.GenericRequest) (interface{}, error) {
	words, err := h.WordService.GetWords()
	if err != nil {
		return nil, err
	}
	// Send response as JSON
	return words, nil
}

// GetWordByIDHandler retrieves a word from the database by its ID
func (h *WordHandler) GetWordByIDHandlerGeneric(r *request.GenericRequest) (interface{}, error) {
	// Get word ID from URL parameter
	wordID, err := strconv.ParseUint(r.PathParams["id"], 10, 64)
	if err != nil {
		return nil, err
	}

	// Get word from service
	word, err := h.WordService.GetWordByID(wordID)
	if err != nil {
		return nil, err
	}

	// Send response as JSON
	return word, nil
}

// CreateWordHandler creates a new word in the database
func (h *WordHandler) CreateWordHandlerGeneric(r *request.GenericRequest) (interface{}, error) {
	// Decode request body into word struct
	var word models.Word
	err := json.Unmarshal(r.Body, &word)
	if err != nil {
		return nil, err
	}

	// Create word using service
	err = h.WordService.CreateWord(&word)
	if err != nil {
		return nil, err
	}

	// Send response as JSON
	return word, nil
}

// UpdateWordHandler updates an existing word in the database
func (h *WordHandler) UpdateWordHandlerGeneric(r *request.GenericRequest) (interface{}, error) {
	// Get word ID from URL parameter
	wordID, err := strconv.ParseUint(r.QueryParams.Get("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	// Decode request body into word struct
	var word models.Word
	err = json.Unmarshal(r.Body, &word)
	if err != nil {
		return nil, err
	}

	// Set word ID to URL parameter ID
	word.ID = wordID

	// Update word using service
	err = h.WordService.UpdateWord(&word)
	if err != nil {
		return nil, err
	}

	// Send response as JSON
	return word, nil
}

// DeleteWordHandler deletes an existing word from the database
func (h *WordHandler) DeleteWordHandlerGeneric(r *request.GenericRequest) (interface{}, error) {
	// Get the word ID from the request URL parameters

	wordID, err := strconv.ParseUint(r.QueryParams.Get("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	// Call the DeleteWord method of the WordService to delete the word
	err = h.WordService.DeleteWord(wordID)
	if err != nil {
		return nil, err
	}

	// Return nil,err a success message

	return nil, nil
}
