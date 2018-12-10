package localconfig

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Clusters []*Cluster
	Hosts    []*Host
}

type Cluster struct {
	Name   string
	Addons []byte
}

type Host struct {
	Name string
	MACs []string
	IPs  []string

	Kernel string
	Initrd string
	Layers map[string]string

	Config []byte
}

func FromBytes(data []byte) (*Config, error) {
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

func FromFile(path string) (*Config, error) {
	ba, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return FromBytes(ba)
}

func (c *Config) ClusterByName(name string) *Cluster {
	for _, cluster := range c.Clusters {
		if cluster.Name == name {
			return cluster
		}
	}
	return nil
}

func (c *Config) HostByIP(ip string) *Host {
	for _, host := range c.Hosts {
		for _, hostIP := range host.IPs {
			if hostIP == ip {
				return host
			}
		}
	}
	return nil
}

func (c *Config) HostByMAC(mac string) *Host {
	// a bit of normalization
	mac = strings.Replace(strings.ToLower(mac), "-", ":", -1)

	for _, host := range c.Hosts {
		for _, hostMAC := range host.MACs {
			if strings.ToLower(hostMAC) == mac {
				return host
			}
		}
	}

	return nil
}
