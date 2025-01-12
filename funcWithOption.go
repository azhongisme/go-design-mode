package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr         string
	Port         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	Timeout      time.Duration
}

type Option func(*Server)

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithPort(port string) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeOut = readTimeout
	}
}

func WithWrtieTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeOut = writeTimeout
	}
}

func WithTimeOut(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer2(options ...Option) *Server {
	srv := &Server{
		Addr:         "localhost",
		Port:         "8000",
		ReadTimeOut:  2 * time.Second,
		WriteTimeOut: 2 * time.Second,
		Timeout:      4 * time.Second,
	}

	for _, option := range options {
		option(srv)
	}

	return srv
}

func NewServer(addr string, port string, readTimeout, writeTimeout, timeout time.Duration) Server {
	return Server{
		Addr:         addr,
		Port:         port,
		ReadTimeOut:  readTimeout,
		WriteTimeOut: writeTimeout,
		Timeout:      timeout,
	}
}

func main() {
	server := NewServer("localhost", ":8000", 2*time.Second, 2*time.Second, 4*time.Second)
	fmt.Println(server)
	server2 := NewServer2(WithPort("8001"))
	fmt.Println(server2)
}
