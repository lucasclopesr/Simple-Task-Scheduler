package transport

import (
	"net"
	"net/http"
)

const UnixSocketAddress = "/var/run/insprd.sock"

// NewUnixSocketClient returna um client que conecta via unix socket
func NewUnixSocketClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			Dial: func(string, string) (net.Conn, error) {
				return net.Dial("unix", UnixSocketAddress)
			},
		},
	}
}
