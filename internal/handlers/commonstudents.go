package handlers

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/util"
	"fmt"
	"net/http"
)

// Get emails of common students under the teachers given
func CommonStudents(w http.ResponseWriter, r *http.Request) {
	// Gets list of teacher emails from query params
	teacherEmails := r.URL.Query()["teacher"]
	// Set res status to 422 if no teacher emails are given as there will definitely be no common students among 0 teachers
	if len(teacherEmails) == 0 {
		errors.ErrorJSON(w, fmt.Errorf("no teacher emails were given"), http.StatusUnprocessableEntity)
		return
	}

	// Check if all the teacher emails given are emails of existing teachers
	err := data.CheckIfTeachersEmailsListIsValid(teacherEmails)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusNotFound)
		return
	}
	
	commonStudentEmails, err := data.GetCommonStudentEmailsUnderTeachers(teacherEmails)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Set res body to list of common student emails and res status to 200 
	res := api.CommonStudentsResponse{Students: commonStudentEmails}
	util.WriteJSON(w, res, http.StatusOK)
}