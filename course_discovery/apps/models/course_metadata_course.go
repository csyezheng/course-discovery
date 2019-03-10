package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Course struct {
	ID                   int         `gorm:"column:id;primary_key" json:"id"`
	Created              time.Time   `gorm:"column:created" json:"created"`
	Modified             time.Time   `gorm:"column:modified" json:"modified"`
	UUID                 string      `gorm:"column:uuid" json:"uuid"`
	Key                  string      `gorm:"column:key" json:"key"`
	Title                null.String `gorm:"column:title" json:"title"`
	ShortDescription     null.String `gorm:"column:short_description" json:"short_description"`
	FullDescription      null.String `gorm:"column:full_description" json:"full_description"`
	CardImageURL         null.String `gorm:"column:card_image_url" json:"card_image_url"`
	Slug                 string      `gorm:"column:slug" json:"slug"`
	Number               null.String `gorm:"column:number" json:"number"`
	LevelTypeID          null.Int    `gorm:"column:level_type_id" json:"level_type_id"`
	PartnerID            int         `gorm:"column:partner_id" json:"partner_id"`
	VideoID              null.Int    `gorm:"column:video_id" json:"video_id"`
	CanonicalCourseRunID null.Int    `gorm:"column:canonical_course_run_id" json:"canonical_course_run_id"`
	Image                null.String `gorm:"column:image" json:"image"`
	Outcome              null.String `gorm:"column:outcome" json:"outcome"`
	PrerequisitesRaw     null.String `gorm:"column:prerequisites_raw" json:"prerequisites_raw"`
	SyllabusRaw          null.String `gorm:"column:syllabus_raw" json:"syllabus_raw"`
}

// TableName sets the insert table name for this struct type
func (c *Course) TableName() string {
	return "course_metadata_course"
}
