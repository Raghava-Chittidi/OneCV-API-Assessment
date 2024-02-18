package data

import (
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/models"
)

// Creates a student in the database when given an email
func CreateStudentByEmail(email string) (error) {
	student := &models.Student{Email: email, RegisteredTeachers: nil}
	result := database.DB.Table("students").Create(student)
	if result.Error != nil {
		return errors.FormatError(models.StudentModelString, result.Error)
	}

	return nil
}

// Gets student from database when given their email
func GetStudentByEmail(email string) (*models.Student, error) {
	var student models.Student
	result := database.DB.Table("students").Where("email = ?", email).First(&student)
	if result.Error != nil {
		return nil, errors.FormatError(models.StudentModelString, result.Error)
	}

	return &student, nil
}

// Gets list of students from database when given a list of their emails
func GetStudentsByEmailsList(studentEmails []string) ([]models.Student, error) {
	var students []models.Student
	result := database.DB.Table("students").Where("email IN ?", studentEmails).Find(&students)
	if result.Error != nil {
		return nil, errors.FormatError(models.StudentModelString, result.Error)
	}

	return students, nil
}

// Gets list of emails of all the common students of the teachers given
func GetCommonStudentEmailsUnderTeachers(teacherEmails []string) ([]string, error) {
	var studentEmails []string
	/* 
	    Join students, teachers and teacher_registered_students tables. Filter by teacher emails given and group 
	    by the student ids. Return the emails of students where the number of teachers they are registered to is 
		equal to the length of the teacher emails list given
	*/
	result := database.DB.Table("students").
						  Select("students.email").
						  Joins("JOIN teacher_registered_students ON teacher_registered_students.student_id = students.id").
						  Joins("JOIN teachers ON teachers.id = teacher_registered_students.teacher_id").
						  Where("teachers.email IN ?", teacherEmails).
						  Group("students.id").
						  Having("COUNT(*) = ?", len(teacherEmails)).
						  Find(&studentEmails)

	if result.Error != nil {
		return nil, errors.FormatError(models.StudentModelString, result.Error)
	}

	return studentEmails, nil
}

// Checks whether a student exists in the database for every email in the student emails list given
func CheckIfStudentsEmailsListIsValid(emails []string) error {
	for _, email := range emails {
		_, err := GetStudentByEmail(email)
		if err != nil {
			return err
		}
	}

	return nil
}