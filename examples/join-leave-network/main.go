package main

import (
	"os"

	one "github.com/erikh/go-ztone"
)

func errExit() {
	panic("supply a command (leave or join) and a network ID.")
}

func main() {
	if len(os.Args) != 3 {
		errExit()
	}

	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))

	switch os.Args[1] {
	case "leave":
		if err := c.LeaveNetwork(os.Args[2]); err != nil {
			panic(err)
		}
	case "join":
		if err := c.JoinNetwork(os.Args[2]); err != nil {
			panic(err)
		}
	default:
		errExit()
	}
}
