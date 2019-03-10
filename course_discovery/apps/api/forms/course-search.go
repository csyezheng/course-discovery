package forms






// Query parameters for GET /api/v1/courses
type CourseSearchForm struct {
	Query         string    `form:"q"`
	Uuid         string    `form:"uuid"`
	Order         string    `form:"order"`
	Offset        int       `form:"offset"`
	Limit        int       `form:"Limit"`
}