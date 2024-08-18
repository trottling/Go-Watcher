package Database

import (
	"Go-Watcher/src"
	_ "modernc.org/sqlite" // Cgo free driver
	"time"
)

func CheckIpBlock(ip string) bool {
	// Check if the IP is blocked
	// If TIMESTAMP_TO > Current timestamp - IP is blocked
	var blockedIpsCount int
	err := src.DBConn.QueryRow("SELECT COUNT(*) FROM blocked_ips WHERE TIMESTAMP_TO > ?", time.Now().Unix()).Scan(&blockedIpsCount)
	if err != nil {
		src.Log.Errorf("Error checking IP block (%s): %s", ip, err)
		return false
	}
	return blockedIpsCount != 0
}

func InsertRequest(request src.Connection) {
	// Insert the request data into the 'requests' table
	_, err := src.DBConn.Exec(`
		INSERT INTO requests (ip_address, port, path, location, status_code, timestamp, allowed)
		VALUES (?, ?, ?, ?, ?, ?)
	`, request.IPAddress, request.Port, request.Path, request.Location, request.StatusCode, request.Timestamp)
	if err != nil {
		src.Log.Error("Error inserting request: ", err)
	}
}
