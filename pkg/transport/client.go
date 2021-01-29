package transport

import (
	"net"
	"net/http"
)

// UnixSocketAddress é o endereço de comunicação em unix socket
const UnixSocketAddress = "/tmp/simpd.sock"

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
