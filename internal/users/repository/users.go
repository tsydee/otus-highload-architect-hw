package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/tsydim/otus-highload-architect-hw/internal/apperrs"

	"github.com/tsydim/otus-highload-architect-hw/internal/databases"
	"github.com/tsydim/otus-highload-architect-hw/internal/users"
)

type Repository struct {
	db *databases.DB
}

func NewRepository(db *databases.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(ctx context.Context, user users.User) error {
	const query = `
		INSERT INTO users (
		    id, first_name, second_name, birthdate, biography, city, gender, password
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.FirstName, user.SecondName, user.BirthDate,
		user.Biography, user.City, user.Gender, user.Password)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, userID users.UserID) (users.User, error) {
	const query = `
		SELECT id, first_name, second_name, birthdate, biography, city, gender, password
		FROM users
		WHERE id = $1
	`
	var user users.User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.BirthDate,
		&user.Biography,
		&user.City,
		&user.Gender,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, apperrs.ErrNotFound
		}
		return user, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}
