package tests

import (
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some student emails test cases
var studentEmailsTestCases = []string {
	"studentone123@gmail.com", "studenttwo123@gmail.com", "studentthree123@gmail.com",
}

// Some student email lists test cases
var studentEmailsListTestCases = [][]string {
	{ "studentone123@gmail.com"}, 
	{ "studentone123@gmail.com", "studenttwo123@gmail.com"}, 
	{ "studentone123@gmail.com", "studenttwo123@gmail.com", "studentthree123@gmail.com"}, 
}

// Deletes a student when given their email
func deleteStudentByEmail(email string) error {
	result := database.DB.Table("students").Where("email = ?", email).Unscoped().Delete(&models.Student{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Cleans up all the students created during tests by deleting them
func cleanUpStudentsCreated(emails []string) {
	for _, email := range emails {
		deleteStudentByEmail(email)
	}
}

// Asserts whether the two students are equal by comparing their fields
func assertEqualStudent(t *testing.T, expected *models.Student, student *models.Student) {
	assert.Equal(t, expected.Email, student.Email)
	assert.Equal(t, expected.Suspended, student.Suspended)
	assert.Equal(t, expected.RegisteredTeachers, student.RegisteredTeachers)
}

// Tests CreateStudentByEmail and GetStudentByEmail functions
func TestCreateAndGetStudentByEmail(t *testing.T) {
	for _, email := range studentEmailsTestCases {
		// Expected student
		expected := models.Student{Email: email, Suspended: false, RegisteredTeachers: nil}
		
		// Create and get student
		err := data.CreateStudentByEmail(email)
		assert.NoError(t, err)
		student, err := data.GetStudentByEmail(email)
		assert.NoError(t, err)
	
		// Check if they are equal
		assertEqualStudent(t, &expected, student)
		err = deleteStudentByEmail(email)
		assert.NoError(t, err)
	}
}

// Tests GetStudentByEmailsList function
func TestGetStudentByEmailsList(t *testing.T) {
	var expectedStudents []models.Student
	// Creates and stores students in expectedStudents
	for _, email := range studentEmailsTestCases {
		err := data.CreateStudentByEmail(email)
		assert.NoError(t, err)
		student, err := data.GetStudentByEmail(email)
		assert.NoError(t, err)
		expectedStudents = append(expectedStudents, *student)
	}

	// Get actual students and compare if expected number of students is equal to actual number of students
	actualStudents, err := data.GetStudentsByEmailsList(studentEmailsListTestCases[2])
	assert.NoError(t, err)
	assert.Equal(t, len(expectedStudents), len(actualStudents))

	// Check if expected student is equal to actual student
	for index, expected := range actualStudents {
		student := actualStudents[index]
		assertEqualStudent(t, &expected, &student)
	}

	// Clean up
	cleanUpStudentsCreated(studentEmailsTestCases)
}

// Tests GetCommonStudentEmailsUnderTeachers function
func TestGetCommonStudentEmailsUnderTeachers(t *testing.T) {
	expectedStudentEmails := studentEmailsListTestCases[0]
	// Create students
	for _, studentEmail := range studentEmailsTestCases {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	for index, teacherEmail := range teacherEmailsTestCases {
		// Create and get teacher
		err := data.CreateTeacherByEmail(teacherEmail)
		assert.NoError(t, err)
		teacher, err := data.GetTeacherByEmail(teacherEmail)
		assert.NoError(t, err)

		// Get array of students based on test case
		students, err := data.GetStudentsByEmailsList(studentEmailsListTestCases[index])
		assert.NoError(t, err)
		// Register these students to the teacher
		err = database.DB.Model(&teacher).Association("RegisteredStudents").Append(students)
		assert.NoError(t, err)
	}

	// Get emails of common students under the teachers created
	actualStudentEmails, err := data.GetCommonStudentEmailsUnderTeachers(teacherEmailsTestCases)
	assert.NoError(t, err)
	// Compare if expected number of student emails is equal to actual number of student emails
	assert.Equal(t, len(expectedStudentEmails), len(actualStudentEmails))

	// Compare expected and actual student emails
	for index, studentEmail := range actualStudentEmails {
		assert.Equal(t, expectedStudentEmails[index], studentEmail)
	}

	// Clean up
	cleanUpStudentsCreated(studentEmailsTestCases)
	cleanUpTeachersCreated(teacherEmailsTestCases)
}

// Tests CheckIfStudentsEmailsListIsValid function
func TestCheckIfStudentsEmailsListIsValid(t *testing.T) {
	// Should get error as intitally no student in the test cases exists in the database
	err := data.CheckIfStudentsEmailsListIsValid(studentEmailsTestCases)
	assert.Error(t, err)

	// Create students
	for _, studentEmail := range studentEmailsTestCases {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Should not get any error as all students in the test cases exist in the database
	err = data.CheckIfStudentsEmailsListIsValid(studentEmailsTestCases)
	assert.NoError(t, err)

	// Should get error in the case where at least one student in the test case does not exist in the database
	invalidStudentEmailsTestCases := []string{"studentone@gmail.com", "studenttwo@gmail.com", "studentfour@gmail.com",}
	err = data.CheckIfStudentsEmailsListIsValid(invalidStudentEmailsTestCases)
	assert.Error(t, err)
	
	// Clean up
	cleanUpStudentsCreated(studentEmailsTestCases)
}