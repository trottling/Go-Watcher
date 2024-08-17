package Proxy_Server

import (
	"Go-Watcher/src"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"runtime"
)

func RunProxyServer() {
	src.Log.Infof("Starting Proxy-Server : 127.0.0.1:%d", src.Config.ProxyServer.Port)
	// Setup proxy server
	src.ProxyServer = goproxy.NewProxyHttpServer()
	src.ProxyServer.Verbose = src.Config.ProxyServer.ShowConnectionsSTDOUT
	src.ProxyServer.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	// Run proxy server
	src.Log.Fatal(http.ListenAndServe(":"+string(rune(src.Config.ProxyServer.Port)), src.ProxyServer))
}

func InitConnectionsPool() {
	src.Log.Infof("Initializing connections pool : %d CPUs : %d threads", runtime.NumCPU(), src.Config.ProxyServer.Threads)
	var err error
	src.ConnectionPool, err = ants.NewMultiPoolWithFunc(runtime.NumCPU(),
		src.Config.ProxyServer.Threads,
		func(i interface{}) {
			fmt.Println(i)
		},
		ants.RoundRobin,
		ants.WithLogger(src.Log),
		ants.WithPreAlloc(src.Config.ProxyServer.PreAllocateMemory))
	if err != nil {
		src.Log.Panic("Cannot create connections pool for Proxy-Server: " + err.Error())
	}
}
