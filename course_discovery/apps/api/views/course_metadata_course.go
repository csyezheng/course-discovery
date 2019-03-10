package views

import (
	"errors"
	// "github.com/csyezheng/course-discovery/course_discovery/apps/models"
	"net/http"

	//"log"
	//"net/http"
	//"strconv"
	// "time"

	//"github.com/csyezheng/course-discovery/course_discovery/apps/models"
	// "github.com/csyezheng/course-discovery/course_discovery/config"
	"github.com/csyezheng/course-discovery/course_discovery/common"
	"github.com/csyezheng/course-discovery/course_discovery/apps/api/forms"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// "strconv"

	// "github.com/smallnest/gen/dbmeta"
)



// GET /api/v1/courses
//
// Parameters:
//   uuid: string Course ID as returned by the API
func ListCourses(c *gin.Context) {

	var form forms.CourseSearchForm
	err := c.MustBindWith(&form, binding.Form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	search := common.NewSearch()
	//err := c.MustBindWith(&form, binding.Form)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	result, err := search.Courses(form)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	//c.Header("x-result-count", strconv.Itoa(form.Count))
	//c.Header("x-result-offset", strconv.Itoa(form.Offset))
	c.JSON(http.StatusOK, result)
}


// GET /api/v1/courses/:uuid
//
// Parameters:
//   uuid: string Course UUID as returned by the API
func RetrieveCourse(c *gin.Context) {
	uuid := c.Param("uuid")
	search := common.NewSearch()

	course, err := search.Course(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Course", errors.New("Invalid uuid")))
		return
	}
	c.JSON(http.StatusOK, course)
}

//func CreateCourse(c *gin.Context) {
//
//	var course = models.Course{
//		Created:             time.Now(),
//		Modified:             time.Now(),
//		UUID:                 c.PostForm("uuid"),
//		Key:                  c.PostForm("key"),
//		Title:                c.PostForm("title"),
//		ShortDescription:     c.PostForm("short_description"),
//		FullDescription:      c.PostForm("full_description"),
//		CardImageURL:         c.PostForm("card_image_url"),
//		Slug:                 c.PostForm("slug"),
//		Number:               c.PostForm("number"),
//		LevelTypeID:          c.PostForm("level_type_id"),
//		PartnerID:            c.PostForm("partner_id"),
//		VideoID:              c.PostForm("video_id"),
//		CanonicalCourseRunID: c.PostForm("canonical_courserun_id"),
//		Image:                c.PostForm("image"),
//		Outcome:              c.PostForm("outcome"),
//		PrerequisitesRaw:     c.PostForm("prerequisites_raw"),
//		SyllabusRaw:          c.PostForm("syllabus_raw"),
//	}
//	db.Create(&course)
//
//	search := common.NewSearch()
//	uuid, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		log.Printf("could not find image for id: %s", err.Error())
//		c.Data(http.StatusNotFound, "image", []byte(""))
//		return
//	}
//
//	photo := search.FindPhotoByID(photoID)
//	photo.PhotoFavorite = true
//	conf.Db().Save(&photo)
//	c.JSON(http.StatusOK, http.Response{})
//}

//func UpdateCourse(c *gin.Context) {
//	DB := conf.Db()
//	id := ps.ByName("id")
//
//	coursemetadatacourse := &models.CourseMetadataCourse{}
//	if DB.First(coursemetadatacourse, id).Error != nil {
//		http.NotFound(w, r)
//		return
//	}
//
//	updated := &models.CourseMetadataCourse{}
//	if err := readJSON(r, updated); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if err := dbmeta.Copy(coursemetadatacourse, updated); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if err := DB.Save(coursemetadatacourse).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	writeJSON(w, coursemetadatacourse)
//}
//
//func DeleteCourse(c *gin.Context) {
//	DB := conf.Db()
//	id := ps.ByName("id")
//	coursemetadatacourse := &models.CourseMetadataCourse{}
//
//	if DB.First(coursemetadatacourse, id).Error != nil {
//		http.NotFound(w, r)
//		return
//	}
//	if err := DB.Delete(coursemetadatacourse).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}
