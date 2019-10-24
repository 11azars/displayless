package main

import (
	"fmt"

	"github.com/coreos/go-iptables/iptables"
)

func main() {
	tables, err := iptables.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(tables.Proto())
	fmt.Println(tables.ListChains("OUTPUT"))
}
