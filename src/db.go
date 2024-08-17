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

func InsertRequest(request Request) {
	// Insert the request data into the 'requests' table
	_, err := DBConn.Exec(`
		INSERT INTO requests (ip_address, protocol, port, path, location, status_code, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, request.IPAddress, request.Protocol, request.Port, request.Path, request.Location, request.StatusCode, request.Timestamp)
	if err != nil {
		Log.Error("Error inserting request: ", err)
	}
}
