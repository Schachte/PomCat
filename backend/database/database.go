package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gchaincl/dotsql"
)

type DB struct {
	Db   *sql.DB
	DSql *dotsql.DotSql
}

func InitializeDatabase(databaseName string) *DB {
	db, _ := sql.Open("sqlite3", databaseName)

	mergedData := loadSqlFiles()
	dbInstance := &DB{db, mergedData}

	initializeTaskTable(dbInstance)

	return dbInstance
}

func loadSqlFiles() *dotsql.DotSql {
	tables, err := dotsql.LoadFromFile("./sql/tables.sql")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	queries, err := dotsql.LoadFromFile("./sql/queries.sql")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mergedData := dotsql.Merge(tables, queries)
	return mergedData
}

func initializeTaskTable(dbInstance *DB) {
	fmt.Println("Initializing tasks table")

	db := (*dbInstance).Db
	dsql := (*dbInstance).DSql

	statement, err := dsql.Prepare(db, "create-tasks-table")

	if err != nil {
		os.Exit(1)
	}

	statement.Exec()
	fmt.Println("Tasks table initialized!")
}
