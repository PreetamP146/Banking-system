package repository

import (
	"banking-system/internal/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UsersRepository interface {
	CreateUser(req *models.RegisterUserRequest) error
	IsUniqueEmail(email string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
}
type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) CreateUser(req *models.RegisterUserRequest) error {
	_, err := r.db.Exec("INSERT INTO users (id, first_name, last_name, email, password_hash, role, is_active, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		uuid.New(), req.First_name, req.Last_name, req.Email, req.Password, "user", true, time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (r *usersRepository) IsUniqueEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
	}
	return true, nil
}
