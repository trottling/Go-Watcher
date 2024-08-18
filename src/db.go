package src

import (
	"database/sql"
	_ "modernc.org/sqlite" // Cgo free driver
	"time"
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

	// Get records count in 'connections' table
	var connectionsCount int
	err = DBConn.QueryRow("SELECT COUNT(*) FROM connections").Scan(&connectionsCount)
	if err != nil {
		Log.Fatal("Error getting records count: " + err.Error())
	}

	// Get records count in 'connections' table
	var blockedIpsCount int
	err = DBConn.QueryRow("SELECT COUNT(*) FROM connections").Scan(&blockedIpsCount)
	if err != nil {
		Log.Fatal("Error getting records count: " + err.Error())
	}

	Log.Info("Database connected successfully")
	Log.Infof("%d records in 'connections' table", connectionsCount)
	Log.Infof("%d records in 'blocked_ips' table", blockedIpsCount)
}

func CheckIpBlock(ip string) bool {
	// Check if the IP is blocked
	// If TIMESTAMP_TO > Current timestamp - IP is blocked (will return true)
	var blockedIpsCount int
	err := DBConn.QueryRow("SELECT COUNT(*) FROM blocked_ips WHERE TIMESTAMP_TO > ?", time.Now().Unix()).Scan(&blockedIpsCount)
	if err != nil {
		Log.Errorf("Error checking IP block (%s): %s", ip, err)
		return false
	}
	if blockedIpsCount > 0 {
		Log.Warnf("IP %s is blocked", ip)
		return true
	} else {
		Log.Infof("IP %s is not blocked", ip)
		return false
	}
}

func InsertRequest(request Connection) {
	// Insert the request data into the 'requests' table
	Log.Infof("Inserting request: %+v", request)
	_, err := DBConn.Exec(`
		INSERT INTO connections (ip_address, port, path, location, status_code, timestamp, allowed)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, request.IPAddress, request.Port, request.Path, request.Location, request.StatusCode, request.Timestamp, request.Allowed)
	if err != nil {
		Log.Error("Error inserting request: ", err)
	}
}
