package unittest

import (
	"fmt"

	"github.com/malfunkt/iprange"

)

func GetipList() string {
	list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	//fmt.Printf("%+v", list)

	rng := list.Expand()
	fmt.Printf("%s",rng)
	return "123"
}
