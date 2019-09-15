package network

import (
	"errors"

	gonm "github.com/BellerophonMobile/gonetworkmanager"
	"github.com/godbus/dbus"
)

type WifiNetworkMode string

const (
	AdHoc WifiNetworkMode = "adhoc"
	Infra WifiNetworkMode = "infra"
	AP    WifiNetworkMode = "ap"
)

type WifiNetwork struct {
	SSID      string
	Frequency uint32
	BSSID     string
	Mode      WifiNetworkMode

	// Bitrate in kilobits
	BitrateKbs     uint32
	SignalStrength float32

	APFlags  uint32
	WPAFlags uint32
	RSNFlags uint32

	rawDBusPath string
}

type WifiDevice struct {
	Device
}

func (w WifiDevice) Scan() ([]WifiNetwork, error) {
	dev, err := gonm.NewWirelessDevice(dbus.ObjectPath(w.rawDBusPath))
	if err != nil {
		return nil, err
	}

	networks := make([]WifiNetwork, 0)
	aps := dev.GetAccessPoints()

	for _, ap := range aps {
		rawmode := ap.GetMode()
		var mode WifiNetworkMode

		switch rawmode {
		case gonm.Nm80211ModeAdhoc:
			mode = AdHoc
			break
		case gonm.Nm80211ModeInfra:
			mode = Infra
			break
		case gonm.Nm80211ModeAp:
			mode = AP
			break
		default:
			return nil, errors.New("failed to determine network device mode")
		}

		networks = append(networks, WifiNetwork{
			SSID:      ap.GetSSID(),
			Frequency: ap.GetFrequency(),
			BSSID:     ap.GetHWAddress(),
			Mode:      mode,

			BitrateKbs:     ap.GetMaxBitrate(),
			SignalStrength: float32(ap.GetStrength()),

			// Flags

			rawDBusPath: string(ap.GetPath()),
		})
	}

	return networks, nil
}

// nmcli c add type wifi ifname wifi-device con-name connection-name autoconnect no ssid hotspot-ssid
// nmcli connection modify connection-name 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared
// nmcli connection modify connection-name wifi-sec.key-mgmt wpa-psk
// nmcli connection modify connection-name wifi-sec.psk "le password"
// nmcli connection up connection-name

// connection := make(map[string]map[string]interface{})
// 	connection["802-11-wireless"] = make(map[string]interface{})
// 	connection["802-11-wireless"]["security"] = "802-11-wireless-security"
// 	connection["802-11-wireless-security"] = make(map[string]interface{})
// 	connection["802-11-wireless-security"]["key-mgmt"] = "wpa-psk"
// 	connection["802-11-wireless-security"]["psk"] = password

// 	_, err := nm.client.AddAndActivateWirelessConnection(
// 		connection,
// 		nm.wifiDevice,
// 		nil, // accessPoint,
// 	)

// AddAndActivateWirelessConnection(connection map[string]map[string]interface{}, device Device, accessPoint AccessPoint) (ac ActiveConnection, err error)

func (w WifiDevice) ConnectTo(network WifiNetwork) error {
	return nil
}

func (w WifiDevice) ConnectToNetwork(ssid, password string) error {
	return nil
}

func (w WifiDevice) ExposeAP(ssid, password string) error {
	return nil
}
