package Activity_Handler

import (
	"Go-Watcher/src"
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"strconv"
	"time"
)

func OnRequest(data src.ConnectionData) (*http.Request, *http.Response) {
	isReqLegit := true
	IPAddress := data.Request.RemoteAddr
	Port := data.Request.URL.Port()
	PortStr, _ := strconv.Atoi(Port)
	Path := data.Request.URL.Path
	Location := data.Request.Host

	// <-- Request check logic with bool result -->
	isReqLegit = src.CheckIpBlock(IPAddress)

	// Send request back
	if isReqLegit {
		return data.Request, nil
	} else {

		// Add request to database as blocked
		src.InsertRequest(src.Connection{
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
