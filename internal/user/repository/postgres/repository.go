package postgres

import (
	"context"
	"database/sql"
	"spotify/internal/user/model"
	"spotify/internal/user/service"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) service.UserRepository {
	return &postgresUserRepository{Conn: Conn}
}

func (m *postgresUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, service.WrapError(err, service.ErrInternal, "failed to execute query")
	}

	defer func() {
		rows.Close()
	}()

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
			return nil, service.WrapError(err, service.ErrInternal, "failed to scan row")
		}
		result = append(result, user)
	}

	return result, nil
}

func (m *postgresUserRepository) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (id, login, email, password_hash, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return service.WrapError(err, service.ErrInternal, "failed to prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		user.ID, user.Login, user.Email, user.PasswordHash,
		user.AvatarURL, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return service.WrapError(err, service.ErrConflict, "user with this login or email already exists")
		}
		return service.WrapError(err, service.ErrInternal, "failed to create user")
	}

	return nil
}

func (m *postgresUserRepository) GetUserByEmail(ctx context.Context, email string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE email = $1`

	list, err := m.fetch(ctx, query, email)
	if err != nil {
		return &model.User{}, err
	}

	if len(list) > 0 {
		res = &list[0]
	} else {
		return res, service.WrapError(nil, service.ErrValidation, "user not found")
	}

	return
}

func (m *postgresUserRepository) GetUserByLogin(ctx context.Context, login string) (res *model.User, err error) {
	query := `SELECT id, login, email, password_hash, avatar_url, created_at, updated_at 
		FROM users WHERE login = $1`

	list, err := m.fetch(ctx, query, login)
	if err != nil {
		return &model.User{}, err
	}

	if len(list) > 0 {
		res = &list[0]
	} else {
		return res, service.WrapError(nil, service.ErrValidation, "user not found")
	}
	return
}

func isUniqueViolation(err error) bool {
	errStr := err.Error()
	return strings.Contains(errStr, "unique constraint") || strings.Contains(errStr, "duplicate key")
}
