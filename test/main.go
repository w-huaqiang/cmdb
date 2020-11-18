package main

import (
	"flag"
	"fmt"
	"net"

	word "github.com/w-huaqiang/myPackage/fileMgt"
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

func testword() {

	flag.Parse()
	argument := flag.Args()
	if len(argument) == 0 {
		fmt.Printf("useage: textcont file\n")
		return
	}

	a, err := word.TextCount(argument[0])
	if err != nil {
		fmt.Println(err)
	} else {
		for i, key := range a.Keys {
			fmt.Printf("%s:%d\n", key, a.Values[i])
		}
	}

}

func main() {
	//testIPparse()
	testword()
}
