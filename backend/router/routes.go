// This package handles the internal business logic for processing/storing/retrieving pomodoro
// tasks registered from a particular user
package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/schachte/pomodoro/backend/database"
	"github.com/schachte/pomodoro/backend/entities"
	"github.com/schachte/pomodoro/backend/utilities"
)

func StoreTask(dbWrapper *database.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()

		//TODO: Err handling
		if err != nil {
			// handle error
		}

		newTask := new(entities.Task)

		decoder := utilities.GenerateDecoder()
		fmt.Println(r.Form)
		err = decoder.Decode(newTask, r.Form)

		//TODO: Err handling
		if err != nil {
			fmt.Println(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(newTask)
	}
}
