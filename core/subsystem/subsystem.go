package subsystem

import "github.com/aplikadote/mango/core/tuple"

type Subsystem struct {
	name      string
	nickname  string
	impact    float64
	minToWork int
	parent    *Subsystem
	children  []*Subsystem
	tuples    []*tuple.Tuple
}

func Init(name, nickname string) *Subsystem {
	ss := &Subsystem{}
	ss.name = name
	ss.nickname = nickname
	return ss
}

func (ss *Subsystem) Children() []*Subsystem {
	return ss.children
}

func (ss *Subsystem) AddChild(child *Subsystem) {
	ss.children = append(ss.children, child)
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
}

func (ss *Subsystem) Tuples() []*tuple.Tuple {
	return ss.tuples
}

func (ss *Subsystem) String() string {
	return ss.name
}
