package bootstrap

import (
	server "github.com/franciscoruizar/quasar-fire/internal/core/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
