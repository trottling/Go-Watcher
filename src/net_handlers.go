package src

import (
	"bufio"
	"encoding/json"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func OnRequestHandler(r *http.Request) (*http.Request, *http.Response) {
	IPAddress := r.RemoteAddr
	PortStr, err := strconv.Atoi(r.URL.Port())
	if err != nil {
		PortStr = 0
	}

	// <-- Request check logic with bool result -->

	// Req is legit if not blocked
	isReqLegit := !CheckIpBlock(IPAddress)

	// Send request back
	if isReqLegit {
		// Connection data will be added to database in OnResponseHandler
		return r, nil
	} else {
		// Create connection object
		connection := Connection{
			IPAddress:   IPAddress,
			Type:        r.Method,
			Port:        PortStr,
			Path:        r.URL.Path,
			Location:    r.Host,
			StatusCode:  http.StatusForbidden,
			Timestamp:   time.Now().Unix(),
			ReqHeaders:  GetReqHeaders(r),
			ReqBody:     GetReqBody(r),
			RespHeaders: map[string]string{},
			RespBody:    "",
			Allowed:     false,
		}

		// Dump connection
		connection.DumpPath = DumpConnection(connection)

		// Insert request to database as blocked
		InsertRequest(connection)

		// Return forbidden response
		return r, goproxy.NewResponse(
			r,
			goproxy.ContentTypeText,
			http.StatusForbidden,
			"Forbidden, bitch")
	}
}

func OnResponseHandler(r *http.Response) {
	// Collect response data, like location, port, etc. and process it with other requests from same IP address

	PortStr, err := strconv.Atoi(r.Request.URL.Port())
	if err != nil {
		PortStr = 0
	}

	// Create connection object
	connection := Connection{
		IPAddress:   r.Request.RemoteAddr,
		Type:        r.Request.Method,
		Port:        PortStr,
		Path:        r.Request.URL.Path,
		Location:    r.Request.Host,
		StatusCode:  r.StatusCode,
		Timestamp:   time.Now().Unix(),
		ReqHeaders:  GetReqHeaders(r.Request),
		ReqBody:     GetReqBody(r.Request),
		RespHeaders: GetRespHeaders(r),
		RespBody:    GetRespBody(r),
		Allowed:     true,
	}

	// Dump connection
	connection.DumpPath = DumpConnection(connection)

	// Insert request to database as blocked
	InsertRequest(connection)

	// Process data
}

func DumpConnection(conn Connection) (filePath string) {
	if !Config.ActivityHandler.DumpRequests {
		return "* Connection dumping disabled *"
	}

	// Check for Requests_Dump_Ignore_Regex config rule match
	for _, ignoreRegex := range RequestsDumpIgnoreRegexList {
		if ignoreRegex.MatchString(conn.Path) {
			return "* Connection dumping disabled by config rule match*"
		}
	}

	// Write connection dump to file and return dump file path
	filePath = filepath.FromSlash(NetDumpsPath + "/" + "connection_" + conn.IPAddress + "_" + conn.Type + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt")

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0755)

	// Use decoder for minimize encoding errors
	writer := bufio.NewWriter(transform.NewWriter(file, unicode.UTF8.NewDecoder()))

	if err != nil {
		Log.Error("Cannot create connection dump file: " + err.Error())
		return ""
	}
	defer func() {
		if err = file.Close(); err != nil {
			Log.Error("Cannot close connection dump file: " + err.Error())
		}
	}()

	// Write connection dump to file
	fileContent, err := json.Marshal(conn)
	if err != nil {
		Log.Error("Cannot marshal connection dump: " + err.Error())
		return ""
	}

	_, err = writer.Write(fileContent)
	if err != nil {
		Log.Error("Cannot write connection dump to file: " + err.Error())
		return ""
	}

	// Flush the writer buffer
	_ = writer.Flush()

	Log.Infof("%s : Connection writed to %s", conn.IPAddress, filePath)
	return filePath
}
