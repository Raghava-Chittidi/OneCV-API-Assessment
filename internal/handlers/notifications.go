package handlers

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/util"
	"net/http"
	"regexp"
	"strings"
)

// Returns a response with all the emails of the students who should be notified as json data
func RetrieveForNotifications(w http.ResponseWriter, r *http.Request) {
	var payload api.RetrieveForNotificationsPayload
	err := util.ReadRequestBody(w, r, &payload)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Gets preloaded teacher with RegisteredStudents from database when given an email
	teacher, err := data.GetPreloadedTeacherByEmail(payload.TeacherEmail)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	// Create a hashmap and add the emails of all the students who are registered to the teacher and are not suspended
	// A hashmap is used to remove any possible duplicates
	hashmap := make(map[string]bool)
	for _, student := range teacher.RegisteredStudents {
		if !student.Suspended {
			hashmap[student.Email] = true
		}
	}

	// Finds all the strings that match the regex 
	re := regexp.MustCompile(`@student.*@gmail.com`)
	matched := re.FindAllString(payload.Notification, -1)

	// If no matches, it means that no other students should be notified
	if len(matched) != 0 {
		// Generate list of emails of all the tagged students
		var taggedStudentEmails []string
		for _, email := range strings.Split(matched[0], " ")  {
			validEmailAddress := email[1:]
			taggedStudentEmails = append(taggedStudentEmails, validEmailAddress)
		}

		// Check if all the student emails given are emails of existing students
		err = data.CheckIfStudentsEmailsListIsValid(taggedStudentEmails)
		if err != nil {
			errors.ErrorJSON(w, err, http.StatusNotFound)
			return
		}
	
		taggedStudents, err := data.GetStudentsByEmailsList(taggedStudentEmails)
		if err != nil {
			errors.ErrorJSON(w, err, http.StatusNotFound)
			return
		}
	
		// Add the emails of all the students who are tagged and are not suspended
		for _, student := range taggedStudents {
			if !student.Suspended {
				hashmap[student.Email] = true
			}
		}
	}

	// Generate a list of emails of all the recipients who should be notified
	var allRecipientEmails []string
	for email := range hashmap {
		allRecipientEmails = append(allRecipientEmails, email)
	}

	// Set res body to list of recipient emails and res status to 200 
	res := api.RetrieveForNotificationsResponse{Recipients: allRecipientEmails}
	util.WriteJSON(w, res, http.StatusOK)
}