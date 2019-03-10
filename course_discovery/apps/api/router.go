package api

import (
	//"encoding/json"
	//"io/ioutil"
	//"net/http"

	// "github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/csyezheng/course-discovery/course_discovery/apps/api/views"
)

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()



func RegisterApi(router *gin.RouterGroup) {
	router.GET("/courses", views.ListCourses)
	//router.POST("/courses", views.CreateCourse)
	router.GET("/courses/:uuid", views.RetrieveCourse)
	//router.PUT("/courses/:uuid", views.UpdateCourse)
	//router.DELETE("/courses/:uuid", views.DeleteCourse)
}