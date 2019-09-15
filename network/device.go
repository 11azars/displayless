package network

import (
	"net"
)

type InterfaceType string
type DeviceState string

const (
	Wifi     InterfaceType = "wifi"
	Ethernet InterfaceType = "ethernet"
)

type Device struct {
	Name          string
	Type          InterfaceType
	State         InterfaceType
	IPv4Addresses []net.IP

	rawDBusPath string
}

type EthernetDevice struct {
	Device
}
