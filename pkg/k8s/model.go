package k8s

type Cluster struct {
	BomFormat    string       `json:"bomFormat,omitempty"`
	SpecVersion  string       `json:"specVersion,omitempty"`
	SerialNumber string       `json:"serialNumber,omitempty"`
	Metadata     Metadata     `json:"metadata,omitempty"`
	Version      int          `json:"version,omitempty"`
	ControlPlane ControlPlane `json:"control_plane,omitempty"`
	NodesInfo    []NodeInfo   `json:"nodes,omitempty"`
	Addons       []*Component `json:"addons,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type NodeInfo struct {
	NodeRole                string `json:"node_role,omitempty"`
	NodeName                string `json:"node_name,omitempty"`
	KubeletVersion          string `json:"kubelet_version,omitempty"`
	ContainerRuntimeVersion string `json:"container_runtime_version,omitempty"`
	OsImage                 string `json:"os_image,omitempty"`
	Hostname                string `json:"host_name,omitempty"`
	KernelVersion           string `json:"kernel_version,omitempty"`
	KubeProxyVersion        string `json:"kube_proxy_version,omitempty"`
	OperatingSystem         string `json:"operating_system,omitempty"`
	Architecture            string `json:"architecture,omitempty"`
}

type ControlPlane struct {
	Components []*Component `json:"components,omitempty"`
}

type Component struct {
	BomRef     string     `json:"bom-ref,omitempty"`
	Type       string     `json:"type,omitempty"`
	Name       string     `json:"name,omitempty"`
	Purl       string     `json:"purl,omitempty"`
	Version    string     `json:"version,omitempty"`
	Licenses   []License  `json:"licenses,omitempty"`
	Properties []Property `json:"properties,omitempty"`
}

type License struct {
	Expression string `json:"expression,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Metadata struct {
	Timestamp string    `json:"timestamp,omitempty"`
	Component Component `json:"component,omitempty"`
	Tools     []Tool    `json:"tools,omitempty"`
}

type Tool struct {
	Vendor  string `json:"vendor,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Dependency struct {
	Ref       string   `json:"ref,omitempty"`
	DependsOn []string `json:"dependsOn,omitempty"`
}
