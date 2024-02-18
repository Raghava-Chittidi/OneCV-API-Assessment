package handlers

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/util"
	"fmt"
	"net/http"
)

// Registers all the students given to the teacher given
func Register(w http.ResponseWriter, r *http.Request) {
	var payload api.RegisterPayload
	err := util.ReadRequestBody(w, r, &payload)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	teacher, err := data.GetTeacherByEmail(payload.TeacherEmail)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	// Set res status to 422 if no student emails are given as a teacher cannot register 0 students
	if len(payload.StudentEmails) == 0 {
		errors.ErrorJSON(w, fmt.Errorf("no student emails were given"), http.StatusUnprocessableEntity)
		return
	}

	// Check if all the student emails given are emails of existing students
	err = data.CheckIfStudentsEmailsListIsValid(payload.StudentEmails)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	newStudents, err := data.GetStudentsByEmailsList(payload.StudentEmails)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Registers all the students whose emails were given to the given teacher
	err = database.DB.Model(&teacher).Association("RegisteredStudents").Append(newStudents)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Set res status to 204 
	util.WriteJSON(w, nil, http.StatusNoContent)
}