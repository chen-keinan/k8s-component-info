package k8s

type BomResult struct {
	Target   string    `json:"Target"`
	Class    string    `json:"Class,omitempty"`
	Type     string    `json:"Type,omitempty"`
	Packages []Package `json:"Packages,omitempty"`
}

type Package struct {
	ID         string     `json:",omitempty"`
	Name       string     `json:",omitempty"`
	Version    string     `json:",omitempty"`
	Properties []KeyValue `json:",omitempty"`
	DependsOn  []string   `json:",omitempty"`
	Digest     string     `json:",omitempty"`
}

type KeyValue struct {
	Name  string
	Value string
}

type TargetMetadata struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
