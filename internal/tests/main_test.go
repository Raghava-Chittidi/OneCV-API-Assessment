package tests

import (
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/router"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var TestRouter = router.Setup()

// Connects to db and runs all the tests defined
func TestMain(m *testing.M) {
	_, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln("Cannot connect to database", err)
	}

	os.Exit(m.Run())
}

// Creates a mock request with payload, method and route url. Returns body, status code of response and any possible error
func CreateRequestAndReturnResponseData(payload interface{}, method string, route string) ([]byte, int, error) {
	var req *http.Request
	// Set body to nil if not given any payload else, set req body to the payload
	if payload == nil {
		req = httptest.NewRequest(method, route, nil)
	} else {
		payloadBytes, err := json.Marshal(payload);
		if err != nil {
			return nil, -1, err
		}
		buffer := bytes.NewBuffer(payloadBytes)
		req = httptest.NewRequest(method, route, buffer)
	}
	
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Writes data and status code to the ResponseWriter according to whichever handler is triggered when request goes to that route 
	TestRouter.ServeHTTP(w, req)
	return w.Body.Bytes(), w.Result().StatusCode, nil
}