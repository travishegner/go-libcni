package cni

import (
	"encoding/json"
)

// Config represents the cni network config file
type Config struct {
	CNIVersion     string      `json:"cniVersion"`
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	PreviousResult *Result     `json:"prevResult,omitempty"`
	Ipam           *IpamConfig `json:"ipam"`
	Args           *Args       `json:"args"`
}

// IpamConfig represents the cni ipam config
type IpamConfig struct {
	Type string `json:"type"`
}

//NewConfig returns a new standard cni config object
func NewConfig(confBytes []byte) (*Config, error) {
	conf := &Config{}
	err := json.Unmarshal(confBytes, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// Args are additional arguments provided by the container runtime
type Args struct {
	Annotations map[string]string `json:"annotations"`
}
