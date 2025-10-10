package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"spotify/internal/user/model"
	"spotify/internal/user/service"
)

type Repository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) service.IRepository {
	return &Repository{Conn: Conn}
}

func (m *Repository) selectUser(ctx context.Context, query string, args ...interface{}) (result []model.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, service.NewServiceError(service.ErrInternal, "failed to execute query")
	}

	defer rows.Close()

	result = make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Email,
			&user.PasswordHash,
			&user.AvatarURL,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, service.NewServiceError(service.ErrInternal, "failed to scan row")
		}
		result = append(result, user)
	}
	if err := rows.Err(); err != nil {
		return nil, service.NewServiceError(service.ErrInternal, "rows iteration failed")
	}

	if len(result) == 0 {
		return nil, service.NewServiceError(service.ErrValidation, "no users found")
	}

	return result, nil
}

func (m *Repository) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (id, login, email, password_hash, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.Conn.ExecContext(ctx,
		query,
		user.ID, user.Login, user.Email, user.PasswordHash,
		user.AvatarURL, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return service.NewServiceError(service.ErrConflict, "user with this login or email already exists")
		}
		return service.NewServiceError(service.ErrInternal, "failed to create user")
	}

	return nil
}

func (m *Repository) GetUserByEmail(ctx context.Context, email string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE email = $1`

	list, err := m.selectUser(ctx, query, email)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	} else {
		return nil, service.NewServiceError(service.ErrValidation, "user not found")
	}
}

func (m *Repository) GetUserByLogin(ctx context.Context, login string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE login = $1`

	list, err := m.selectUser(ctx, query, login)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	} else {
		return nil, service.NewServiceError(service.ErrValidation, "user not found")
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
