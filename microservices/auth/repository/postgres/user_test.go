package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"spotify/internal/model"
)

func TestRepository_GetUserByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	testLogin := "testuser"
	mockUser := &model.User{
		ID:           uuid.New(),
		Login:        testLogin,
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		AvatarURL:    "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("success - user found", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at FROM "user" WHERE login = $1`)

		mock.ExpectQuery(query).
			WithArgs(testLogin).
			WillReturnRows(
				sqlmock.NewRows([]string{"user_id", "login", "email", "password_hash", "avatar_url", "created_at", "updated_at"}).
					AddRow(mockUser.ID, mockUser.Login, mockUser.Email, mockUser.PasswordHash, mockUser.AvatarURL, mockUser.CreatedAt, mockUser.UpdatedAt),
			)

		user, err := repo.GetUserByLogin(context.Background(), testLogin)

		assert.NoError(t, err)
		require.NotNil(t, user)
		assert.Equal(t, mockUser, user)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("fail - user not found", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at FROM "user" WHERE login = $1`)

		mock.ExpectQuery(query).
			WithArgs(testLogin).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserByLogin(context.Background(), testLogin)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, user)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	newUser := model.User{
		ID:           uuid.New(),
		Login:        "newuser",
		Email:        "new@example.com",
		PasswordHash: "new_hashed_password",
		AvatarURL:    "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("success - user created", func(t *testing.T) {
		query := regexp.QuoteMeta(`INSERT INTO "user" (user_id, login, email, password_hash, avatar_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)

		mock.ExpectExec(query).
			WithArgs(newUser.ID, newUser.Login, newUser.Email, newUser.PasswordHash, newUser.AvatarURL, newUser.CreatedAt, newUser.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.CreateUser(context.Background(), newUser)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("fail - user already exists", func(t *testing.T) {
		query := regexp.QuoteMeta(`INSERT INTO "user" (user_id, login, email, password_hash, avatar_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)

		mock.ExpectExec(query).
			WithArgs(newUser.ID, newUser.Login, newUser.Email, newUser.PasswordHash, newUser.AvatarURL, newUser.CreatedAt, newUser.UpdatedAt).
			WillReturnError(&pgconn.PgError{Code: "23505"})

		err := repo.CreateUser(context.Background(), newUser)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConflict)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	mockUser := &model.User{
		ID:           uuid.New(),
		Login:        "testuser",
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("GetUserByEmail - success", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at FROM "user" WHERE email = $1`)

		mock.ExpectQuery(query).
			WithArgs(mockUser.Email).
			WillReturnRows(
				sqlmock.NewRows([]string{"user_id", "login", "email", "password_hash", "avatar_url", "created_at", "updated_at"}).
					AddRow(mockUser.ID, mockUser.Login, mockUser.Email, mockUser.PasswordHash, mockUser.AvatarURL, mockUser.CreatedAt, mockUser.UpdatedAt),
			)

		user, err := repo.GetUserByEmail(context.Background(), mockUser.Email)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetUserByID - success", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at FROM "user" WHERE user_id = $1`)

		mock.ExpectQuery(query).
			WithArgs(mockUser.ID.String()).
			WillReturnRows(
				sqlmock.NewRows([]string{"user_id", "login", "email", "password_hash", "avatar_url", "created_at", "updated_at"}).
					AddRow(mockUser.ID, mockUser.Login, mockUser.Email, mockUser.PasswordHash, mockUser.AvatarURL, mockUser.CreatedAt, mockUser.UpdatedAt),
			)

		user, err := repo.GetUserByID(context.Background(), mockUser.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_UpdateUserAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	userID := uuid.New().String()
	avatarPath := "path/to/avatar.jpg"

	t.Run("success - avatar updated", func(t *testing.T) {
		query := regexp.QuoteMeta(`UPDATE "user" SET avatar_url = $1 WHERE user_id = $2`)

		mock.ExpectExec(query).
			WithArgs(avatarPath, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdateUserAvatar(context.Background(), userID, avatarPath)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_UpdateUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	user := model.User{
		ID:           uuid.New(),
		Login:        "updated_login",
		Email:        "updated@example.com",
		PasswordHash: "new_hash",
	}

	t.Run("success", func(t *testing.T) {
		query := regexp.QuoteMeta(`
		UPDATE "user"
		SET login = $1,
			email = $2,
			password_hash = $3
		WHERE user_id = $4`)

		mock.ExpectExec(query).
			WithArgs(user.Login, user.Email, user.PasswordHash, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdateUserProfile(context.Background(), user)
		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("db error mapped", func(t *testing.T) {
		query := regexp.QuoteMeta(`
		UPDATE "user"
		SET login = $1,
			email = $2,
			password_hash = $3
		WHERE user_id = $4`)

		mock.ExpectExec(query).
			WithArgs(user.Login, user.Email, user.PasswordHash, user.ID).
			WillReturnError(&pgconn.PgError{Code: "23505"})

		err := repo.UpdateUserProfile(context.Background(), user)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "user already exist")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestRepository_GetUsersByIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("success - returns two users", func(t *testing.T) {
		ids := []string{"id1", "id2"}

		uid1 := uuid.New()
		uid2 := uuid.New()

		query := regexp.QuoteMeta(`
			SELECT user_id, login, email, password_hash, avatar_url, created_at, updated_at
			FROM "user"
			WHERE user_id IN ($1,$2)
		`)

		rows := sqlmock.NewRows([]string{
			"user_id", "login", "email", "password_hash", "avatar_url", "created_at", "updated_at",
		}).
			AddRow(uid1, "alice", "a@example.com", "h1", "", time.Now(), time.Now()).
			AddRow(uid2, "bob", "b@example.com", "h2", "", time.Now(), time.Now())

		mock.ExpectQuery(query).
			WithArgs("id1", "id2").
			WillReturnRows(rows)

		users, err := repo.GetUsersByIDs(context.Background(), ids)
		require.NoError(t, err)
		require.Len(t, users, 2)
		assert.Equal(t, "alice", users[0].Login)
		assert.Equal(t, "bob", users[1].Login)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
