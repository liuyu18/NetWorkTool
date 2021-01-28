package iputil

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/malfunkt/iprange"
)

type iputilInterface interface {
	GetPorts(selection string) ([]int, error)
	GetIpList(ips string)([]net.IP,error)
}

type DefaultIputil struct {}



func (util *DefaultIputil) GetPorts(selection string) ([]int, error) {
	ports := make([]int, 0)
	if selection == "" {
		return ports, nil
	}

	ranges := strings.Split(selection, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid port selection segment: '%s'", r)
			}

			p1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", parts[0])
			}

			p2, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", parts[1])
			}

			if p1 > p2 {
				return nil, fmt.Errorf("Invalid port range: %d-%d", p1, p2)
			}

			for i := p1; i <= p2; i++ {
				ports = append(ports, i)
			}

		} else {
			if port, err := strconv.Atoi(r); err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'", r)
			} else {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}

func (util *DefaultIputil) GetIpList(ips string)([]net.IP,error) {
	addressList, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	list := addressList.Expand()

	return list, err
}
