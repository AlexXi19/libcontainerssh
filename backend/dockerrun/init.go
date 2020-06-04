package dockerrun

import (
	"containerssh/backend"
	"context"
	"github.com/docker/docker/client"
)

func createSession(sessionId string, username string) (backend.Session, error) {
	cli, err := client.NewClient("tcp://127.0.0.1:2375", "", nil, make(map[string]string))
	if err != nil {
		return nil, err
	}

	session := &dockerRunSession{}
	session.sessionId = sessionId
	session.username = username
	session.env = map[string]string{}
	session.cols = 80
	session.rows = 25
	session.pty = false
	session.containerId = ""
	session.client = cli
	session.ctx = context.Background()

	return session, nil
}

type dockerRunSession struct {
	username    string
	sessionId   string
	env         map[string]string
	cols        uint
	rows        uint
	width       uint
	height      uint
	pty         bool
	containerId string
	ctx         context.Context
	client      *client.Client
}

func Init(registry *backend.Registry) {
	dockerRunBackend := backend.Backend{}
	dockerRunBackend.Name = "dockerrun"
	dockerRunBackend.CreateSession = createSession
	registry.Register(dockerRunBackend)
}
