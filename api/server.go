package api

import (
	"context"
	"log/slog"
	"net/http"
	// "github.com/ohhfishal/schedule/db"
)

type Database interface {
	// CreateEvent(ctx context.Context, arg db.CreateEventParams) (db.Event, error)
}

type Server struct {
	Logger   *slog.Logger
	Database Database
}

func (server Server) getEvents(ctx context.Context, userId int, params GetEventsParams) http.Handler {
	return nil
}

func (server Server) createEvent(ctx context.Context, userId int) http.Handler {
	return nil
}

func (server Server) getEvent(ctx context.Context, userId int, eventId int) http.Handler {
	return nil
}

func (server Server) getUser(ctx context.Context, username string) http.Handler {
	return nil
}

func (server Server) putUser(ctx context.Context, username string) http.Handler {
	return nil
}

// ServerInterface bindings

// (GET /dev/event/{userId})
func (server Server) GetEvents(w http.ResponseWriter, r *http.Request, userId int, params GetEventsParams) {
}

// (POST /dev/event/{userId})
func (server Server) CreateEvent(w http.ResponseWriter, r *http.Request, userId int) {}

// (GET /dev/event/{userId}/{eventId})
func (server Server) GetEventById(w http.ResponseWriter, r *http.Request, userId int, eventId int) {}

// (GET /dev/user/{username})
func (server Server) GetUserByUsername(w http.ResponseWriter, r *http.Request, username string) {}

// (POST /dev/user/{username})
func (server Server) PutUserByUsername(w http.ResponseWriter, r *http.Request, username string) {}
