package one

// NetworkStatus is the API response to a /status call.
type NetworkStatus struct {
	Address           string `json:"address"`
	PublicIdentity    string `json:"publicIdentity"`
	WorldID           int64  `json:"worldId"`
	WorldTimestamp    int64  `json:"worldTimestamp"`
	Online            bool   `json:"online"`
	TCPFallbackActive bool   `json:"tcpFallbackActive"`
	// Enum: ALWAYS / TRUSTED / NEVER
	RelayPolicy  string `json:"relayPolicy"`
	VersionMajor int64  `json:"versionMajor"`
	VersionMinor int64  `json:"versionMinor"`
	VersionRev   int64  `json:"versionRev"`
	Version      string `json:"version"`
	Clock        int64  `json:"clock"`
}

// Network is the data structure that encapsulates a network.
type Network struct {
	ID                string   `json:"id"`
	LegacyNetworkID   string   `json:"nwid"`
	MAC               string   `json:"mac"`
	Name              string   `json:"name"`
	Status            string   `json:"status"`
	Type              string   `json:"type"`
	MTU               int      `json:"mtu"`
	DHCP              bool     `json:"dhcp"`
	Bridge            bool     `json:"bridge"`
	BroadcastEnabled  bool     `json:"broadcastEnabled"`
	PortError         int      `json:"portError"`
	NetconfRevision   int64    `json:"netconfRevision"`
	AssignedAddresses []string `json:"assignedAddresses"`
	Routes            []Route  `json:"routes"`
	PortDeviceName    string   `json:"portDeviceName"`
	AllowManaged      bool     `json:"allowManaged"`
	AllowGlobal       bool     `json:"allowGlobal"`
	AllowDefault      bool     `json:"allowDefault"`
	AllowDNS          bool     `json:"allowDNS"`
}

// Route encapsulates network routes. See Network.
type Route struct {
	Target string `json:"target"`
	Via    string `json:"via"`
	Flags  int64  `json:"flags"`
	Metric int64  `json:"metric"`
}

// Peer encapsulates ZeroTier One peers.
type Peer struct {
	Address      string `json:"address"`
	VersionMajor int64  `json:"versionMajor"`
	VersionMinor int64  `json:"versionMinor"`
	VersionRev   int64  `json:"versionRev"`
	Version      string `json:"version"`
	Latency      int64  `json:"latency"`
	// Enum: LEAF / UPSTREAM / ROOT / PLANET
	Role  string `json:"role"`
	Paths []Path `json:"paths"`
}

// Path is the path on the network.
type Path struct {
	Address       string `json:"address"`
	LastSend      int64  `json:"lastSend"`
	LastReceive   int64  `json:"lastReceive"`
	Active        bool   `json:"active"`
	Expired       bool   `json:"expired"`
	Preferred     bool   `json:"preferred"`
	TrustedPathID int64  `json:"trustedPathId"`
}

// Status returns the status of the ZeroTier One instance
func (c *Client) Status() (*NetworkStatus, error) {
	ns := &NetworkStatus{}
	return ns, c.wrapJSON("/status", ns)
}

// Networks returns all networks that ZeroTier One knows about.
func (c *Client) Networks() ([]*Network, error) {
	nws := []*Network{}
	return nws, c.wrapJSON("/network", &nws)
}

// Network queries a specific network.
func (c *Client) Network(id string) (*Network, error) {
	nw := &Network{}
	return nw, c.wrapJSON("/network/"+id, nw)
}

// Peers queries the peers that ZeroTier One knows about.
func (c *Client) Peers() ([]*Peer, error) {
	peers := []*Peer{}
	return peers, c.wrapJSON("/peer", &peers)
}
