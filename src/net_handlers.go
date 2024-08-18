package src

import (
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"strconv"
	"time"
)

func OnRequestHandler(r *http.Request) (*http.Request, *http.Response) {
	IPAddress := r.RemoteAddr
	Port := r.URL.Port()
	PortStr, _ := strconv.Atoi(Port)
	Path := r.URL.Path
	Location := r.Host

	// <-- Request check logic with bool result -->
	isReqLegit := CheckIpBlock(IPAddress)

	// Send request back
	if isReqLegit {
		return r, nil
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
		return r, goproxy.NewResponse(r,
			goproxy.ContentTypeText, http.StatusForbidden,
			"Forbidden, bitch")
	}
}

func OnResponseHandler(r *http.Response) (*http.Request, *http.Response) {
	// dd
	return nil, nil
}

func DumpConnection() {
}
