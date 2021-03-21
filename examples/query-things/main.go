package main

import (
	"encoding/json"
	"fmt"
	"os"

	one "github.com/erikh/go-ztone"
)

func main() {
	if len(os.Args) != 2 {
		panic("supply a zerotier network id")
	}

	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))
	network, err := c.Network(os.Args[1])
	if err != nil {
		panic(err)
	}

	content, err := json.MarshalIndent(network, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
