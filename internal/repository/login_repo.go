package repository

import (
	"banking-system/internal/models"
	"database/sql"
)

func (r *usersRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id,password_hash, role, is_active FROM users WHERE email = $1", email).Scan(&user.ID, &user.Password_Hash, &user.Role, &user.Is_active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if !user.Is_active {
		return nil, nil
	}
	return &user, nil
}
