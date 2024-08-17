package src

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func ConnectDB() {
	Log.Info("Loading DB...")
	// Open the database connection
	var err error
	DBConn, err = sql.Open("sqlite", DBPath)
	if err != nil {
		Log.Fatal("Error opening database: " + err.Error())
	}

	// Check the database connection
	err = DBConn.Ping()
	if err != nil {
		Log.Fatal("Error pinging database: " + err.Error())
	}

	// Get records count
	var count int
	err = DBConn.QueryRow("SELECT COUNT(*) FROM connections").Scan(&count)
	if err != nil {
		Log.Fatal("Error getting records count: " + err.Error())
	}

	Log.Infof("Database connected successfully : %d records in 'connections' table", count)
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
