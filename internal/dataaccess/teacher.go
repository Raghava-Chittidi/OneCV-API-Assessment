package data

import (
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/models"
)

// Creates a teacher in the database when given an email
func CreateTeacherByEmail(email string) (error) {
	teacher := &models.Teacher{Email: email, RegisteredStudents: nil}
	result := database.DB.Table("teachers").Create(teacher)
	if result.Error != nil {
		return errors.FormatError(models.TeacherModelString, result.Error)
	}

	return nil
}

// Gets teacher from database when given their email
func GetTeacherByEmail(email string) (*models.Teacher, error) {
	var teacher models.Teacher
	result := database.DB.Table("teachers").Where("email = ?", email).First(&teacher)
	if result.Error != nil {
		return nil, errors.FormatError(models.TeacherModelString, result.Error)
	}

	return &teacher, nil
}

// Gets preloaded teacher with RegisteredStudents from database when given an email
func GetPreloadedTeacherByEmail(email string) (*models.Teacher, error) {
	var teacher models.Teacher
	result := database.DB.Table("teachers").
						  Preload("RegisteredStudents").
						  Where("email = ?", email).First(&teacher)

	if result.Error != nil {
		return nil, errors.FormatError(models.TeacherModelString, result.Error)
	}

	return &teacher, nil
}

// Checks whether a teacher exists in the database for every email in the teacher emails list given
func CheckIfTeachersEmailsListIsValid(emails []string) error {
	for _, email := range emails {
		_, err := GetTeacherByEmail(email)
		if err != nil {
			return err
		}
	}

	return nil
}