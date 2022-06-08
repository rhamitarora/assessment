package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func Test_socketHandler(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(socketHandler))
	defer s.Close()
	wsURL := "ws" + strings.TrimPrefix(s.URL, "http")
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	// Write a message.
	c.WriteMessage(websocket.TextMessage, []byte("TestServer"))

	// Expect the server to echo the message back.
	c.SetReadDeadline(time.Now().Add(time.Second * 2))
	mt, msg, err := c.ReadMessage()
	if err != nil {
		t.Error(err)
	}
	if mt != websocket.TextMessage || string(msg) != "TestServer" {
		t.Errorf("expected text TestServer, got %d: %s", mt, msg)
	}
}
