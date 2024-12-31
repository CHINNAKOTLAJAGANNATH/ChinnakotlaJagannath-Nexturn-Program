package repository

import (
	"a2_ems_project/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create a new user in the database
func (repo *UserRepository) CreateUser(user *model.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

// Check if a username already exists
func (repo *UserRepository) IsUsernameTaken(username string) (bool, error) {
	var existingUsername string
	err := repo.DB.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUsername)
	if err == sql.ErrNoRows {
		return false, nil // Username is not taken
	}
	if err != nil {
		return false, err // Database error
	}
	return true, nil // Username exists
}

// Authenticate user by username and password
func (repo *UserRepository) AuthenticateUser(username, password string) (*model.User, error) {
	var user model.User
	err := repo.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}
	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}
	return &user, nil
}
