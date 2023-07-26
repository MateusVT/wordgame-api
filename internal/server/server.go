package server

import (
	"github.com/fleetdm/wordgame/internal/api"
	"log"
	"net/http"
	"time"
)

// Server structure holds a pointer to the Handler structure from api package.
type Server struct {
	H *api.Handler
}

// NewServer function returns a new instance of the Server structure
func NewServer(h *api.Handler) *Server {
	return &Server{
		H: h,
	}
}

// SetupRoutes function sets up the necessary routes for the application.
func (s *Server) SetupRoutes() {

	// Create a file server to serve the Swagger UI
	fs := http.FileServer(http.Dir("./docs/swaggerui"))

	http.Handle("/new", HandlerMiddleware(http.HandlerFunc(s.H.NewGameHandler))) // Route to start a new game
	http.Handle("/guess", HandlerMiddleware(http.HandlerFunc(s.H.GuessHandler))) // Route to make a guess in a game
	http.Handle("/docs/", http.StripPrefix("/docs", fs))                         // Route to serve the API Docs via the Swagger UI

}

// Serve function starts the HTTP server on the provided address. It returns an error if anything goes wrong.
func (s *Server) Serve(addr string) error {
	log.Printf("Starting server on http://%s", addr)
	return http.ListenAndServe(addr, nil)
}

// HandlerMiddleware function is a middleware that sets the Content-Type header to application/json.
func HandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
