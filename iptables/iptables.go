package iptables

type IPTableCmd []string
type Profile []IPTableCmd

var (
  // P2P allows the following connections
  // incoming 5000 UDP
  // outgoing 5000 UDP
  P2P Profile = []IPTableCmd{}

  // Consumer allows the following connections
  // outgoing 80, 443 TCP, UDP
  Consumer Profile = []IPTableCmd{}

  // WiFiAP allows the following connections
  // incoming 80, 443 TCP
  // outgoing 80, 443 TCP
  // preroute from any to localhost:8000
  // postroute back to original destination
  WiFiAP Profile = []IPTableCmd{}
)

func Activate(profiles ...Profile) error {
  return nil
}

func Deactivate(profiles ...Profile) error {
  return nil
}
