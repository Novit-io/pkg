package localconfig

import (
	"io"

	yaml "gopkg.in/yaml.v2"
)

type Host struct {
	Name string
	MACs []string
	IPs  []string

	Kernel string
	Initrd string
	Layers map[string]string

	Config []byte
}

func (h *Host) WriteHashDataTo(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(Host{
		Kernel: h.Kernel,
		Initrd: h.Initrd,
		Layers: h.Layers,
		Config: h.Config,
	})
}
