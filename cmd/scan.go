// GetPath() dbus.ObjectPath

// // GetInterface gets the name of the device's control (and often data)
// // interface.
// GetInterface() (string, error)

// // GetIpInterface gets the IP interface name of the device.
// GetIpInterface() (string, error)

// // GetState gets the current state of the device.
// GetState() (NmDeviceState, error)

// // GetIP4Config gets the Ip4Config object describing the configuration of the
// // device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED
// // state.
// GetIP4Config() (IP4Config, error)

// // GetDHCP4Config gets the Dhcp4Config object describing the configuration of the
// // device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED
// // state.
// GetDHCP4Config() (DHCP4Config, error)

// // GetDeviceType gets the general type of the network device; ie Ethernet,
// // WiFi, etc.
// GetDeviceType() (NmDeviceType, error)

// // GetAvailableConnections gets an array of object paths of every configured
// // connection that is currently 'available' through this device.
// GetAvailableConnections() ([]Connection, error)

package main

import (
	"fmt"
	"os"
	"time"

	gonm "github.com/BellerophonMobile/gonetworkmanager"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Info().Msg("starting server...")

	nm, err := gonm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	for {
		devices := nm.GetDevices()
		for _, device := range devices {
			ifname := device.GetInterface()
			state := device.GetState()
			deviceType := device.GetDeviceType()
			path := device.GetPath()

			fmt.Println(ifname, deviceType, state, path)
			if deviceType == gonm.NmDeviceTypeWifi {
				wifiDev, err := gonm.NewWirelessDevice(path)
				if err != nil {
					panic(err)
				}

				aps := wifiDev.GetAccessPoints()
				for _, ap := range aps {
					fmt.Println(ap.GetSSID())
				}
			}
		}

		<-time.NewTimer(time.Second * 5).C
	}
}
