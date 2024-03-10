package functional_patterns

import (
	"fmt"
	"time"
)

type simpleServer struct {
	host    string
	port    string
	timeout time.Duration
	maxConn int
}

func (ss *simpleServer) String() string {
	return fmt.Sprintf("Host: %v\nPost: %v\nTimeout: %v minutes\nMax Connections: %v", ss.host, ss.port, ss.timeout.Minutes(), ss.maxConn)
}

func (s *simpleServer) start() {
	fmt.Println("Starting server...")
	fmt.Println("********************************")
	fmt.Println(s)
	fmt.Println("********************************")
	fmt.Println("Server started.")
}

func newServer(options ...func(*simpleServer)) *simpleServer {
	server := &simpleServer{}

	for _, o := range options {
		o(server)
	}

	return server
}

func withHost(host string) func(*simpleServer) {
	return func(ss *simpleServer) {
		ss.host = host
	}
}

func withPort(port string) func(*simpleServer) {
	return func(ss *simpleServer) {
		ss.port = port
	}
}

func withTimeout(timeout time.Duration) func(*simpleServer) {
	return func(ss *simpleServer) {
		ss.timeout = timeout
	}
}

func withMaxConn(maxConn int) func(*simpleServer) {
	return func(ss *simpleServer) {
		ss.maxConn = maxConn
	}
}

func OptionalPattern() {
	svr := newServer(
		withHost("localhost"),
		withPort("2233"),
		withTimeout(time.Minute),
		withMaxConn(120),
	)

	svr.start()
}
