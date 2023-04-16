package server

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/alichaddad/goresp/pkg/resp"
)

const NETWORK = "tcp"

type Server struct {
	addr       string
	httpClient *http.Client
}

func New(addr string) *Server {
	return &Server{
		addr:       addr,
		httpClient: &http.Client{Timeout: time.Duration(2 * time.Second)},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting Server on %s", s.addr)
	listener, err := net.Listen(NETWORK, s.addr)
	if err != nil {
		return err
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.handleRequest(conn)
	}
}
func (s *Server) handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		if _, err := reader.Peek(1); err == io.EOF {
			break
		}
		values := resp.ParseMessage(reader)
		switch values[0] {
		case "myping":
			msg := resp.NewBulkMessage(s.MyPing())
			conn.Write(msg.GetValue())
		case "testurl":
			if len(values) < 2 {
				conn.Write(resp.NewError(errors.New("missing command argument")).GetValue())
				break
			}
			res := "true"
			err := s.TestUrl(values[1])
			if err != nil {
				res = "false"
			}
			conn.Write(resp.NewBulkMessage(res).GetValue())
		default:
			conn.Write(resp.NewError(errors.New("unknown command")).GetValue())
		}

	}
	conn.Close()
}

func (s *Server) MyPing() string {
	return "pong"
}

func (s *Server) TestUrl(url string) error {
	_, err := s.httpClient.Get(url)
	return err
}
