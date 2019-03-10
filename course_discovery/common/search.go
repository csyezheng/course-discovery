package common


import (
	// "github.com/csyezheng/course-discovery/course_discovery/config"
	"time"
	// "fmt"

	"github.com/jinzhu/gorm"
	"github.com/csyezheng/course-discovery/course_discovery/apps/models"
	"github.com/csyezheng/course-discovery/course_discovery/context"
	"github.com/csyezheng/course-discovery/course_discovery/apps/api/forms"
)

// Search searches given an originals path and a db instance.
type Search struct {
	// originalsPath string
	db            *gorm.DB
}

// SearchCount is the total number of search hits.
type SearchCount struct {
	Total int
}

// PhotoSearchResult is a found mediafile.
type CourseSearchResult struct {
	// Course
	ID                   int
	Created              time.Time
	Modified             time.Time
	UUID                 string
	Key                  string
	Title                string
	ShortDescription     string
	FullDescription      string
	CardImageURL         string
	Slug                 string
	Number               string
	LevelTypeID          int
	PartnerID            int
	VideoID              int
	CanonicalCourseRunID int
	Image                string
	Outcome              string
	PrerequisitesRaw     string
	SyllabusRaw          string

}

// NewSearch returns a new Search type with a given path and db instance.
func NewSearch() *Search {
	instance := &Search{
		// originalsPath: originalsPath,
		db: context.DB,
	}

	return instance
}

// Photos searches for photos based on a Form and returns a PhotoSearchResult slice.
func (s *Search) Courses(form forms.CourseSearchForm) ([]models.Course, error) {
	var courses []models.Course
	s.db.Find(&courses)

	return courses, nil
}

// Photos searches for photos based on a Form and returns a PhotoSearchResult slice.
func (s *Search) Course(uuid string) (course models.Course, err error) {
	err = s.db.Where("uuid = ?", uuid).First(&course).Error
	return course, err
}

//// FindFiles finds files returning maximum results defined by limit
//// and finding them from an offest defined by offset.
//func (s *Search) FindFiles(limit int, offset int) (files []models.File) {
//	s.db.Where(&models.File{}).Limit(limit).Offset(offset).Find(&files)
//
//	return files
//}

//// FindFileByID returns a mediafile given a certain ID.
//func (s *Search) FindFileByID(id string) (file models.File) {
//	s.db.Where("id = ?", id).First(&file)
//
//	return file
//}
//
//// FindFileByHash finds a file with a given hash string.
//func (s *Search) FindFileByHash(fileHash string) (file models.File) {
//	s.db.Where("file_hash = ?", fileHash).First(&file)
//	return file
//}
//
//// FindPhotoByID returns a Photo based on the ID.
//func (s *Search) FindPhotoByID(photoID uint64) (photo models.Photo) {
//	s.db.Where("id = ?", photoID).First(&photo)
//
//	return photo
//}