package main

import (
	"encoding/json"
	"fmt"
	"os"

	gonm "github.com/BellerophonMobile/gonetworkmanager"
	"github.com/rs/zerolog"

	"gitlab.com/newrx/displayless"
)

type NetworkManager struct {
	client         gonm.NetworkManager
	wifiDevice     gonm.WirelessDevice
	ethernetDevice gonm.Device
}

func (nm *NetworkManager) ConnectWiFi(ssid string) error {
	return nil
}

// nmcli c add type wifi ifname wifi-device con-name connection-name autoconnect no ssid hotspot-ssid
// nmcli connection modify connection-name 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared
// nmcli connection modify connection-name wifi-sec.key-mgmt wpa-psk
// nmcli connection modify connection-name wifi-sec.psk "le password"
// nmcli connection up connection-name

func (nm *NetworkManager) CreateAP(ssid, password string) error {
	connection := make(map[string]map[string]interface{})
	connection["802-11-wireless"] = make(map[string]interface{})
	connection["802-11-wireless"]["security"] = "802-11-wireless-security"
	connection["802-11-wireless-security"] = make(map[string]interface{})
	connection["802-11-wireless-security"]["key-mgmt"] = "wpa-psk"
	connection["802-11-wireless-security"]["psk"] = password

	_, err := nm.client.AddAndActivateWirelessConnection(
		connection,
		nm.wifiDevice,
		nil, // accessPoint,
	)

	return err
}

func main() {
	logger := zerolog.New(os.Stdout)

	nm, err := gonm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

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

	e := nm.Subscribe()
	defer nm.Unsubscribe()

	logger.Info().Msg("starting server...")
	server := displayless.Start("0.0.0.0:80", true, true, logger)
	defer server.Close()

	for {
		event := <-e
		json.NewEncoder(os.Stdout).Encode(event)
	}
}
