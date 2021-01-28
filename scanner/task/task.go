package task

import (
	"fmt"
	"myTools/scanner/tcpConnectScanner/tcpScannerTool"
	"myTools/scanner/vars"
	"net"
	"strings"
	"sync"
)

type TaskInterface interface {
	GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int)
	RunTask(tasks []map[string]int)
	Scan(taskChan chan map[string]int, wg *sync.WaitGroup)
	SaveResult(ip string, port int, err error) error
	PrintResult()
}

type DefaultTask struct {
	TcpTool tcpScannerTool.MyTcpTool
}

func (task* DefaultTask) GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	tasks := make([]map[string]int, 0)

	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String(): port}
			tasks = append(tasks, ipPort)
		}
	}
	return tasks, len(tasks)
}

func (task* DefaultTask) RunTask(tasks []map[string]int) {
	wg := &sync.WaitGroup{}
	taskChan := make(chan map[string]int, vars.ThreadNum*2)

	for index := 0; index < vars.ThreadNum; index++ {
		go task.Scan(taskChan, wg)
	}
	for _, taskItem := range tasks {
		wg.Add(1)
		taskChan <- taskItem
	}
	close(taskChan)
	wg.Wait()
}

func (task* DefaultTask) Scan(taskChan chan map[string]int,
	wg *sync.WaitGroup) {

	for taskItem := range taskChan {
		for ip, port := range taskItem {
			err := task.SaveResult(task.TcpTool.MyConnect(ip, port))
			_ = err
			wg.Done()
		}
	}
}

func (task* DefaultTask) SaveResult(ip string, port int, err error) error {
	if err != nil {
		return err
	}
	v, ok := vars.Result.Load(ip)
	if ok {
		ports, ok1 := v.([]int)
		if ok1 {
			ports = append(ports, port)
			vars.Result.Store(ip, ports)
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		vars.Result.Store(ip, ports)
	}
	return err
}

func (task* DefaultTask) PrintResult() {
	vars.Result.Range(func(key, value interface{}) bool {
		fmt.Printf("ip:%v\n", key)
		fmt.Printf("ports: %v\n", value)
		fmt.Println(strings.Repeat("-", 100))
		return true
	})
}
