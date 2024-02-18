package tests

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuspendValidInputs(t *testing.T) {
	// Create test payload
	err := data.CreateStudentByEmail(studentEmailsTestCases[0])
	assert.NoError(t, err)
	student, err := data.GetStudentByEmail(studentEmailsTestCases[0])
	assert.NoError(t, err)
	// Initially, student is not be suspended
	assert.Equal(t, false, student.Suspended)
	var suspendPayload = api.SuspendPayload{StudentEmail: studentEmailsTestCases[0]}

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(suspendPayload, "POST", "/api/suspend")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, statusCode)
	// Student should now be suspended
	assert.Equal(t, false, student.Suspended)
	
	// Clean up
	deleteStudentByEmail(studentEmailsTestCases[0])
}

func TestSuspendInvalidStudentEmail(t *testing.T) {
	// Student with this email does not exist yet
	var suspendPayload = api.SuspendPayload{StudentEmail: studentEmailsTestCases[0]}

	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(suspendPayload, "POST", "/api/suspend")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestSuspendInvalidPayload(t *testing.T) {
	// Compare expected and actual status codes
	_, statusCode, err := CreateRequestAndReturnResponseData(nil, "POST", "/api/suspend")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}