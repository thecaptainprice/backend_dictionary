package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thecaptainprice/dictionary-app/backend/models"
	"github.com/thecaptainprice/dictionary-app/backend/services"
	"github.com/thecaptainprice/dictionary-app/backend/utils"
)

type WordHandler struct {
	WordService *services.WordService
}

func NewWordHandler(wordService *services.WordService) *WordHandler {
	return &WordHandler{WordService: wordService}
}

// GetWordsHandler retrieves all words from the database
func (h *WordHandler) GetWordsHandler(w http.ResponseWriter, r *http.Request) {
	words, err := h.WordService.GetWords()
	if err != nil {
		utils.Logger(w, err, "Unable to retrieve words")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

// GetWordsHandler retrieves all words from the database
func (h *WordHandler) GetWordsHandler2(w http.ResponseWriter, r *http.Request) {
	words, err := h.WordService.GetWords()
	if err != nil {
		utils.Logger(w, err, "Unable to retrieve words")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

// GetWordByIDHandler retrieves a word from the database by its ID
func (h *WordHandler) GetWordByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get word ID from URL parameter
	vars := mux.Vars(r)
	wordID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid word ID")
		return
	}

	// Get word from service
	word, err := h.WordService.GetWordByID(wordID)
	if err != nil {
		utils.Logger(w, err, "Unable to retrieve word")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

// CreateWordHandler creates a new word in the database
func (h *WordHandler) CreateWordHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body into word struct
	var word models.Word
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		utils.Logger(w, err, "Invalid request body")
		return
	}

	// Create word using service
	err = h.WordService.CreateWord(&word)
	if err != nil {
		utils.Logger(w, err, "Unable to create word")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

// UpdateWordHandler updates an existing word in the database
func (h *WordHandler) UpdateWordHandler(w http.ResponseWriter, r *http.Request) {
	// Get word ID from URL parameter
	vars := mux.Vars(r)
	wordID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid word ID")
		return
	}

	// Decode request body into word struct
	var word models.Word
	err = json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		utils.Logger(w, err, "Invalid request body")
		return
	}

	// Set word ID to URL parameter ID
	word.ID = wordID

	// Update word using service
	err = h.WordService.UpdateWord(&word)
	if err != nil {
		utils.Logger(w, err, "Unable to update word")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

// DeleteWordHandler deletes an existing word from the database
func (h *WordHandler) DeleteWordHandler(w http.ResponseWriter, r *http.Request) {
	// Get the word ID from the request URL parameters

	vars := mux.Vars(r)
	wordID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid word ID")
		return
	}

	// Call the DeleteWord method of the WordService to delete the word
	err = h.WordService.DeleteWord(wordID)
	if err != nil {
		utils.Logger(w, err, err.Error())
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Word with ID %d has been deleted", wordID)
}
