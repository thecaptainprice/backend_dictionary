package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thecaptainprice/dictionary-app/backend/models"
	"github.com/thecaptainprice/dictionary-app/backend/services"
	"github.com/thecaptainprice/dictionary-app/backend/utils"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// GetUserHandler returns a specific user from the database
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get word ID from URL parameter
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid word ID")
		return
	}

	// Get word from service
	word, err := h.UserService.GetUserByID(userID)
	if err != nil {
		utils.Logger(w, err, "Unable to retrieve word")
		return
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

// CreateUserHandler creates a new user in the database
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a new User object
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Logger(w, err, "Invalid JSON request body")
		return
	}

	// Call the CreateUser method of the UserService to create the new user in the database
	user, err := h.UserService.CreateUser(&newUser), err
	if err != nil {
		utils.Logger(w, err, "User is not created")
		return
	}

	// Return the newly created user as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler Update a user in the database
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request URL parameters
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid  user ID")
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Logger(w, err, "Invalid request body")
		return
	}
	// Set word ID to URL parameter ID
	user.ID = userID

	//Check the exist user by email
	checkUser, err := h.UserService.GetUserByEmail(user.Email)
	if err != nil {
		utils.Logger(w, err, "Couldn't fetch email")
		return
	}

	if checkUser != nil {
		utils.Logger(w, errors.New("Email is Taken"), "Email is Taken")
		return
	}

	// Update word using service
	err = h.UserService.UpdateUser(&user)
	if err != nil {
		utils.Logger(w, err, "Unable to update user")
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, " user with ID %d has been Update", userID)
}

// DeleteUserHandler deletes an existing user from the database
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request URL parameters
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.Logger(w, err, "Invalid  user ID")
		return
	}

	// Call the Delete user method of the  userService to delete the  user
	err = h.UserService.DeleteUser(userID)
	if err != nil {
		utils.Logger(w, err, err.Error())
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, " user with ID %d has been deleted", userID)
}
