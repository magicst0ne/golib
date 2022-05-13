package net

import (
	"testing"
)

func TestTcpGather(t *testing.T) {


	ports := []string{
		"443",
	}

	portsStatus, err := TcpGather("www.baidu.com", ports)
	
	if err != nil {
		t.Error(err.Error())
	}

	p443, ok := portsStatus["443"]
	if !(ok && p443) {
		t.Error("check port 443 failed")
	}

}