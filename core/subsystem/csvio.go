package subsystem

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CsvLine struct {
	id       int
	idParent int
	ss       *Subsystem
}

func ImportChain(f *os.File) *Chain {
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	//	lines := []CsvLine{}
	idToLine := make(map[int]CsvLine)
	var root *CsvLine

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		//		fmt.Println(record)

		id, _ := strconv.Atoi(record[0])
		idParent, _ := strconv.Atoi(record[1])
		name := record[2]
		nick := record[3]
		ssType, _ := ParseSubsystemType(record[4])
		impact := parseDouble(record[5])
		minToWork, _ := strconv.Atoi(record[6])

		ss := NewSubsystem(name, nick)
		ss.SsType = ssType
		ss.Impact = impact
		ss.MinToWork = minToWork

		line := CsvLine{id, idParent, ss}
		//		lines = append(lines, line)

		idToLine[id] = line

		if idParent == -1 {
			root = &line
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
	return chain
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
		w.WriteString(strconv.FormatFloat(ss.Impact, 'f', 6, 64))
		w.WriteString(";")
		w.WriteString(strconv.Itoa(ss.MinToWork))
		w.WriteString("\n")
	})

	w.Flush()
}

func parseDouble(s string) float64 {
	return float64(1)
}
