package tupletype

type TupleType int

const (
	MC TupleType = iota
	MP
	DONP
	DO
)

var names = [...]string{
	"MC",
	"MP",
	"DONP",
	"DO",
}

var fullnames = [...]string{
	"Mantencion Correctiva",
	"Mantencion Preventiva",
	"Detencion Operacional No Programada",
	"Detencion Operacional",
}

func (tt TupleType) Name() string {
	return names[tt]
}

func (tt TupleType) Fullname() string {
	return fullnames[tt]
}

func (tt TupleType) String() string {
	return names[tt]
}
