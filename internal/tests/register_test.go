package tests

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterValidInputs(t *testing.T) {
	// Create test payload
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsListTestCases[1]}

	// Create test teachers and students in the database
	err := data.CreateTeacherByEmail(registerPayload.TeacherEmail)
	assert.NoError(t, err)
	for _, studentEmail := range registerPayload.StudentEmails {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, statusCode)

	// Clean up
	deleteTeacherByEmail(registerPayload.TeacherEmail)
	for _, studentEmail := range registerPayload.StudentEmails {
		deleteStudentByEmail(studentEmail)
	}
}

func TestRegisterInvalidTeacherEmail(t *testing.T) {
	// Create test payload
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsListTestCases[1]}

	// Create only test students in the database, making the teacher email provided invalid
	for _, studentEmail := range registerPayload.StudentEmails {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)

	// Clean up
	for _, studentEmail := range registerPayload.StudentEmails {
		deleteStudentByEmail(studentEmail)
	}
}

func TestRegisterInvalidStudentEmail(t *testing.T) {
	// Create test payload
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsListTestCases[1]}

	// Create only one test teacher and one test student in the database, making the second student email provided invalid
	err := data.CreateTeacherByEmail(registerPayload.TeacherEmail)
	assert.NoError(t, err)
	err = data.CreateStudentByEmail(registerPayload.StudentEmails[0])
	assert.NoError(t, err)

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)

	// Clean up
	deleteTeacherByEmail(registerPayload.TeacherEmail)
	deleteStudentByEmail(registerPayload.StudentEmails[0])
}

func TestRegisterInvalidPayload(t *testing.T) {
	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(nil, "POST", "/api/register")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestRegisterNoStudentEmails(t *testing.T) {
	// Create test payload
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: []string{}}

	// Create only one test teacher
	err := data.CreateTeacherByEmail(registerPayload.TeacherEmail)
	assert.NoError(t, err)

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)

	// Clean up
	deleteTeacherByEmail(registerPayload.TeacherEmail)
}