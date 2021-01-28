package task

import (
	"myTools/scanner/tcpConnectScanner/tcpScannerTool"
	"net"
	"testing"
	// "github.com/golang/mock/gomock"
)

func TestGenerateTask(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	task := DefaultTask{
		tcpTool: tcpScannerTool.DefaultTcpTool{},
	}
	want := make([]map[string]int, 6)

	for index := range want {
		want[index] = make(map[string]int, 1)
	}

	want[0]["114.114.114.114"] = 22
	want[1]["114.114.114.114"] = 23
	want[2]["114.114.114.114"] = 24
	want[3]["192.168.1.1"] = 22
	want[4]["192.168.1.1"] = 23
	want[5]["192.168.1.1"] = 24

	wantDataKey := []string{
		"114.114.114.114",
		"114.114.114.114",
		"114.114.114.114",
		"192.168.1.1",
		"192.168.1.1",
		"192.168.1.1",
	}
	wantLenght := 6

	TestIpDatas := []net.IP{
		net.ParseIP("114.114.114.114"),
		net.ParseIP("192.168.1.1")}

	TestPortDatas := []int{22, 23, 24}

	resultData, len := task.GenerateTask(TestIpDatas, TestPortDatas)

	if len != wantLenght {
		t.Errorf("test GenerateTask Error")
	}
	for index := 0; index < cap(want); index++ {
		wantMap := want[index]
		resultDataMap := resultData[index]
		key := wantDataKey[index]
		if wantMap[key] != resultDataMap[key] {
			t.Errorf("test GenerateTask Error")
		}
	}
}
