package server

import (
	"fmt"
	"strings"

	"github.com/desertbit/grumble"
	"wgp.ratserver/public"
)

var SessionCommand *grumble.Command

func init() {
	SessionCommand = &grumble.Command{
		Name: "session",
		Help: "show all webshell",
		Run: func(c *grumble.Context) error {
			showSession()
			return nil
		},
	}
	App.AddCommand(SessionCommand)

	useCommand := &grumble.Command{
		Name: "use",
		Help: "use session",
		Args: func(a *grumble.Args) {
			a.Int("index", "select the session id", grumble.Default(-1))
		},
		Run: func(c *grumble.Context) error {
			sessionId := c.Args.Int("index")
			if sessionId != -1 {
				useSession(sessionId, c)
			}
			return nil
		},
	}
	App.AddCommand(useCommand)
}

func useSession(sessionId int, c *grumble.Context) {
	if len(public.Clientlist) >= sessionId {
		public.CurrClient = public.Clientlist[sessionId]
		public.CurrClient.Sid = sessionId
		c.App.SetPrompt(fmt.Sprintf("session %d > ", sessionId))
	} else {
		fmt.Println("[-] session id is too large")
		showSession()
	}
}

func showSession() {
	fmt.Println("")
	if public.CurrClient != nil {
		fmt.Printf("[+] current session %d\n\n", public.CurrClient.Sid)
	}
	if len(public.Clientlist) == 0 {
		fmt.Println("[-] No connection is working, please keep waiting!")
	}
	for index, client := range public.Clientlist {
		fmt.Printf("[%d] HostName: %s \t LanIP:%s \t WlanIP:%s \n", index, client.PcName, strings.TrimSpace(client.LanIp), client.WlanIp)
	}
	fmt.Println("")
}
