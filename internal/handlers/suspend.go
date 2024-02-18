package handlers

import (
	"Raghava/OneCV-Assignment/internal/api"
	data "Raghava/OneCV-Assignment/internal/dataaccess"
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/errors"
	"Raghava/OneCV-Assignment/internal/util"
	"net/http"
)

// Suspends the student whose email is given
func Suspend(w http.ResponseWriter, r *http.Request) {
	var payload api.SuspendPayload
	err := util.ReadRequestBody(w, r, &payload)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	student, err := data.GetStudentByEmail(payload.StudentEmail)
	if err != nil {
		errors.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	// Set the suspended boolean value of the student given to true
	student.Suspended = true
	database.DB.Save(&student)

	// Set res status to 204
	util.WriteJSON(w, nil, http.StatusNoContent)
}