package xterm

import (
	"golang.org/x/crypto/ssh"
)

type ptyRequestMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}

type Terminal struct {
	Columns uint32 `json:"cols"`
	Rows    uint32 `json:"rows"`
}

type SSHClient struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IpAddress string `json:"ipaddress"`
	Port      int    `json:"port"`
	Session   *ssh.Session
	Client    *ssh.Client
	channel   ssh.Channel
}

func NewSSHClient() *SSHClient {
	return &SSHClient{}
}

func (client *SSHClient) Auth(username string, port int) {
	client.Username = username
	client.Port = port
}
