package net

import (
	"testing"
	"fmt"
)

func TestTcpGather(t *testing.T) {


	ports := []string{
		"80",
		"443",
		"8080",
	}

	portsStatus, err := TcpGather("www.baidu.com", ports)
	fmt.Println(portsStatus)
	
	if err != nil {
		t.Error(err.Error())
	}

	p443, ok := portsStatus["443"]
	if !(ok && p443) {
		t.Error("check port 443 failed")
	}

}