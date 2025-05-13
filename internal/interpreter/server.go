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

func (server *Server) Upload(source string, target string) error {
	return server.Client.Upload(source, target)
}

func (server *Server) Download(source string, target string) error {
	return server.Client.Download(source, target)
}

func (localhost *Server) Delete(target string) error {
	return nil
}
