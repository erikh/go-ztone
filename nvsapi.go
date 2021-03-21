package one

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

type Route struct {
	Target string `json:"target"`
	Via    string `json:"via"`
	Flags  int64  `json:"flags"`
	Metric int64  `json:"metric"`
}

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

type Path struct {
	Address       string `json:"address"`
	LastSend      int64  `json:"lastSend"`
	LastReceive   int64  `json:"lastReceive"`
	Active        bool   `json:"active"`
	Expired       bool   `json:"expired"`
	Preferred     bool   `json:"preferred"`
	TrustedPathID int64  `json:"trustedPathId"`
}

func (c *Client) Status() (*NetworkStatus, error) {
	ns := &NetworkStatus{}
	return ns, c.wrapJSON("/status", ns)
}

func (c *Client) Networks() ([]*Network, error) {
	nws := []*Network{}
	return nws, c.wrapJSON("/network", &nws)
}

func (c *Client) Network(id string) (*Network, error) {
	nw := &Network{}
	return nw, c.wrapJSON("/network/"+id, nw)
}

func (c *Client) Peers() ([]*Peer, error) {
	peers := []*Peer{}
	return peers, c.wrapJSON("/peer", &peers)
}
