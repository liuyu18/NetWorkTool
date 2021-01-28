package main

import (
	"fmt"
	"myTools/scanner/iputil"
	"myTools/scanner/task"
	"myTools/scanner/tcpConnectScanner/tcpScannerTool"
	"os"
	"runtime"
)

func main() {

	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]

		iputil := new(iputil.DefaultIputil)
		task := &task.DefaultTask{
			TcpTool: tcpScannerTool.DefaultTcpTool{},
		}

		ips, err := iputil.GetIpList(ipList)
		if err != nil {
			fmt.Printf("parse ip list error: %+v\n", err)
		}
		ports, err := iputil.GetPorts(portList)

		if err != nil {
			fmt.Printf("parse ports error: %+v\n", err)
		}

		tasks, _ := task.GenerateTask(ips, ports)
		task.RunTask(tasks)
		task.PrintResult()
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}

// go run .\main.go 45.22.2.156,114.114.114.114 22,23,24,25,80-100
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
