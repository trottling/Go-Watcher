package src

import (
	"github.com/panjf2000/ants/v2"
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"runtime"
	"strconv"
)

func RunProxyServer() {
	Log.Infof("Starting Proxy-Server : 127.0.0.1:%d", Config.ProxyServer.Port)

	// Setup proxy server
	ProxyServer = goproxy.NewProxyHttpServer()

	// Show requests in STDOUT
	ProxyServer.Verbose = Config.ProxyServer.ShowConnectionsSTDOUT

	// Enable MITM
	ProxyServer.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// Connect request handler
	ProxyServer.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		connectionData := ConnectionData{
			Request:  r,
			Response: nil,
			Type:     "req",
		}

		err := ConnectionPool.Invoke(connectionData)

		if err != nil {
			Log.Error("Connection Handling Error: " + err.Error())
			return nil, nil
		}

		return r, nil
	})

	// Run proxy server
	Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.ProxyServer.Port), ProxyServer))
}

func InitConnectionsPool() {
	Log.Infof("Initializing connections pool : %d CPUs : %d threads", runtime.NumCPU(), Config.ProxyServer.Threads)
	var err error
	ConnectionPool, err = ants.NewMultiPoolWithFunc(runtime.NumCPU(),
		Config.ProxyServer.Threads,
		func(i interface{}) {
			HandleConnection(i)
		},
		ants.RoundRobin,
		ants.WithLogger(Log),
		ants.WithPreAlloc(Config.ProxyServer.PreAllocateMemory))
	if err != nil {
		Log.Panic("Cannot create connections pool for Proxy-Server: " + err.Error())
	}
}
