package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"spotify/internal/model"
)

func (m *Repository) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (id, login, email, password_hash, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.Conn.ExecContext(ctx,
		query,
		user.ID, user.Login, user.Email, user.PasswordHash,
		user.AvatarURL, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		return handlePostgresError(err)
	}
	return nil
}

func (m *Repository) GetUserByEmail(ctx context.Context, email string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE email = $1`

	user, err := m.selectUser(ctx, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Repository) GetUserByLogin(ctx context.Context, login string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE login = $1`

	user, err := m.selectUser(ctx, query, login)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Repository) selectUser(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	rows := m.Conn.QueryRowContext(ctx, query, args...)
	user := &model.User{}
	err := rows.Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user not found: %w", ErrNotFound)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", ErrInternal)
	}

	return user, nil
}

func handlePostgresError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return fmt.Errorf("user already exist: %w", ErrConflict)
		default:
			return fmt.Errorf("postgres error (%s): %s", ErrInternal, pgErr.Message)
		}
	}
	return fmt.Errorf("unknown database error: %w", ErrInternal)
}
