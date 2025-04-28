package api

import "net/http"

type Server struct {
}

func (server Server) getEvents(userId int, params GetEventsParams) http.Handler {
	return nil
}

func (server Server) createEvent(userId int) http.Handler {
	return nil
}

func (server Server) getEventById(userId int, eventId int) http.Handler {
	return nil
}

func (server Server) getUserByUsername(w http.ResponseWriter, r *http.Request, username string) http.Handler {
	return nil
}

func (server Server) putUserByUsername(w http.ResponseWriter, r *http.Request, username string) http.Handler {
	return nil
}

// ServerInterface bindings

// (GET /dev/event/{userId})
func (server Server) GetEvents(w http.ResponseWriter, r *http.Request, userId int, params GetEventsParams) {
	handler := server.getEvents(userId, params)
	handler.ServeHTTP(w, r)
}

// (POST /dev/event/{userId})
func (server Server) CreateEvent(w http.ResponseWriter, r *http.Request, userId int) {}

// (GET /dev/event/{userId}/{eventId})
func (server Server) GetEventById(w http.ResponseWriter, r *http.Request, userId int, eventId int) {}

// (GET /dev/user/{username})
func (server Server) GetUserByUsername(w http.ResponseWriter, r *http.Request, username string) {}

// (POST /dev/user/{username})
func (server Server) PutUserByUsername(w http.ResponseWriter, r *http.Request, username string) {}
