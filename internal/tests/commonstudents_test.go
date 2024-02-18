package tests

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonStudentsValidInputs(t *testing.T) {
	// Expected response data
	var commonPayload = api.CommonStudentsResponse{Students: studentEmailsListTestCases[1]}
	expected, err := json.Marshal(commonPayload)
	assert.NoError(t, err)

	// Create test students in the database
	for _, studentEmail := range commonPayload.Students {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Create test teachers in the database and register these students to the teacher
	for _, teacherEmail := range teacherEmailsListTestCases[1] {
		err := data.CreateTeacherByEmail(teacherEmail)
		assert.NoError(t, err)
		var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmail, StudentEmails: studentEmailsListTestCases[1]}
		_, _, err = CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
		assert.NoError(t, err)
	}
	
	data, status, err := CreateRequestAndReturnResponseData(nil, "GET", fmt.Sprintf("/api/commonstudents?teacher=%s&teacher=%s", 
																					teacherEmailsTestCases[0], teacherEmailsTestCases[1]))
	assert.NoError(t, err)

	// Compare expected and actual status codes
	assert.Equal(t, http.StatusOK, status)
	// Compare expected and actual response data
	assert.Equal(t, expected, data)

	// Clean up
	for _, teacherEmail := range teacherEmailsListTestCases[1] {
		deleteTeacherByEmail(teacherEmail)
	}
	for _, studentEmail := range commonPayload.Students {
		deleteStudentByEmail(studentEmail)
	}
}

func TestCommonStudentsInvalidTeacherEmail(t *testing.T) {
	// Create test students in the database
	for _, studentEmail := range studentEmailsListTestCases[1] {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Create only one teacher, making the other teacher email provided invalid
	for _, teacherEmail := range teacherEmailsListTestCases[0] {
		err := data.CreateTeacherByEmail(teacherEmail)
		assert.NoError(t, err)
		var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmail, StudentEmails: studentEmailsListTestCases[1]}
		_, _, err = CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
		assert.NoError(t, err)
	}
	
	_, status, err := CreateRequestAndReturnResponseData(nil, "GET", fmt.Sprintf("/api/commonstudents?teacher=%s&teacher=%s", 
																					teacherEmailsTestCases[0], teacherEmailsTestCases[1]))
	assert.NoError(t, err)

	// Compare expected and actual status codes
	assert.Equal(t, http.StatusNotFound, status)

	// Clean up
	for _, teacherEmail := range teacherEmailsListTestCases[0] {
		deleteTeacherByEmail(teacherEmail)
	}
	for _, studentEmail := range studentEmailsListTestCases[1] {
		deleteStudentByEmail(studentEmail)
	}
}

func TestCommonStudentsNoTeacherEmails(t *testing.T) {
	_, status, err := CreateRequestAndReturnResponseData(nil, "GET", "/api/commonstudents")
	assert.NoError(t, err)

	// Compare expected and actual status codes
	assert.Equal(t, http.StatusUnprocessableEntity, status)
}