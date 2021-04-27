package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/cmd/server"
)

var r *gin.Engine
var sv server.Server

func Listen() {
	sv, err := server.NewServer(8080, server.TestMode)

	if err != nil {
		panic(err)
	}

	r = sv.Router

	taskRoutes()
	listRoutes()

	r.NoRoute(noRouteHandler)
	r.NoMethod(methodNotAllowedHandler)

	sv.Listen()
}

func tempHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Handler test OK",
	})
}

func taskRoutes() {
	group := r.Group("/task")

	group.GET("", tempHandler)
	group.GET("/{id}", tempHandler)
	group.POST("/create", tempHandler)
	group.PATCH("/{id}", tempHandler)
	group.PUT("/{id}", tempHandler)
	group.DELETE("/{id}", tempHandler)
}

func listRoutes() {
	group := r.Group("/list")

	group.GET("", tempHandler)
	group.GET("/{id}", tempHandler)
	group.POST("/create", tempHandler)
	group.PATCH("/{id}", tempHandler)
	group.PUT("/{id}", tempHandler)
	group.DELETE("/{id}", tempHandler)
}

func noRouteHandler(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Endpoint not found",
	})
}

func methodNotAllowedHandler(c *gin.Context) {
	c.JSON(405, gin.H{
		"message": "Method not allowed",
	})
}
