package cmb

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"io"
	"log"
	"os"
)

type CMB struct {
	Statistics Statistics `json:"statistics,omitempty"`
	LineNum    int        `json:"line_num,omitempty"`
	Orders     []Order    `json:"orders,omitempty"`
}

// New creates a new CMB provider.
func New() *CMB {
	return &CMB{
		Statistics: Statistics{},
		LineNum:    0,
		Orders:     make([]Order, 0),
	}
}

func (w *CMB) Translate(filename string) (*ir.IR, error) {
	log.SetPrefix("[Provider-CMB] ")

	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	// If FieldsPerRecord is negative, no check is made and records
	// may have a variable number of fields.
	reader.FieldsPerRecord = -1

	var account string

	for {
		line, err := reader.Read()
		for index, item := range line {
			line[index] = item
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		w.LineNum++

		if w.LineNum == 3 {
			account = convertAccount(line[0])
		}

		if w.LineNum <= 7 {
			// The first 8 lines are useless for us.
			continue
		}

		err = w.translateToOrders(line, account)
		if err != nil {
			return nil, fmt.Errorf("failed to translate bill: line %d: %v", w.LineNum, err)
		}
	}
	log.Printf("Finished to parse the file %s", filename)
	return w.convertToIR(), nil
}
