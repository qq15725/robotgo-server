package main

import (
	"context"
	"fmt"
	"github.com/sourcegraph/jsonrpc2"
	"os"
	"sync"
)

type RobotgoServer struct {
	conn      *jsonrpc2.Conn
	connMutex sync.Mutex
	shutdown  bool
}

func NewRobotgoServer() *RobotgoServer {
	return &RobotgoServer{}
}

func (s *RobotgoServer) Initialize(ctx context.Context) error {
	// to implement
	return nil
}

func (s *RobotgoServer) Handle(context.Context, *jsonrpc2.Conn, *jsonrpc2.Request) (result interface{}, err error) {
	fmt.Println("Handling request...")
	// to implement
	return nil, nil
}

func (s *RobotgoServer) Serve(ctx context.Context) {
	fmt.Println("Starting Robotgo RPC server...")

	s.conn = jsonrpc2.NewConn(
		context.Background(),
		jsonrpc2.NewBufferedStream(os.Stdin, jsonrpc2.VSCodeObjectCodec{}),
		jsonrpc2.HandlerWithError(s.Handle),
	)

	<-s.conn.DisconnectNotify()
}
