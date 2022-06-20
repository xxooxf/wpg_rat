package main

import (
	"github.com/desertbit/grumble"
	"wgp.ratserver/server"
)

func main() {
	grumble.Main(server.App)
}
