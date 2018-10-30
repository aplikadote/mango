package tuple

import (
	"time"

	"github.com/vjeantet/jodaTime"
)

type Tuple struct {
	from time.Time
	to   time.Time
	tt   Type
}

func Init(from, to time.Time, tt Type) *Tuple {
	tuple := &Tuple{}
	tuple.from = from
	tuple.to = to
	tuple.tt = tt
	return tuple
}

func (t *Tuple) String() string {
	fromStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", t.from)
	toStr := jodaTime.Format("dd/MM/YYYY HH:mm:ss", t.to)
	return "[" + fromStr + " - " + toStr + "; " + t.tt.Name() + "]"
}
