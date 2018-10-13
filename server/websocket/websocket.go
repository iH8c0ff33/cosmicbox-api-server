package websocket

import (
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type client struct {
	ws *websocket.Conn
	mu sync.Mutex
}

func (cli *client) send(mtype int, message string) <-chan bool {
	c := make(chan bool)

	cli.mu.Lock()

	go func(cli *client, c chan<- bool) {
		err := <-checkWebSocket(cli.ws)
		if err != nil {
			logrus.Debugf("err -> %s", err)
			c <- false
			return
		}
		c <- true

		// <-time.After(time.Second * 3)
		cli.ws.WriteMessage(mtype, []byte(message))

		cli.mu.Unlock()
	}(cli, c)

	return c
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	match, err := regexp.MatchString("https://eee.lsgalfer.it", origin)
	if !match {
		match, err = regexp.MatchString("https?://localhost:\\d+", origin)
	}
	if !match {
		match, err = regexp.MatchString("https?://192\\.168\\.1\\.\\d{1,3}:\\d+", origin)
	}
	if err == nil && match {
		return true
	}
	return false
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}
var connections []*client

// Upgrade http connection to websocket connection
func Upgrade(writer http.ResponseWriter, request *http.Request) error {
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return err
	}

	cli := &client{
		ws: connection,
		mu: sync.Mutex{},
	}

	connections = append(connections, cli)
	connection.SetCloseHandler(func(code int, text string) error {
		logrus.Debugf("disconnected %s", text)
		return nil
	})
	logrus.Debugf("ws: added a new connection for %s", connection.RemoteAddr())
	return nil
}

func removeClient(ws *client) {
	for i, a := range connections {
		if a == ws {
			connections[i] = connections[len(connections)-1]
			connections = connections[:len(connections)-1]
		}
	}
}

func checkWebSocket(ws *websocket.Conn) <-chan error {
	c := make(chan error)
	go func(c chan<- error) {
		c <- ws.WriteControl(websocket.PingMessage, []byte(""), time.Now().Add(time.Second))
	}(c)
	return c
}

// Broadcast a message to all connected clients
func Broadcast(mtype int, message string) {
	for _, cli := range connections {
		logrus.Debugf("sending update to %s", cli.ws.RemoteAddr())

		go func(cli *client) {
			success := cli.send(mtype, message)
			if !<-success {
				removeClient(cli)
			}
		}(cli)
	}
}
