package main

import (
	"encoding/json"
	"fmt"
	"os"

	one "github.com/erikh/go-ztone"
)

func errExit() {
	panic("supply a thing to query (network or peer) and a value: either the netID for networks, or the address for peers")
}

func main() {
	if len(os.Args) != 3 {
		errExit()
	}

	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))

	var (
		value interface{}
		err   error
	)

	switch os.Args[1] {
	case "network":
		value, err = c.GetNetwork(os.Args[2])
		if err != nil {
			panic(err)
		}
	case "peer":
		value, err = c.GetPeer(os.Args[2])
		if err != nil {
			panic(err)
		}
	default:
		errExit()
	}

	content, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
