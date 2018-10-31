package tuple

import "errors"

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

func ParseTupleType(str string) (TupleType, error) {
	index := -1
	for i, val := range names {
		if str == val {
			index = i
			break
		}
	}

	if index != -1 {
		return TupleType(index), nil
	} else {
		err := errors.New("tipo no soportado")
		return -1, err
	}
}
