nmcli d
# lists devices with conns
nmcli con down id <connection>

nmcli con up id <connection>

https://developer.gnome.org/NetworkManager/unstable/gdbus-org.freedesktop.NetworkManager.html


nmcli c add type wifi ifname wifi-device con-name connection-name autoconnect no ssid hotspot-ssid
nmcli connection modify connection-name 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared
nmcli connection modify connection-name wifi-sec.key-mgmt wpa-psk
nmcli connection modify connection-name wifi-sec.psk "le password"
nmcli connection up connection-name

https://medium.com/@lane.c.wagner/bitbanging-spi-in-go-an-explanation-315d9c05737

https://github.com/davecheney/gpio

https://github.com/rtr7/router7
