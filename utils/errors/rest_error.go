package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func CustomBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}

func CustomNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func CustomInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
