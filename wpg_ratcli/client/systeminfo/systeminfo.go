package systeminfo

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func GetSystem() (pcname string, wlanip string, lanip string) {
	pcname, _ = os.Hostname()
	lanip = getLanIp()
	wlanip = getWlanIp()
	return
}

func getLanIp() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	ip = strings.TrimSpace(ip)
	return
}

func getWlanIp() (wlanIp string) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	wlanIp = strings.Split(conn.LocalAddr().String(), ":")[0]
	return
}
