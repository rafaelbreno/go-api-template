package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/cmd/server"
	"github.com/rafaelbreno/go-api-template/api/internal/handler"
)

var r *gin.Engine
var sv server.Server

func Listen() {
	sv, err := server.NewServer(8070, server.TestMode)

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
	h := handler.NewTaskHandler()
	group := r.Group("/task")

	group.GET("", h.FindAll)
	group.GET("/:id", h.FindById)
	group.POST("/create", h.Create)
	group.PATCH("/:id", h.Update)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

func listRoutes() {
	group := r.Group("/list")
	h := handler.NewListHandler()

	group.GET("", h.FindAll)
	group.GET("/:id", h.FindByID)
	group.POST("/create", h.Create)
	group.PATCH("/:id", h.Update)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
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
