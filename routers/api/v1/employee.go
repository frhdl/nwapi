package v1

import (
	"net/http"

	"github.com/getchipman/nwapi/pkg/app"

	"github.com/getchipman/nwapi/pkg/e"

	"github.com/getchipman/nwapi/models"
	"github.com/gin-gonic/gin"
)

// GetEmployees test
func GetEmployees(c *gin.Context) {
	appG := app.Gin{C: c}

	var employees []*models.Employees

	employees, err := models.GetEmployees()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 1001, nil)
	}

	data := make(map[string]interface{})
	data["lists"] = employees

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
