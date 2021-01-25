package iputil

import (
	"net"
	"testing"
)

func TestGetipList(t *testing.T) {
	wantData := []net.IP{
		net.ParseIP("114.114.114.114"),
		net.ParseIP("192.168.1.1")}
	testData, err := GetIpList("192.168.1.1, 114.114.114.114")

	if err != nil {
		t.Errorf("TestGetipList: faild resolve get error: %v", err)
	}

	for index := 0; index < len(wantData); index++ {
		if !wantData[index].Equal(testData[index]) {
			t.Errorf("TestGetipList: want Data does not equal parse Data")
		}
	}
}

func TestGetport(t *testing.T) {
	ports,err := GetPorts("22,23,24,80-81")
	if err != nil {
		t.Errorf("TestGetport: can not pasre ports:%v",err)
	}
	wantDatas := []int{22,23,24,80,81}
	for index := 0; index < len(wantDatas); index++ {
		if ports[index] != wantDatas[index] {
			t.Errorf("TestGetport: want data does not equal to parse data")
		}
	}
}
