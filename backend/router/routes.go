// This package handles the internal business logic for processing/storing/retrieving pomodoro
// tasks registered from a particular user
package router

import (
	"net/http"

	"github.com/schachte/pomodoro/backend/database"
	"github.com/schachte/pomodoro/backend/entities"
	"github.com/schachte/pomodoro/backend/utilities"
)

func StoreTask(dbWrapper *database.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := &utilities.DynamicStruct{&entities.Task{}, rw, r}
		utilities.PersistRequest(data)
	}
}
