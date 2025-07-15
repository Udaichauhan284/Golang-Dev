package response

import (
	"encoding/json"
	"net/http"
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