package tests

import (
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some teacher emails test cases
var teacherEmailsTestCases = []string {
	"teacherone123@gmail.com", "teachertwo123@gmail.com", "teacherthree123@gmail.com",
}

// Some teacher email lists test cases
var teacherEmailsListTestCases = [][]string {
	{ "teacherone123@gmail.com"}, 
	{ "teacherone123@gmail.com", "teachertwo123@gmail.com"}, 
	{ "teacherone123@gmail.com", "teachertwo123@gmail.com", "teacherthree123@gmail.com"}, 
}

// Deletes a teacher when given their email
func deleteTeacherByEmail(email string) error {
	result := database.DB.Table("teachers").Where("email = ?", email).Unscoped().Delete(&models.Teacher{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Cleans up all the teachers created during tests by deleting them
func cleanUpTeachersCreated(emails []string) {
	for _, email := range emails {
		deleteTeacherByEmail(email)
	}
}

// Asserts whether the two teachers are equal by comparing their fields
func assertEqualTeacher(t *testing.T, teacher *models.Teacher, expected *models.Teacher) {
	assert.Equal(t, teacher.Email, expected.Email)
	assert.Equal(t, teacher.RegisteredStudents, expected.RegisteredStudents)
}

// Tests CreateTeacherByEmail and GetTeacherByEmail functions
func TestCreateAndGetTeacherByEmail(t *testing.T) {
	for _, email := range teacherEmailsTestCases {
		// Expected teacher
		expected := models.Teacher{Email: email, RegisteredStudents: nil}
		
		// Create and get teacher
		err := data.CreateTeacherByEmail(email)
		assert.NoError(t, err)
		teacher, err := data.GetTeacherByEmail(email)
		assert.NoError(t, err)
	
		// Check if they are equal
		assertEqualTeacher(t, teacher, &expected)
		err = deleteTeacherByEmail(email)
		assert.NoError(t, err)
	}
}

// Tests GetPreloadedTeacherByEmail function
func TestGetPreloadedTeacherByEmail(t *testing.T) {
	teacherEmail := teacherEmailsTestCases[0]
	expected := models.Teacher{Email: teacherEmail, RegisteredStudents: nil}

	// Create and get teacher
	err := data.CreateTeacherByEmail(teacherEmail)
	assert.NoError(t, err)
	teacher, err := data.GetTeacherByEmail(teacherEmail)
	assert.NoError(t, err)

	// Create students
	for _, studentEmail := range studentEmailsTestCases {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	students, err := data.GetStudentsByEmailsList(studentEmailsTestCases)
	assert.NoError(t, err)
	// Register the students to the teacher
	err = database.DB.Model(&teacher).Association("RegisteredStudents").Append(students)
	assert.NoError(t, err)
	
	// Get actual preloaded teacher
	preloadedTeacher, err := data.GetPreloadedTeacherByEmail(teacherEmail)
	assert.NoError(t, err)

	// Set expected teacher's registered students to the students created
	expected.RegisteredStudents = students
	// Compare the expected preloaded teacher with the actual preloaded teacher
	assertEqualTeacher(t, &expected, preloadedTeacher)

	// Clean up
	cleanUpStudentsCreated(studentEmailsTestCases)
	err = deleteTeacherByEmail(teacherEmail)
	assert.NoError(t, err)
}

// Tests CheckIfTeachersEmailsListIsValid function
func TestCheckIfTeachersEmailsListIsValid(t *testing.T) {
	// Should get error as intitally no teacher in the test cases exists in the database
	err := data.CheckIfTeachersEmailsListIsValid(teacherEmailsListTestCases[2])
	assert.Error(t, err)

	// Create teachers
	for _, teacherEmail := range teacherEmailsTestCases {
		err := data.CreateTeacherByEmail(teacherEmail)
		assert.NoError(t, err)
	}

	// Should not get any error as all teachers in the test cases exist in the database
	err = data.CheckIfTeachersEmailsListIsValid(teacherEmailsTestCases)
	assert.NoError(t, err)

	// Should get error in the case where at least one teacher in the test case does not exist in the database
	invalidteacherEmailsTestCases := []string{"teacherone@gmail.com", "teachertwo@gmail.com", "teacherfour@gmail.com",}
	err = data.CheckIfTeachersEmailsListIsValid(invalidteacherEmailsTestCases)
	assert.Error(t, err)
	
	// Clean up
	cleanUpTeachersCreated(teacherEmailsTestCases)
}