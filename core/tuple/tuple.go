package tuple

import (
	"time"

	"github.com/vjeantet/jodaTime"
)

type Tuple struct {
	from      time.Time
	to        time.Time
	tupleType TupleType
}

func NewTuple(from, to time.Time, tupleType TupleType) *Tuple {
	tuple := &Tuple{}
	tuple.from = from
	tuple.to = to
	tuple.tupleType = tupleType
	return tuple
}

func (tuple *Tuple) String() string {
	fromStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", tuple.from)
	toStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", tuple.to)
	return "[" + fromStr + " - " + toStr + "; " + tuple.tupleType.Name() + "]"
}
