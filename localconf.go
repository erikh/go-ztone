package one

import (
	"encoding/json"
	"io"
)

// LocalPhysConfig is the local.conf physical network configuration
type LocalPhysConfig struct {
	Blacklist     bool `json:"blacklist"`
	TrustedPathID int  `json:"trustedPathId,omitempty"`
	MTU           int  `json:"mtu,omitempty"`
}

// LocalVirtConfig is the local.conf virtual network configuration
type LocalVirtConfig struct {
	Try       []string `json:"try"`
	Blacklist []string `json:"blacklist"`
}

// LocalSettings is the bag of local.conf settings.
type LocalSettings struct {
	PrimaryPort        uint16 `json:"primaryPort,omitempty"`
	SecondaryPort      uint16 `json:"secondaryPort,omitempty"`
	TertiaryPort       uint16 `json:"tertiaryPort,omitempty"`
	PortMappingEnabled bool   `json:"portMappingEnabled,omitempty"`
	AllowSecondaryPort bool   `json:"allowSecondaryPort,omitempty"`
	// Enum of `apply`, `download`, or `disable`
	SoftwareUpdate string `json:"softwareUpdate,omitempty"`
	// Enum of `release`, or `beta`
	SoftwareUpdateChannel    string   `json:"softwareUpdateChannel,omitempty"`
	SoftwareUpdateDist       bool     `json:"softwareUpdateDist,omitempty"`
	InterfacePrefixBlacklist []string `json:"interfacePrefixBlacklist,omitempty"`
	AllowManagementFrom      []string `json:"allowManagementFrom,omitempty"`
	Bind                     []string `json:"bind,omitempty"`
	AllowTCPFallbackRelay    bool     `json:"allowTcpFallbackRelay,omitempty"`
	MultipathMode            uint     `json:"multipathMode,omitempty"`
}

// LocalConfiguration is the entire local.conf. Note that if you need to create
// an empty one, sometimes an empty struct is better than a fancy method. ;)
type LocalConfiguration struct {
	Physical map[string]LocalPhysConfig `json:"physical"`
	Virtual  map[string]LocalVirtConfig `json:"virtual"`
	Settings LocalSettings              `json:"settings"`
}

// NewLocalConfiguration constructs a new *LocalConfiguration from the incoming
// io.Reader.
func NewLocalConfiguration(r io.Reader) (*LocalConfiguration, error) {
	lc := &LocalConfiguration{}
	return lc, json.NewDecoder(r).Decode(lc)
}

// Write writes out the entire configuration to the io.Writer
func (lc *LocalConfiguration) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(lc)
}
