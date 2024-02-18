package api

// Payloads for all the requests recieved

type RegisterPayload struct {
	TeacherEmail string`json:"teacher"`
	StudentEmails []string `json:"students"`
}

type SuspendPayload struct {
	StudentEmail string `json:"student"`
}

type RetrieveForNotificationsPayload struct {
	TeacherEmail string `json:"teacher"`
	Notification string `json:"notification"`
}

// Structs for all the responses to be sent

type Response struct {
	Message string `json:"message,omitempty"`
}

type CommonStudentsResponse struct {
	Students []string `json:"students"`
}

type RetrieveForNotificationsResponse struct {
	Recipients []string `json:"recipients"`
}
