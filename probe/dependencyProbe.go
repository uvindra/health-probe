package probe

type DependencyProbe struct {
	clientName string
	serverName string
	*BaseProbe
}

func NewDependencyProbe(clientName, serverName string) *DependencyProbe {
	return &DependencyProbe{clientName: clientName, serverName: serverName}
}

func (p *DependencyProbe) GetClientName() string {
	return p.clientName
}

func (p *DependencyProbe) GetServerName() string {
	return p.serverName
}
