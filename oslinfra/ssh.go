package oslinfra

type Host struct {
	Name string
	Port int
}

type Inventory interface {
	Refresh() error
}

func SSHRefresh() string {
	return "helloworld"
}
