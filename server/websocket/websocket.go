package websocket

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}
var connections []*websocket.Conn

func Upgrade(writer http.ResponseWriter, request *http.Request) error {
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return err
	}

	connections = append(connections, connection)
	connection.SetCloseHandler(func(code int, text string) error {
		logrus.Debugf("disconnected %s", text)
		return nil
	})
	logrus.Debugf("ws: added a new connection for %s", connection.RemoteAddr())
	return nil
}

func removeClient(ws *websocket.Conn) {
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

func Broadcast(mtype int, message string) {
	for _, ws := range connections {
		logrus.Debugf("sending update to %s", ws.RemoteAddr())

		go func(ws *websocket.Conn) {
			err := <-checkWebSocket(ws)
			if err != nil {
				removeClient(ws)
				logrus.Debugf("err -> %s", err)
			}
			<-time.After(time.Second * 3)
			ws.WriteMessage(mtype, []byte(message))
		}(ws)
	}
}
