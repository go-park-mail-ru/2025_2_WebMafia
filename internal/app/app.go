package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/signal"
	"spotify/internal/middleware"
	httpDelivery "spotify/internal/user/delivery/http"
	userRepository "spotify/internal/user/repository/postgres"
	userService "spotify/internal/user/service"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/postgres"
	"syscall"
)

type App struct {
	server *http.Server
	cfg    *Config
	db     *sql.DB
}

func NewApp(cfg *Config) (*App, error) {
	db, err := postgres.New(context.Background(), cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	router := mux.NewRouter()
	router.Use(middleware.CORS(cfg.CORS))

	repo := userRepository.NewUserRepository(db)
	svc := userService.NewUserService(repo)

	jwtManager := jwtmanager.NewManager(cfg.JWTSecretKey, cfg.AccessTokenTTL)
	handler := httpDelivery.NewHandler(svc, jwtManager)

	authMiddleware := middleware.NewAuthMiddleware(jwtManager)
	handler.RegisterRouter(router, authMiddleware)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &App{
		server: server,
		cfg:    cfg,
		db:     db,
	}, nil
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)
	go func() {
		fmt.Println("Server is starting on port ", a.server.Addr)
		serverErrors <- a.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		fmt.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()

		if err := a.server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

	}

	fmt.Println("Server exiting")
	return nil
}
