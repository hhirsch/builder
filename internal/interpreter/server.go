package interpreter

import (
	"github.com/melbahja/goph"
)

type Server struct {
	Client *goph.Client
}

func NewServer(client *goph.Client) *Server {
	return &Server{
		Client: client,
	}
}

func (server *Server) Execute(command string) (string, error) {
	result, err := server.Client.Run(command)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
