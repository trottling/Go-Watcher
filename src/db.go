package src

import (
	"database/sql"
)

func GetDB() *sql.DB {
	Log.Info("Loading DB...")
	// Open the database connection
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		Log.Fatal("Error opening database: " + err.Error())
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		Log.Fatal("Error pinging database: " + err.Error())
	}
	Log.Info("Database connected successfully")
	return db
}
