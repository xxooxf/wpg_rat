package public

import "net"

type Client struct {
	Sid    int
	Id     string
	PcName string
	LanIp  string
	WlanIp string
	Os     string
	Conn   net.Conn
}

var (
	Clientlist = make([]*Client, 0)
	CurrClient *Client
)
