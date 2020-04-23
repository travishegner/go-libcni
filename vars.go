package cni

import (
	"os"
	"strings"
)

//Vars represents the environment variables that the runtime should have populated
type Vars struct {
	Command            string
	NetworkNamespace   string
	ContainerInterface string
	ContainerID        string
	Arguments          string
	Path               string
	args               map[string]string
	parsed             bool
}

// NewVars returns an object populated with the standard CNI environment variables
func NewVars() *Vars {
	v := &Vars{
		Command:            os.Getenv("CNI_COMMAND"),
		NetworkNamespace:   os.Getenv("CNI_NETNS"),
		ContainerInterface: os.Getenv("CNI_IFNAME"),
		ContainerID:        os.Getenv("CNI_CONTAINERID"),
		Arguments:          os.Getenv("CNI_ARGS"),
		Path:               os.Getenv("CNI_PATH"),
	}

	return v
}

func (v *Vars) parseArgs() {
	if v.parsed {
		return
	}

	v.args = make(map[string]string)

	kvs := strings.Split(v.Arguments, ";")
	for _, kv := range kvs {
		a := strings.Split(kv, "=")
		if len(a) > 1 {
			v.args[a[0]] = a[1]
		}
	}

	v.parsed = true
}

//GetArg gets a specific argument passed in by the CNI_ARGS variable
func (v *Vars) GetArg(key string) (string, bool) {
	if !v.parsed {
		v.parseArgs()
	}

	value, ok := v.args[key]
	return value, ok
}
