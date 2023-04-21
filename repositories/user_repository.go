// repositories/user_repository.go

package repositories

import (
	"database/sql"
	"errors"

	"github.com/thecaptainprice/dictionary-app/backend/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserByID retrieves a user from the database by its ID
func (r *UserRepository) GetUserByID(id uint64) (*models.User, error) {
	query := "SELECT id, email, password FROM user WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrieves a user from the database by its email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT email FROM user WHERE email = ?"
	row := r.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO user (name,email, password) VALUES (?, ?,?)"
	result, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = uint64(userID)

	return nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE user SET name=?, email=?, password=? WHERE id=?"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes an existing user from the database
func (r *UserRepository) DeleteUser(id uint64) error {
	query := "DELETE FROM user WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// AuthenticateUser authenticates a user with the given email and password
func (r *UserRepository) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
