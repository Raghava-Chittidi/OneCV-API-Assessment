package tests

import (
	"Raghava/OneCV-Assignment/internal/api"
	"Raghava/OneCV-Assignment/internal/util"
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some payloads test cases
var testPayloadCases = []interface{} {
	api.RegisterPayload{TeacherEmail: teacherEmailsTestCases[0], StudentEmails: studentEmailsTestCases},
	api.SuspendPayload{StudentEmail: studentEmailsTestCases[0]},
	api.RetrieveForNotificationsPayload{TeacherEmail: teacherEmailsTestCases[2], Notification: "hello @studentone@gmail.com"},
}

// Some status code test cases
var statusCodeTestCases = []int {200, 204, 400, 422, 500}

// Test ReadRequestBody function with the different payload test cases
func TestReadRequestBody(t *testing.T) {
	for _, expected := range testPayloadCases {
		// Expected json encoding of the payload
		expected, err := json.Marshal(expected);
		assert.NoError(t, err)
		
		body := bytes.NewBuffer(expected)
		// Make a post request with the payload
		r := httptest.NewRequest("POST", "/test", body)
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
	
		var data interface{}
		// Read the data into the data parameter
		err = util.ReadRequestBody(w, r, &data)
		assert.NoError(t, err)

		actual, err := json.Marshal(data);
		assert.NoError(t, err)
		// Compare the expected and actual json strings
		assert.JSONEq(t, string(expected), string(actual))
	}
}

// Test WriteJSON function with the different payload test cases
func TestWriteJSON(t *testing.T) {
	for index, expected := range testPayloadCases {
		// Make a new request
		r := httptest.NewRequest("GET", "/test", nil)
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Write the payload onto the response body and set response status
		err := util.WriteJSON(w, expected, statusCodeTestCases[index])
		assert.NoError(t, err)

		// Get the data written in response body
		res := w.Result()
		defer res.Body.Close()
		actual, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		expected, err := json.Marshal(expected);
		assert.NoError(t, err)

		// Compare the expected and actual json strings
		assert.JSONEq(t, string(expected), string(actual))
		// Compare the expected and actual status codes
		assert.Equal(t, res.StatusCode, statusCodeTestCases[index])
	}
}
