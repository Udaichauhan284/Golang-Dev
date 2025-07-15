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
)

//creation of new student
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		var student types.Student

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

		slog.Info("Creating a student");

		response.WriteJson(w, http.StatusCreated, map[string]string{"success" : "OK"});


		// w.Write([]byte("Welcome To Students APIs"));
	}
}