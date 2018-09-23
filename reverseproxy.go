package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	origin, _ := url.Parse(target)
	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(res, req)
}
