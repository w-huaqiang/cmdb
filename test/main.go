package main

import (
	"fmt"
	"net"

	ipparse "github.com/w-huaqiang/myPackage/netMgt"
)

func testIPparse() {
	s := "3.1.20.0/24"
	_, network, err := net.ParseCIDR(s)
	if err == nil {
		ipList := ipparse.IPTable(network)
		length := len(ipList)
		fmt.Println("最小IP为: ", ipList[0])
		fmt.Println("最大IP为： ", ipList[length-1])
		fmt.Println("子网IP共: ", length)
		fmt.Println(ipList)
	} else {
		fmt.Println("Error: ", err)
	}

}

func main() {
	testIPparse()
}
