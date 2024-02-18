package models

import (
	"gorm.io/gorm"
)

const TeacherModelString = "teacher"

type Teacher struct {
	gorm.Model
	Email string `json:"email" gorm:"unique;not null"`
	RegisteredStudents []Student `json:"registeredStudents" gorm:"many2many:teacher_registeredStudents;constraint:OnDelete:CASCADE"`
}
