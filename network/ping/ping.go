package main

import (
	"log"
	"net"
	"strings"
	"time"
	"zhongying/util/fastping"
)

func main() {
	ip := "10.96.38.4"
	active_status := ping(ip)
	log.Println("%s---%s", ip, active_status)
}

func ping(ip string) bool {

	alive := false

	p := fastping.NewPinger()

	netProto := "ip4:icmp"
	if strings.Index(ip, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, ip)
	if err != nil {
		return false
	}

	p.AddIPAddr(ra)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		log.Println("IP Addr: " + addr.String() + " receive, RTT: " + rtt.String())
		alive = true
	}

	//p.OnIdle = func() {
	//}

	//p.OnErr = func(addr *net.IPAddr, t int) {
	//	//fmt.Printf("Error %s : %d\n", addr.IP.String(), t)
	//}

	p.MaxRTT = time.Second
	err = p.Run()
	if err != nil {
		return false
	}

	//fmt.Printf("%s : %v\n", ip, p.Alive)

	return alive
}
