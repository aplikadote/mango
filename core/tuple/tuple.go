package tuple

import (
	"time"

	"github.com/aplikadote/mango/core/tupletype"
	"github.com/vjeantet/jodaTime"
)

type Tuple struct {
	From time.Time
	To   time.Time
	Tt   tupletype.TupleType
}

func (t *Tuple) String() string {
	fromStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", t.From)
	toStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", t.To)
	return "[" + fromStr + " - " + toStr + "]"
}
