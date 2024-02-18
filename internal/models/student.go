package models

import (
	"gorm.io/gorm"
)

const StudentModelString = "student"

type Student struct {
	gorm.Model
	Email string `json:"email" gorm:"unique;not null"`
	Suspended bool `json:"suspended" gorm:"default:false"`
	RegisteredTeachers []Teacher `json:"registeredTeachers" gorm:"many2many:student_registeredTeachers;constraint:OnDelete:CASCADE"`
}
