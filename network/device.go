package network

import (
	"errors"
	"net"

	gonm "github.com/BellerophonMobile/gonetworkmanager"
	"github.com/godbus/dbus"
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

// GetPath() dbus.ObjectPath

//     // GetFlags gets flags describing the capabilities of the access point.
//     GetFlags() (uint32, error)

//     // GetWPAFlags gets flags describing the access point's capabilities
//     // according to WPA (Wifi Protected Access).
//     GetWPAFlags() (uint32, error)

//     // GetRSNFlags gets flags describing the access point's capabilities
//     // according to the RSN (Robust Secure Network) protocol.
//     GetRSNFlags() (uint32, error)

//     // GetSSID returns the Service Set Identifier identifying the access point.
//     GetSSID() (string, error)

//     // GetFrequency gets the radio channel frequency in use by the access point,
//     // in MHz.
//     GetFrequency() (uint32, error)

//     // GetHWAddress gets the hardware address (BSSID) of the access point.
//     GetHWAddress() (string, error)

//     // GetMode describes the operating mode of the access point.
//     GetMode() (Nm80211Mode, error)

//     // GetMaxBitrate gets the maximum bitrate this access point is capable of, in
//     // kilobits/second (Kb/s).
//     GetMaxBitrate() (uint32, error)

//     // GetStrength gets the current signal quality of the access point, in
//     // percent.
//     GetStrength() (uint8, error)
