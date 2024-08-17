package Activity_Handler

import (
	"Go-Watcher/src"
)

func HandleConnection(data interface{}) {
	connectionData := data.(src.ConnectionData)

	switch connectionData.Type {
	case "req":
		// Handle request
		request, response := OnRequest(connectionData.Request)
		if response != nil {
			connectionData.Response = response
		} else {
			connectionData.Response = nil
		}
		connectionData.Request = request
	case "res":
		// Handle response
		request, response := OnResponse(connectionData.Response)
		if request != nil {
			connectionData.Request = request
		} else {
			connectionData.Request = nil
		}
		connectionData.Response = response
	default:
		src.Log.Errorf("Unknown connection type: %s", connectionData.Type)
	}
}
