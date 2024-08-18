package src

import (
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"strconv"
	"time"
)

func OnRequestHandler(r *http.Request) (*http.Request, *http.Response) {
	IPAddress := r.RemoteAddr
	PortStr, _ := strconv.Atoi(r.URL.Port())

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
			Path:       r.URL.Path,
			Location:   r.Host,
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
