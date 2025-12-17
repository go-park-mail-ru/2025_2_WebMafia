package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"spotify/internal/model"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

func (m *Repository) CreateUser(ctx context.Context, user model.User) error {
	const op = "repository.CreateUser "
	query := `INSERT INTO "user" (user_id, login, email, password_hash, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.Conn.ExecContext(ctx,
		query,
		user.ID, user.Login, user.Email, user.PasswordHash,
		user.AvatarURL, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, handlePostgresError(err))
	}
	return nil
}

func (m *Repository) GetUserByEmail(ctx context.Context, email string) (res *model.User, err error) {
	const op = "repository.GetUserByEmail"
	query := `SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM "user" WHERE email = $1`

	user, err := m.selectUser(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func (m *Repository) GetUserByLogin(ctx context.Context, login string) (res *model.User, err error) {
	const op = "repository.GetUserByLogin"

	query := `SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM "user" WHERE login = $1`

	user, err := m.selectUser(ctx, query, login)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func (m *Repository) UpdateUserAvatar(ctx context.Context, userID string, avatarPath string) error {
	const op = "repository.UpdateUserAvatar"

	query := `UPDATE "user" SET avatar_url = $1 WHERE user_id = $2`
	_, err := m.Conn.ExecContext(ctx, query, avatarPath, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, handlePostgresError(err))
	}
	return nil
}

func (m *Repository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	const op = "repository.GetUserByID"

	query := `SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at 
			  FROM "user" WHERE user_id = $1`

	user, err := m.selectUser(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func (m *Repository) UpdateUserProfile(ctx context.Context, user model.User) error {
	const op = "repository.UpdateUserProfile"

	query := `
		UPDATE "user"
		SET login = $1,
			email = $2,
			password_hash = $3
		WHERE user_id = $4`

	_, err := m.Conn.ExecContext(ctx, query,
		user.Login,
		user.Email,
		user.PasswordHash,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, handlePostgresError(err))
	}
	return nil
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
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return user, nil
}

func (m *Repository) GetUsersByIDs(ctx context.Context, ids []string) ([]model.User, error) {
	const op = "repository.GetUsersByIDs"

	if len(ids) == 0 {
		return []model.User{}, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))

	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at FROM "user"
		WHERE user_id IN (%s)`, strings.Join(placeholders, ","))

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s query failed: %w", op, err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(
			&u.ID,
			&u.Login,
			&u.Email,
			&u.PasswordHash,
			&u.AvatarURL,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s scan failed: %w", op, err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s rows iteration failed: %w", op, err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found: %w", ErrNotFound)
	}
	return users, nil
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
			return fmt.Errorf("postgres error: %w", err)
		}
	}
	return fmt.Errorf("unknown database error: %w", err)
}
