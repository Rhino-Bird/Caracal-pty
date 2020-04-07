package xterm

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ShellWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		msg := base64.StdEncoding.EncodeToString(message)
		log.Printf(msg, "msg is")

		cols := "150"
		rows := "35"
		col, _ := strconv.Atoi(cols)
		row, _ := strconv.Atoi(rows)
		terminal := Terminal{
			Columns: uint32(col),
			Rows:    uint32(row),
		}
		sshClient, err := DecodedMsgToSSHClient(msg)
		if err != nil {
			log.Printf("Decode err is", err)
		}
		if sshClient.IpAddress == "" || sshClient.Password == "" {
			log.Printf("Auth err is", err)
		}
		err = sshClient.GenerateClient()
		if err != nil {
			log.Printf("Generate client err is", err)
		}
		sshClient.RequestTerminal(terminal)
		sshClient.Connect(c)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Printf("write:", err)
			break
		}
	}
}
