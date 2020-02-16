//Backend application that handles the implementation of storing and processing
//tasks for the Pomodoro application
package main

import (
	"log"
	"net/http"

	"github.com/schachte/pomodoro/backend/database"
	"github.com/schachte/pomodoro/backend/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbWrapper := *database.InitializeDatabase("./pomodoro.db")

	multiplexer := router.InitializeRouter(dbWrapper.Db)
	initializeRoutes(multiplexer, &dbWrapper)

	if err := http.ListenAndServe(":8180", multiplexer); err != nil {
		log.Fatal(err)
	}
}

func initializeRoutes(multiplexer *router.CustomRouter, dbWrapper *database.DB) {
	multiplexer.POST("storeTask", router.StoreTask(dbWrapper))
}
