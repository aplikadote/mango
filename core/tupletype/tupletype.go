package tupletype

type TupleType int

const (
	MC TupleType = iota
	MP
	DONP
	DO
)

var Names = [...]string{
	"MC",
	"MP",
	"DONP",
	"DO",
}
