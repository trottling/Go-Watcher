package src

import (
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
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
		// OnRequestHandler will block or allow request
		return OnRequestHandler(r)
	})

	// Connect response handler
	ProxyServer.OnResponse().DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		// OnResponseHandler will not change response
		OnResponseHandler(r)
		return r
	})

	// Run proxy server
	Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.ProxyServer.Port), ProxyServer))
}
