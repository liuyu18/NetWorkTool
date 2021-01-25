package main

import (
	"fmt"
	"os"
	//"toos/scanner/tcpConnectScanner/util"

)

func main() {

	fmt.Printf("0: %s\n",os.Args[0])
	fmt.Printf("1: %s\n",os.Args[1])
	fmt.Printf("2: %s\n",os.Args[2])


	// ipList := os.Args[1]
	// portList := os.Args[2]
	// fmt.Println("ipList:",ipList)
	// fmt.Println("portList:",portList)

	// if len(os.Args) == 3 {
	// 	ipList := os.Args[1]
	// 	portList := os.Args[2]
	// 	fmt.Println("ipList:",ipList)
	// 	fmt.Println("portList:",portList)
		//ips, err := util.GetIpList(ipList)
		//ports, err := util.GetPorts(portList)
		//_ = err
		// fmt.Printf("iplist: %v, portList: %v, err: %v\n", ips, ports, err)
		// for _, ip := range ips {
		// 	for _, port := range ports {
		// 		_, err := scanner.Connect(ip.String(), port)
		// 		if err != nil {
		// 			continue
		// 		}
		// 		fmt.Printf("ip: %v, port: %v is open\n", ip, port)
		// 	}
		// }

	// } else {
	// 	fmt.Printf("%v iplist port\n", os.Args[0])
	// }
}