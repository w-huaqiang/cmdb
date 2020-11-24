package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-ping/ping"
	ipparse "github.com/w-huaqiang/myPackage/netMgt"
)

var wg sync.WaitGroup
var usage = `
pingnet  [-c count] [-t timeout] subnet

Example:
    # ping subnet 3.1.20.0/24 default count is 2, timeout is 2 seconds
	pingnet 3.1.20.0/24

	# ping subnet 172.16.5.0/26 and count is 3 timeout 5 seconds
	pingnet -c 3 -t 5s 172.16.5.0/26
`

func goIPparse(s string, t time.Duration, c int) {
	_, network, err := net.ParseCIDR(s)
	if err == nil {
		ipList := ipparse.IPTable(network)
		length := len(ipList)
		for i := 0; i < length; i++ {
			ip := fmt.Sprint(ipList[i])
			wg.Add(1)
			go pingnet(ip, t, c)
		}
	} else {
		fmt.Println("Error: ", err)
	}

}

func pingnet(s string, t time.Duration, c int) {

	pinger, err := ping.NewPinger(s)
	if err != nil {
		panic(err)
	}
	pinger.Timeout = t
	pinger.Count = c
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv == stats.PacketsSent {
		fmt.Printf("%s-->true\n", stats.Addr)
	} else {
		fmt.Printf("%s-->false\n", stats.Addr)
	}

	defer wg.Done()
}

func main() {

	//netWorkcidr := flag.String("n", "None", "The subnet like 3.1.20.0/24")
	timeout := flag.Duration("t", time.Second*3, "timeout of ping")
	count := flag.Int("c", 2, "the count of ping per time")
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	netWorkcidr := flag.Arg(0)
	goIPparse(netWorkcidr, *timeout, *count)
	wg.Wait()

}
