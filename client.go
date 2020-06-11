package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

var customHTTPTransport = &http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	DialContext:           timeoutDialer(5000, 5000),
	MaxIdleConns:          1,
	MaxIdleConnsPerHost:   1,
	IdleConnTimeout:       1 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	ExpectContinueTimeout: 1 * time.Second,
}

func timeoutDialer(connectionTimeout, readTimeout int) func(ctx context.Context, net, addr string) (c net.Conn, err error) {
	cTimeout := time.Duration(connectionTimeout) * time.Millisecond
	rTimeout := time.Duration(readTimeout) * time.Millisecond
	return func(ctx context.Context, netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rTimeout))
		return conn, nil
	}
}

var httpClient = http.Client{
	Transport: customHTTPTransport,
}

func main() {
	for {
		if true {
			resp, err := httpClient.Get("http://localhost:8000/idle")
			log.Print(resp, err)
		}
	}
}
