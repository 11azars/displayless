package network

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
