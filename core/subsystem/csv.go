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

func MakeChain(f *os.File) *Chain {
	// Load a TXT file.
	//f, _ := os.Open("test.csv")

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
		ssType, _ := ParseType(record[3])
		impact, _ := strconv.ParseFloat(record[5], 64)
		minToWork, _ := strconv.Atoi(record[6])

		ss := NewSubsystem(name, nick)
		ss.Type = ssType
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
