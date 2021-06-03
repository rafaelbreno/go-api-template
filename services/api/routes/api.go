package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/cmd/server"
	"github.com/rafaelbreno/go-api-template/api/internal/handler"
)

var r *gin.Engine
var sv server.Server

func Listen() {
	sv, err := server.NewServer(8070, server.DebugMode)

	if err != nil {
		panic(err)
	}

	r = sv.Router

	taskRoutes()
	listRoutes()

	r.GET("/api/ping", pingHandler)

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

func pingHandler(c *gin.Context) {
	response, err := http.Get(fmt.Sprintf("%s/ping", os.Getenv("API_HOST")))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Pong API",
			"error":   "Unable to reach Auth service",
		})
		return
	}

	responseAuth, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Pong API",
			"error":   "Unable to read Auth's service response",
		})
		return
	}

	var responseData struct {
		Message string `json:"message"`
	}

	err = json.Unmarshal(responseAuth, &responseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Pong API",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Pong API",
		"auth_message": responseData.Message,
	})
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Endpoint not found",
	})
}

func methodNotAllowedHandler(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"message": "Method not allowed",
	})
}
