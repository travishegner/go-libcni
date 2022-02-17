package cni

import (
	"encoding/json"
	"fmt"
)

//Result represents the cni result given back to the caller
type Result struct {
	CNIVersion string       `json:"cniVersion"`
	Interfaces []*Interface `json:"interfaces,omitempty"`
	IPs        []*IP        `json:"ips"`
	Routes     []*Route     `json:"route,omitempty"`
	DNS        *DNS         `json:"dns,omitempty"`
}

//Interface represents the interface in the Result
type Interface struct {
	Name    string `json:"name"`
	MAC     string `json:"mac,omitempty"`
	Sandbox string `json:"sandbox,omitempty"`
}

//IP represents the ip address in the Result
type IP struct {
	Version   string `json:"version"`
	Address   string `json:"address"`
	Gateway   string `json:"gateway,omitempty"`
	Interface *int   `json:"interface,omitempty"`
}

//Route represents any route in the Result
type Route struct {
	Destination string `json:"dst"`
	Gateway     string `json:"gw,omitempty"`
}

//DNS represents the DNS structure in the Result
type DNS struct {
	Nameservers []string `json:"nameservers,omitempty"`
	Domain      string   `json:"domain,omitempty"`
	Search      []string `json:"search,omitempty"`
	Options     []string `json:"options,omitempty"`
}

//Marshal marshals the result into a json byte array
func (r *Result) Marshal() []byte {
	ebytes, err := json.Marshal(r)
	if err != nil {
		return []byte(fmt.Sprintf("{\"cniVersion\": \"%v\", \"code\": 99, \"msg\":\"error marshaling result\", \"details\":\"there was an error marshaling the result to json\"}", CNIVersion))
	}

	return ebytes
}
