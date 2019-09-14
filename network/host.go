package network

import (
	"net"

	gonm "github.com/BellerophonMobile/gonetworkmanager"
)

func WifiDevices() ([]WifiDevice, error) {
	nm, err := gonm.NewNetworkManager()
	if err != nil {
		return nil, err
	}

	devs := make([]WifiDevice, 0)
	devices := nm.GetDevices()

	for _, dev := range devices {
		if dev.GetDeviceType() == gonm.NmDeviceTypeWifi {
			rips := dev.GetIP4Config().GetAddresses()
			ips := make([]net.IP, len(rips))

			for _, rip := range rips {
				ip := net.ParseIP(rip.Address)
				if ip.String() != "" {
					ips = append(ips, ip)
				}
			}

			devs = append(devs, WifiDevice{
				Device{
					Name: dev.GetInterface(),
					Type: Wifi,
					// Status: ,
					IPv4Addresses: ips,
					rawDBusPath:   string(dev.GetPath()),
				},
			})
		}
	}

	return devs, nil
}

func EthernetDevices() ([]EthernetDevice, error) {
	nm, err := gonm.NewNetworkManager()
	if err != nil {
		return nil, err
	}

	devs := make([]EthernetDevice, 0)
	devices := nm.GetDevices()

	for _, dev := range devices {
		if dev.GetDeviceType() == gonm.NmDeviceTypeEthernet {
			rips := dev.GetIP4Config().GetAddresses()
			ips := make([]net.IP, len(rips))

			for _, rip := range rips {
				ip := net.ParseIP(rip.Address)
				if ip.String() != "" {
					ips = append(ips, ip)
				}
			}

			devs = append(devs, EthernetDevice{
				Device{
					Name: dev.GetInterface(),
					Type: Wifi,
					// Status: ,
					IPv4Addresses: ips,
					rawDBusPath:   string(dev.GetPath()),
				},
			})
		}
	}

	return devs, nil
}
