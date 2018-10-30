package filehandling

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func makeChain() {
	// Load a TXT file.
	f, _ := os.Open("test.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}
