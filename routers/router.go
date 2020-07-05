package routers

import (
	v1 "github.com/getchipman/nwapi/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	// Get Employees list
	apiv1.GET("/employees", v1.GetEmployees)

	return r
}
