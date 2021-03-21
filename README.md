# go-ztone: a client to the ZeroTier One local configuration socket

`go-ztone` is for using against your local [ZeroTier One](https://www.zerotier.com) instance. You must use the "secret authtoken" described in [this document](https://github.com/zerotier/zerotierone#running). For the examples specifically, the convention is to use `ZEROTIER_ONE_TOKEN` on the commandline to assign it into the environment.

```go
package main

import (
	"fmt"
	"os"

	one "github.com/erikh/go-ztone"
)

func main() {
	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))
	networks, err := c.Networks()
	if err != nil {
		panic(err)
	}

	peers, err := c.Peers()
	if err != nil {
		panic(err)
	}

	fmt.Println("Networks w/ MAC:")
	for _, network := range networks {
		fmt.Println(network.ID, network.MAC)
	}

	fmt.Println("Peers w/ Latency:")
	for _, peer := range peers {
		fmt.Println(peer.Address, peer.Latency)
	}
}
```

## Functionality

`ztone` has a few basic functions, most of them [listed here](https://github.com/zerotier/ZeroTierOne/blob/master/service/README.md#network-virtualization-service-api), as well as on the [GoDoc](https://pkg.go.dev/github.com/erikh/go-ztone).

## Example Code

You can see examples in the `examples/` directory. There are two:

- `list-things`: lists a few different properties and takes no arguments.
- `query-things`: takes a network ID and returns a few properties about it.

# Author

Erik Hollensbe <github@hollensbe.org>

# License

BSD 3-Clause
