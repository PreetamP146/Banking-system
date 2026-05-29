package repository

import (
	"banking-system/internal/models"
	"database/sql"
)

func (r *usersRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT 1 FROM users WHERE email = $1", email).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
