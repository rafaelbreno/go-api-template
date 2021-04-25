package routes

import "github.com/gin-gonic/gin"

var r *gin.Engine

func init() {
	r = gin.Default()

	taskRoutes()
	listRoutes()

	r.Run(":80")
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
