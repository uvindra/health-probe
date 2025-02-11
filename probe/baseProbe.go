package probe

import "sync/atomic"

type BaseProbe struct {
	errorCount   atomic.Uint32
	successCount atomic.Uint32
}

func (p *BaseProbe) IncrementErrorCount() {
	p.errorCount.Add(1)
}

func (p *BaseProbe) IncrementSuccessCount() {
	p.successCount.Add(1)
}

func (p *BaseProbe) GetErrorCount() uint32 {
	return p.errorCount.Load()
}

func (p *BaseProbe) GetSuccessCount() uint32 {
	return p.successCount.Load()
}

func (p *BaseProbe) Reset() {
	p.errorCount.Store(0)
	p.successCount.Store(0)
}
