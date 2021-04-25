package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ServerMode
// 0 - Debug Mode
// 1 - Release Mode
// 2 - Test Mode
type ServerMode int

const (
	DebugMode   ServerMode = iota // 0
	ReleaseMode                   // 1
	TestMode                      // 2
)

// Server struct
// Base info for the server
type Server struct {
	port   int
	Router *gin.Engine
}

// Return Server
func NewServer(port int, mode ServerMode) (sv Server) {
	sv = Server{
		port:   port,
		Router: gin.New(),
	}

	// Setting the server mode
	// More info:
	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case ReleaseMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	sv.Router.Use(gin.Recovery())

	return
}

// Enable Cors
func (s *Server) EnableCors(engine *gin.Engine, allowedOrigins []string) *Server {
	s.
		Router.
		Use(cors.New(cors.Config{
			AllowOrigins:     allowedOrigins,
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{""},
			AllowCredentials: true,
			MaxAge:           50 * time.Second,
		}))

	return s
}

// Start server
func (s *Server) Listen() {
	s.Router.Run(":" + strconv.Itoa(int(s.port)))
}
