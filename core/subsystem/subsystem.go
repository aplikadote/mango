package subsystem

import "github.com/aplikadote/mango/core/tuple"

type Subsystem struct {
	Name      string
	Nickname  string
	SsType    SubsystemType
	Impact    float64
	MinToWork int
	parent    *Subsystem
	children  []*Subsystem
	tuples    []*tuple.Tuple
}

func NewSubsystem(name, nick string) *Subsystem {
	return &Subsystem{Name: name, Nickname: nick}
}

func (ss *Subsystem) Children() []*Subsystem {
	return ss.children
}

func (ss *Subsystem) AddChild(child *Subsystem) {
	ss.children = append(ss.children, child)
	child.parent = ss
}

func (ss *Subsystem) DeleteChild(child *Subsystem) {
	var index int = 0
	for i, v := range ss.children {
		if v == child {
			index = i
			break
		}
	}
	ss.children = append(ss.children[:index], ss.children[index+1:]...)
	child.parent = nil
}

func (ss *Subsystem) Parent() *Subsystem {
	return ss.parent
}

func (ss *Subsystem) Tuples() []*tuple.Tuple {
	return ss.tuples
}

func (ss *Subsystem) PreOrden(visitor func(*Subsystem)) {
	visitor(ss)
	for _, child := range ss.children {
		child.PreOrden(visitor)
	}
}

func (ss *Subsystem) PostOrden(visitor func(*Subsystem)) {
	for _, child := range ss.children {
		child.PreOrden(visitor)
	}
	visitor(ss)
}

func (ss *Subsystem) String() string {
	return ss.Name
}
