package errors

import "net/http"

type RestErr struct {
	Message string 	`json:"message"`
	Status 	int 	`json:"status"`
	Error 	string 	`json:"error"`
}

func CustomBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}