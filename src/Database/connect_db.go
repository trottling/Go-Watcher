package Database

import (
	"Go-Watcher/src"
	"database/sql"
	_ "modernc.org/sqlite" // Cgo free driver
)

func ConnectDB() {
	src.Log.Info("Loading DB...")
	// Open the database connection
	var err error
	src.DBConn, err = sql.Open("sqlite", src.DBPath)
	if err != nil {
		src.Log.Fatal("Error opening database: " + err.Error())
	}

	// Check the database connection
	err = src.DBConn.Ping()
	if err != nil {
		src.Log.Fatal("Error pinging database: " + err.Error())
	}

	// Get records count in 'connections' table
	var connectionsCount int
	err = src.DBConn.QueryRow("SELECT COUNT(*) FROM connections").Scan(&connectionsCount)
	if err != nil {
		src.Log.Fatal("Error getting records count: " + err.Error())
	}

	// Get records count in 'connections' table
	var blockedIpsCount int
	err = src.DBConn.QueryRow("SELECT COUNT(*) FROM connections").Scan(&blockedIpsCount)
	if err != nil {
		src.Log.Fatal("Error getting records count: " + err.Error())
	}

	src.Log.Info("Database connected successfully")
	src.Log.Infof("%d records in 'connections' table", connectionsCount)
	src.Log.Infof("%d records in 'blocked_ips' table", blockedIpsCount)
}
