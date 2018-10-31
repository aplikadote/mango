package subsystem

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

type CsvLine struct {
	id       int
	idParent int
	ss       *Subsystem
}

func ImportChain(f *os.File) (*Chain, error) {
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	//	lines := []CsvLine{}
	idToLine := make(map[int]*CsvLine)
	var root *CsvLine
	var err error
	var record []string
	var line *CsvLine

	for {
		// Stop at EOF.
		if record, err = r.Read(); err == io.EOF {
			break
		}

		line, err = parseCsvLine(record)
		if err != nil {
			return nil, err
		}

		idToLine[line.id] = line
		if line.idParent == -1 {
			root = line
		}
	}

	//fmt.Println(root.ss)
	for _, csvChild := range idToLine {
		if csvChild.idParent != -1 {
			csvParent, _ := idToLine[csvChild.idParent]
			csvParent.ss.AddChild(csvChild.ss)
		}

	}

	chain := NewChain()
	chain.SetRootSubsystem(root.ss)
	return chain, nil
}

func parseCsvLine(record []string) (*CsvLine, error) {
	var err error
	var id, idParent int

	if id, err = strconv.Atoi(record[0]); err != nil {
		return nil, err
	}

	if idParent, err = strconv.Atoi(record[1]); err != nil {
		return nil, err
	}

	name := record[2]
	nick := record[3]

	var ssType SubsystemType
	if ssType, err = ParseSubsystemType(record[4]); err != nil {
		return nil, err
	}

	var impact float64
	if impact, err = parseDouble(record[5]); err != nil {
		return nil, err
	}

	minToWork, _ := strconv.Atoi(record[6])

	ss := NewSubsystem(name, nick)
	ss.SsType = ssType
	ss.Impact = impact
	ss.MinToWork = minToWork

	return &CsvLine{id, idParent, ss}, nil
}

func ExportChain(f *os.File, chain *Chain) {
	w := bufio.NewWriter(f)

	//	w.WriteString("id;idParent;Name;Nick;Tipo;Impacto;MinToWork")

	idToSubsystem := make(map[*Subsystem]int)

	var index int = 0
	chain.RootSubsystem().PreOrden(func(ss *Subsystem) {
		idToSubsystem[ss] = index
		index++
	})

	chain.RootSubsystem().PreOrden(func(ss *Subsystem) {

		idParent := -1
		if ss.parent != nil {
			idParent = idToSubsystem[ss.parent]
		}

		w.WriteString(strconv.Itoa(idToSubsystem[ss]))
		w.WriteString(";")
		w.WriteString(strconv.Itoa(idParent))
		w.WriteString(";")
		w.WriteString(ss.Name)
		w.WriteString(";")
		w.WriteString(ss.Nickname)
		w.WriteString(";")
		w.WriteString(ss.SsType.String())
		w.WriteString(";")
		w.WriteString(formatDouble(ss.Impact))
		w.WriteString(";")
		w.WriteString(strconv.Itoa(ss.MinToWork))
		w.WriteString("\n")
	})

	w.Flush()
}

func parseDouble(str string) (float64, error) {
	return strconv.ParseFloat(strings.Replace(str, ",", ".", -1), 64)
}

func formatDouble(value float64) string {
	return strings.Replace(strconv.FormatFloat(value, 'f', 6, 64), ".", ",", -1)
}
