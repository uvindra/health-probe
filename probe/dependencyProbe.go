package probe

import (
	"sync/atomic"
)

type DependencyProbe struct {
	clientName string
	serverName string
	*BaseProbe
}

func NewDependencyProbe(clientName, serverName string) *DependencyProbe {
	BaseProbe := &BaseProbe{errorCount: atomic.Uint32{}, successCount: atomic.Uint32{}}
	return &DependencyProbe{clientName: clientName, serverName: serverName, BaseProbe: BaseProbe}
}

func (p *DependencyProbe) GetClientName() string {
	return p.clientName
}

func (p *DependencyProbe) GetServerName() string {
	return p.serverName
}
