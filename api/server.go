package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ohhfishal/schedule/db"
)

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

// /api/.../event

func (server Server) getEvents(ctx context.Context, userId int, params GetEventsParams) http.Handler {
	return nil
}

func (server Server) createEvent(ctx context.Context, userId int) http.Handler {
	return nil
}

func (server Server) getEvent(ctx context.Context, userId int, eventId int) http.Handler {
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

// ServerInterface bindings

// (GET /dev/event/{userId})
func (server Server) GetEvents(w http.ResponseWriter, r *http.Request, userId int, params GetEventsParams) {
	handler := server.getEvents(r.Context(), userId, params)
	handler.ServeHTTP(w, r)
}

// (POST /dev/event/{userId})
func (server Server) CreateEvent(w http.ResponseWriter, r *http.Request, userId int) {
	handler := server.createEvent(r.Context(), userId)
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
