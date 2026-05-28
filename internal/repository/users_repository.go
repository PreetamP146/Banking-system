package repository

import "database/sql"

type UsersRepository interface {
	CreateUser() error
	GetUserByEmail(email string) (bool, error)
}
type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) CreateUser() error {
	// Implement the logic to create a user in the database
	return nil
}
func (r *usersRepository) GetUserByEmail(email string) (bool, error) {
	var userID int
	err := r.db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		return false, err
	}
	if userID != 0 {
		return false, nil
	}
	return true, nil
}
