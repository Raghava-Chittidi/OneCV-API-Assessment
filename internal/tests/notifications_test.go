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

func TestRetrieveNotificationsValidInputs(t *testing.T) {
	// Create 3 test students in the database
	for _, studentEmail := range studentEmailsListTestCases[2] {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Create test teacher in the database and register first 2 students to the teacher
	err := data.CreateTeacherByEmail(teacherEmailsTestCases[0])
	assert.NoError(t, err)
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsListTestCases[1]}
	_, _, err = CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)

	// Create test payload
	var retrieveNotificationsPayload = api.RetrieveForNotificationsPayload{TeacherEmail: teacherEmailsTestCases[0], 
																		   Notification: fmt.Sprintf("Hello students! @%s", studentEmailsTestCases[2])}
				
	// Expected response body is all 3 students as 2 students were registered to the teacher and the last student was tagged in the notification																	   
	var expected = api.RetrieveForNotificationsResponse{Recipients: studentEmailsListTestCases[2]}

	// Compare expected and actual status codes
	data, statusCode, err := CreateRequestAndReturnResponseData(retrieveNotificationsPayload, "POST", "/api/retrievefornotifications")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Compare the expected and actual slice of recipients recieved
	var actualData api.RetrieveForNotificationsResponse
	json.Unmarshal(data, &actualData)
	assert.ElementsMatch(t, expected.Recipients, actualData.Recipients)

	// Clean up
	deleteTeacherByEmail(teacherEmailsTestCases[0])
	for _, studentEmail := range studentEmailsListTestCases[2] {
		deleteStudentByEmail(studentEmail)
	}
}

func TestRetrieveNotificationsInvalidTeacherEmail(t *testing.T) {
	// Create 3 test students in the database
	for _, studentEmail := range studentEmailsListTestCases[2] {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Create test payload. No teacher is created to ensure that the teacher email provided is invalid
	var retrieveNotificationsPayload = api.RetrieveForNotificationsPayload{TeacherEmail: teacherEmailsTestCases[0], 
																		   Notification: fmt.Sprintf("Hello students! @%s", studentEmailsTestCases[2])}

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(retrieveNotificationsPayload, "POST", "/api/retrievefornotifications")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)

	// Clean up
	for _, studentEmail := range studentEmailsListTestCases[2] {
		deleteStudentByEmail(studentEmail)
	}
}

func TestRetrieveNotificationsInvalidStudentEmail(t *testing.T) {
	// Create only 2 test students in the database
	for _, studentEmail := range studentEmailsListTestCases[1] {
		err := data.CreateStudentByEmail(studentEmail)
		assert.NoError(t, err)
	}

	// Create test teacher in the database and register these 2 students to the teacher
	err := data.CreateTeacherByEmail(teacherEmailsTestCases[0])
	assert.NoError(t, err)
	var registerPayload = api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsListTestCases[1]}
	_, _, err = CreateRequestAndReturnResponseData(registerPayload, "POST", "/api/register")
	assert.NoError(t, err)

	// Create test payload and since the 3rd student was not created, tagging them would make it invalid
	var retrieveNotificationsPayload = api.RetrieveForNotificationsPayload{TeacherEmail: teacherEmailsTestCases[0], 
																		   Notification: fmt.Sprintf("Hello students! @%s", studentEmailsTestCases[2])}
				
	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(retrieveNotificationsPayload, "POST", "/api/retrievefornotifications")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)

	// Clean up
	deleteTeacherByEmail(teacherEmailsTestCases[0])
	for _, studentEmail := range studentEmailsListTestCases[1] {
		deleteStudentByEmail(studentEmail)
	}
}

func TestRetrieveNotificationsInvalidPayload(t *testing.T) {
	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(nil, "POST", "/api/retrievefornotifications")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}