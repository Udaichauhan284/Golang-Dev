package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Udaichauhan284/Golang-Dev/internal/types"
	"github.com/Udaichauhan284/Golang-Dev/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

//creation of new student
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		var student types.Student
		slog.Info("Creating a student");

		//now decode this student body with json deconder
		err := json.NewDecoder(r.Body).Decode(&student)
		//now check the error
		if errors.Is(err, io.EOF){
			//if body payload is empty, then check if that error is End of file
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err));

			//now sending the customize error
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")));

			return;
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err));
			return;
		}

		//request validation
		if err := validator.New().Struct(student); err != nil {
			// response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(err)); the err which we are passing into the ValidatorError is giving me error because it want validationError to we pass, so for that we need to type cast the err to validation error

			//type casting
			validateErrs := err.(validator.ValidationErrors);
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(validateErrs));
			return;
		}

		

		response.WriteJson(w, http.StatusCreated, map[string]string{"success" : "OK"});


		// w.Write([]byte("Welcome To Students APIs"));
	}
}