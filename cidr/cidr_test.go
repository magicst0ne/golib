package cidr

import (
	"testing"
)

func TestAddressRange(t *testing.T) {


	hosts, err := AddressRange("172.25.104.0/24")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(hosts)!=254 {
		t.Error("ip range wrong")
	}
}