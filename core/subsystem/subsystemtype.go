package subsystem

import "errors"

type SubsystemType int

const (
	LINEAR SubsystemType = iota
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

func (tt *SubsystemType) Name() string {
	return names[*tt]
}

func (tt *SubsystemType) Fullname() string {
	return fullnames[*tt]
}

func (tt *SubsystemType) String() string {
	return names[*tt]
}

func ParseSubsystemType(str string) (SubsystemType, error) {
	index := -1
	for i, val := range names {
		if str == val {
			index = i
			break
		}
	}

	if index != -1 {
		return SubsystemType(index), nil
	} else {
		err := errors.New("tipo no soportado")
		return -1, err
	}
}
