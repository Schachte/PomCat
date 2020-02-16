//Backend application that handles the implementation of storing and processing
//tasks for the Pomodoro application
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/schachte/pomodoro/backend/database"
	"github.com/schachte/pomodoro/backend/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// database instance
	dbWrapper := *database.InitializeDatabase("./pomodoro.db")

	multiplexer := router.InitializeRouter(dbWrapper.Db)
	initializeRoutes(multiplexer, dbWrapper.Db)

	//TODO: Update to loading port via env variables
	if err := http.ListenAndServe(":8180", multiplexer); err != nil {
		log.Fatal(err)
	}
}

// Handles the initialization of all routes within the application to be
// registered within the custom multiplexer/router
func initializeRoutes(multiplexer *router.CustomRouter, db *sql.DB) {
	multiplexer.POST("storeTask", router.StoreTask(db))
}
