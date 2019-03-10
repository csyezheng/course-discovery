package course_discovery

import (
	//"encoding/json"
	//"io/ioutil"
	//"net/http"

	// "github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/csyezheng/course-discovery/course_discovery/apps/api"
)



func RegisterRoutes(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	api.RegisterApi(v1)
}