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
	err = DBConn.QueryRow("SELECT COUNT(*) FROM blocked_ips").Scan(&blockedIpsCount)
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

	if Config.ProxyServer.ShowConnectionsSTDOUT {
		Log.Infof("Inserting request: %+v", request)
	}

	_, err := DBConn.Exec(`
		INSERT INTO connections (ip_address, port, path, location, status_code, timestamp, allowed, type, dump_path)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, request.IPAddress, request.Port, request.Path, request.Location, request.StatusCode, request.Timestamp, request.Allowed, request.Type, request.DumpPath)
	if err != nil {
		Log.Error("Error inserting request: ", err)
	}
}

func GetIPConnections(ip string, timestamp int64) []Connection {
	// Get all connections from IP
	rows, err := DBConn.Query("SELECT * FROM connections WHERE ip_address = ? AND timestamp > ?", ip, timestamp)
	if err != nil {
		Log.Errorf("Error getting IP connections count (%s): %s", ip, err)
		return nil
	}

	var connections []Connection

	for rows.Next() {
		var connection Connection
		err := rows.Scan(&connection.IPAddress, &connection.Port, &connection.Path, &connection.Location, &connection.StatusCode, &connection.Timestamp, &connection.Allowed, &connection.Type, &connection.DumpPath)
		if err != nil {
			Log.Error("Error getting IP connections count: ", err)
			continue
		}
		connections = append(connections, connection)
	}

	return connections
}

func BlockIP(ip string, reason string) {
	// Block IP
	_, err := DBConn.Exec(`
		INSERT INTO blocked_ips (ip_address, timestamp_from, timestamp_to, reason)
		VALUES (?, ?, ?, ?)
	`, ip, time.Now().Unix(), time.Now().Unix()+int64(Config.ActivityHandler.BlockIPsTime), reason)
	if err != nil {
		Log.Error("Error inserting request: ", err)
	}
	Log.Infof("IP %s blocked for %d seconds, reason: %s", ip, Config.ActivityHandler.BlockIPsTime, reason)
}
