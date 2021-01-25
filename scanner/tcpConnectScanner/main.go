package main

import (
	"fmt"
	"os"

	"myTools/scanner/tcpConnectScanner/iputil"
	"myTools/scanner/tcpConnectScanner/tcpScannerTool"
)

func main() {

	fmt.Printf("0: %s\n",os.Args[0])
	fmt.Printf("1: %s\n",os.Args[1])
	fmt.Printf("2: %s\n",os.Args[2])

	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]
		ips, err := iputil.GetIpList(ipList)
		if err != nil {
			fmt.Printf("parse ip list error: %+v\n", err)
		}
		fmt.Println(ips)
		ports, err := iputil.GetPorts(portList)

		if err != nil {
			fmt.Printf("parse ports error: %+v\n", err)
		}
		fmt.Println(ports)

		for _, ipItem := range ips {
			for _, portItem := range ports {
				_, err := tcpScannerTool.Connect(ipItem.String(),portItem)
				if err != nil {
					continue
				}
				fmt.Printf("ip: %v, port: %v is open\n", ipItem, portItem)
			}
		}
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}


// go run .\main.go 45.22.2.156,114.114.114.114 22,23,24,25,80-100