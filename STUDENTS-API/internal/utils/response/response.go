package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

//now these Status and Error is comming with
//captial letter, we can decide, in json to send them in small
type Response struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

//making the constant to return the error
const (
	StatusOk = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(status);

	return json.NewEncoder(w).Encode(data);
}


func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error : err.Error(),
	}
}

//need to create the validation error function also to handle the ValidatorError which we just imported now
func ValidatorError(errs validator.ValidationErrors) Response {
	var errMsgs []string;

	for _, err := range errs {
		switch err.ActualTag(){
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()));
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()));
		}
	}

	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs, ", "),
	}
}