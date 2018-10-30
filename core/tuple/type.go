package tuple

type Type int

const (
	MC Type = iota
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

func (tt Type) Name() string {
	return names[tt]
}

func (tt Type) Fullname() string {
	return fullnames[tt]
}

func (tt Type) String() string {
	return names[tt]
}
