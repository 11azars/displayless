package main

import (
	"encoding/json"
	"os"

	"github.com/11azars/displayless/network"
)

func main() {
	devices, err := network.WifiDevices()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(devices)

	for _, dev := range devices {
		list, err := dev.Scan()
		if err != nil {
			panic(err)
		}

		json.NewEncoder(os.Stdout).Encode(list)
	}
}
