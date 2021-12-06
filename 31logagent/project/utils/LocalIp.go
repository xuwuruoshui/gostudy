package utils

import (
	"net"
	"strings"
)

func GetLocalIP() (ip string){
	conn,err := net.Dial("udp","8.8.8.8:80")
	if err!=nil{
		return
	}
	defer  conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.IP.String(),":")[0]
	return ip
}