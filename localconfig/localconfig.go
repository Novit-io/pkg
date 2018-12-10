package localconfig

import "strings"

type LocalConfig struct {
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

func (c *LocalConfig) ClusterByName(name string) *Cluster {
	for _, cluster := range c.Clusters {
		if cluster.Name == name {
			return cluster
		}
	}
	return nil
}

func (c *LocalConfig) HostByIP(ip string) *Host {
	for _, host := range c.Hosts {
		for _, hostIP := range host.IPs {
			if hostIP == ip {
				return host
			}
		}
	}
	return nil
}

func (c *LocalConfig) HostByMAC(mac string) *Host {
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
