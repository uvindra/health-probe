package probe

import "sync/atomic"

type LocalProbe struct {
	name string
	*BaseProbe
}

func NewLocalProbe(name string) *LocalProbe {
	BaseProbe := &BaseProbe{errorCount: atomic.Uint32{}, successCount: atomic.Uint32{}}
	return &LocalProbe{name: name, BaseProbe: BaseProbe}
}

func (p *LocalProbe) GetName() string {
	return p.name
}
