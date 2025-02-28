package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	websocketjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
	"log"
	"net"
	"net/http"
)

type WsRpcConn struct {
	ctx context.Context
	rpc *jsonrpc2.Conn
}

type RPCServer struct {
	conns []*WsRpcConn
}

func (s *RPCServer) Notify(method string, params interface{}) error {
	for _, conn := range s.conns {
		if err := conn.rpc.Notify(conn.ctx, method, params); err != nil {
			return err
		}
	}
	return nil
}

func (s *RPCServer) addConn(target *WsRpcConn) {
	s.conns = append(s.conns, target)
}

func (s *RPCServer) removeConn(target *WsRpcConn) {
	for i, conn := range s.conns {
		if conn == target {
			s.conns = append(s.conns[:i], s.conns[i+1:]...)
			return
		}
	}
}

func (s *RPCServer) Serve(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to get available port: %v", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Failed to upgrader.Upgrade: %v", err)
			return
		}
		defer c.Close()
		ctx := r.Context()
		conn := jsonrpc2.NewConn(
			ctx,
			websocketjsonrpc2.NewObjectStream(c),
			&RPCHandler{},
		)
		wsRpcConn := &WsRpcConn{
			rpc: conn,
			ctx: ctx,
		}
		s.addConn(wsRpcConn)
		<-conn.DisconnectNotify()
		s.removeConn(wsRpcConn)
	})

	log.Printf("Server listening on :%d", port)
	if err := http.Serve(listener, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
