package config

import (
	"github.com/mycoria/mycoria/m"
)

// Store holds all configuration in a storable format.
type Store struct {
	Router Router `json:"router,omitempty" yaml:"router,omitempty"`
	System System `json:"system,omitempty" yaml:"system,omitempty"`

	ServiceConfigs []ServiceConfig   `json:"services,omitempty" yaml:"services,omitempty"`
	FriendConfigs  []FriendConfig    `json:"friends,omitempty"  yaml:"friends,omitempty"`
	ResolveConfig  map[string]string `json:"resolve,omitempty"  yaml:"resolve,omitempty"`
}

// Router defines all configuration regarding the overlay network itself.
type Router struct { //nolint:maligned
	// Address it the identity of the router.
	Address m.AddressStorage `json:"address,omitempty" yaml:"address,omitempty"`

	// Universe holds the "universe" the router is in.
	Universe       string `json:"universe,omitempty"       yaml:"universe,omitempty"`
	UniverseSecret string `json:"universeSecret,omitempty" yaml:"universeSecret,omitempty"`

	// Isolate constrains outgoing traffic to friends.
	Isolate bool `json:"isolate,omitempty" yaml:"isolate,omitempty"`

	// Listen holds the peering URLs to listen on.
	// URLs must have an IP address as host.
	Listen []string `json:"listen,omitempty" yaml:"listen,omitempty"`

	// IANA holds a list of domains or IPs assigne by IANA through which the router can be reached.
	IANA []string `json:"iana,omitempty" yaml:"iana,omitempty"`

	// Connect holds the peering URLs the router
	// tries to always hold a connection to.
	Connect []string `json:"connect,omitempty" yaml:"connect,omitempty"`

	// AutoConnect specifies whether the router should automatically peer with
	// other routers (based on live usage data) to improve network flow.
	AutoConnect bool `json:"autoConnect,omitempty" yaml:"autoConnect,omitempty"`

	// Bootstrap holds peering URLs that the router uses to bootstrap to the network.
	Bootstrap []string `json:"bootstrap,omitempty" yaml:"bootstrap,omitempty"`
}

// FriendConfig is a trusted router in the network.
type FriendConfig struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	IP   string `json:"ip,omitempty"   yaml:"ip,omitempty"`
}

// ServiceConfig defines an endpoint other routers can send traffic to.
type ServiceConfig struct { //nolint:maligned
	Name        string `json:"name,omitempty"        yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Domain      string `json:"domain,omitempty"      yaml:"domain,omitempty"`
	URL         string `json:"url,omitempty"         yaml:"url,omitempty"`

	// Access Control
	Public  bool     `json:"public,omitempty"  yaml:"public,omitempty"`
	Friends bool     `json:"friends,omitempty" yaml:"friends,omitempty"`
	For     []string `json:"for,omitempty"     yaml:"for,omitempty"`

	Advertise bool `json:"advertise,omitempty" yaml:"advertise,omitempty"`
}

// System defines all configuration regarding the system.
type System struct {
	TunName   string `json:"tunName,omitempty"   yaml:"tunName,omitempty"`
	StatePath string `json:"statePath,omitempty" yaml:"statePath,omitempty"`
}