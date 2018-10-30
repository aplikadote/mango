package subsystem

import "errors"

type Type int

const (
	LINEAR Type = iota
	FRACTION
	PARALLEL
	REDUNDANT
	STANDY
	EQUIPMENT
	STOCKPILE
)

var names = [...]string{
	"LINEAR",
	"FRACTION",
	"PARALLEL",
	"REDUNDANT",
	"STANDY",
	"EQUIPMENT",
	"STOCKPILE",
}

var fullnames = [...]string{
	"LINEAR",
	"FRACTION",
	"PARALLEL",
	"REDUNDANT",
	"STANDY",
	"EQUIPMENT",
	"STOCKPILE",
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

func ParseType(str string) (Type, error) {
	index := -1
	for i, val := range names {
		if str == val {
			index = i
			break
		}
	}

	if index != -1 {
		return Type(index), nil
	} else {
		err := errors.New("tipo no soportado")
		return -1, err
	}
}
