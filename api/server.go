package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ohhfishal/schedule/db"
)

var _ ServerInterface = &Server{}

type Database interface {
	CreateUser(ctx context.Context, username string) (db.User, error)
	GetUserByUsername(ctx context.Context, username string) (db.User, error)
}

type Server struct {
	Logger   *slog.Logger
	Database Database
}

func NewServer(logger *slog.Logger, database Database) Server {
	return Server{
		Logger:   logger,
		Database: database,
	}
}

// Entry Point from CLI
func (server *Server) Run(ctx context.Context, addr string) error {
	router := http.NewServeMux()
	handler := HandlerFromMux(server, router)
	s := &http.Server{
		Handler: handler,
		Addr:    addr,
	}
	go func() {
		server.Logger.Info(`starting server`, `addr`, addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			server.Logger.Error(`server died due to error`, `error`, err)
		}
	}()

	server.Logger.Info(`main goroutine sleeping`)
	<-ctx.Done()

	err := ctx.Err()
	server.Logger.Info(`context ended`, `error`, err)

	if err := s.Shutdown(context.TODO()); err != nil {
		return fmt.Errorf(`shutting down: %w`, err)
	}
	return nil
}

// /api/.../event

func (server Server) createEvent(ctx context.Context, userId int, event Event) http.Handler {
	return nil
}

func (server Server) getEvent(ctx context.Context, userId int, eventId int) http.Handler {
	return nil
}

func (server Server) getEvents(ctx context.Context, userId int, params GetEventsParams) http.Handler {
	return nil
}

// /api/.../user

func (server Server) getUser(ctx context.Context, username string) http.Handler {
	user, err := server.Database.GetUserByUsername(ctx, username)
	if err != nil {
		return Error(
			http.StatusInternalServerError,
			fmt.Errorf(`failed to get user: %w`, err),
		)
	}
	return JSON(http.StatusCreated, user)
}

func (server Server) putUser(ctx context.Context, username string) http.Handler {
	user, err := server.Database.CreateUser(ctx, username)
	if err != nil {
		return Error(
			http.StatusInternalServerError,
			fmt.Errorf(`failed to create user: %w`, err),
		)
	}
	return JSON(http.StatusCreated, user)
}

func JSON(status int, v any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(v)
	})
}

func Text(status int, content any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprint(w, content)
	})
}

func Error(status int, err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, err.Error(), status)
	})
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	err := json.NewDecoder(r.Body).Decode(&v)
	if errors.Is(err, io.EOF) {
		return v, errors.New("body is empty or incomplete")
	}

	if err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

// ServerInterface bindings

// (GET /dev/event/{userId})
func (server Server) GetEvents(w http.ResponseWriter, r *http.Request, userId int, params GetEventsParams) {
	handler := server.getEvents(r.Context(), userId, params)
	handler.ServeHTTP(w, r)
}

// (POST /dev/event/{userId})
func (server Server) CreateEvent(w http.ResponseWriter, r *http.Request, userId int) {
	event, err := decode[Event](r)
	if err != nil {
		Error(400, err).ServeHTTP(w, r)
		return
	}
	handler := server.createEvent(r.Context(), userId, event)
	handler.ServeHTTP(w, r)
}

// (GET /dev/event/{userId}/{eventId})
func (server Server) GetEventById(w http.ResponseWriter, r *http.Request, userId int, eventId int) {
	handler := server.getEvent(r.Context(), userId, eventId)
	handler.ServeHTTP(w, r)
}

// (GET /dev/user/{username})
func (server Server) GetUserByUsername(w http.ResponseWriter, r *http.Request, username string) {
	handler := server.getUser(r.Context(), username)
	handler.ServeHTTP(w, r)
}

// (POST /dev/user/{username})
func (server Server) PutUserByUsername(w http.ResponseWriter, r *http.Request, username string) {
	handler := server.putUser(r.Context(), username)
	handler.ServeHTTP(w, r)
}

// (GET /health)
func (server Server) Health(w http.ResponseWriter, r *http.Request) {
	Text(200, "OK").ServeHTTP(w, r)
}
