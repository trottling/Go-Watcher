package src

import (
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"strconv"
	"time"
)

func HandleConnection(data interface{}) ConnectionData {
	connectionData := data.(ConnectionData)

	switch connectionData.Type {
	case "req":
		// Handle request
		request, response := OnRequest(connectionData)
		if response != nil {
			connectionData.Response = response
		} else {
			connectionData.Response = nil
		}
		connectionData.Request = request
	case "res":
		// Handle response
		request, response := OnResponse(connectionData)
		if request != nil {
			connectionData.Request = request
		} else {
			connectionData.Request = nil
		}
		connectionData.Response = response
	default:
		Log.Errorf("Unknown connection type: %s", connectionData.Type)
	}
	return connectionData
}

func OnRequest(data ConnectionData) (*http.Request, *http.Response) {
	IPAddress := data.Request.RemoteAddr
	Port := data.Request.URL.Port()
	PortStr, _ := strconv.Atoi(Port)
	Path := data.Request.URL.Path
	Location := data.Request.Host

	// <-- Request check logic with bool result -->
	isReqLegit := CheckIpBlock(IPAddress)

	// Send request back
	if isReqLegit {
		return data.Request, nil
	} else {

		// Add request to database as blocked
		InsertRequest(Connection{
			IPAddress:  IPAddress,
			Port:       PortStr,
			Path:       Path,
			Location:   Location,
			StatusCode: http.StatusForbidden,
			Timestamp:  time.Now().Unix(),
			Allowed:    false,
		})

		// Return forbidden response
		return data.Request, goproxy.NewResponse(data.Request,
			goproxy.ContentTypeText, http.StatusForbidden,
			"Forbidden, bitch")
	}
}

func OnResponse(data interface{}) (*http.Request, *http.Response) {
	// dd
	return nil, nil
}
