package subsystem

type Chain struct {
	rootSubsystem *Subsystem
}

func NewChain() *Chain {
	return &Chain{}
}

func (c *Chain) RootSubsystem() *Subsystem {
	return c.rootSubsystem
}

func (c *Chain) SetRootSubsystem(ss *Subsystem) {
	c.rootSubsystem = ss
}

func (c *Chain) String() string {
	return c.rootSubsystem.Name
}
