// This package handles the internal business logic for processing/storing/retrieving pomodoro
// tasks registered from a particular user
package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/schachte/pomodoro/backend/entities"
)

func StoreTask(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Printf("Error processing the following request %s", err)
			return
		}

		taskName := r.PostFormValue("name")

		startTime, _ := time.Parse(time.RFC3339, r.PostFormValue("start_time"))
		endTime, _ := time.Parse(time.RFC3339, r.PostFormValue("end_time"))
		workLength, _ := strconv.Atoi(r.PostFormValue("work_length"))
		breakLength, _ := strconv.Atoi((r.PostFormValue("break_length")))

		newTask := new(entities.Task)
		newTask.SetName(taskName)
		newTask.SetStart(startTime)
		newTask.SetEnd(endTime)
		newTask.SetWorkLength(workLength)
		newTask.SetBreakLength(breakLength)

		statement, _ := db.Prepare("INSERT INTO tasks (task_name, start_time, end_time, work_length, break_length) VALUES (?, ?, ?, ?, ?)")
		statement.Exec(newTask.Name, newTask.StartTime.String(), newTask.EndTime.String(), strconv.Itoa(newTask.WorkLength), strconv.Itoa(newTask.BreakLength))

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(newTask)
	}
}
