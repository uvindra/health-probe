package probe

type LocalProbe struct {
	name string
	*BaseProbe
}

func NewLocalProbe(name string) *LocalProbe {
	return &LocalProbe{name: name}
}

func (p *LocalProbe) GetName() string {
	return p.name
}
