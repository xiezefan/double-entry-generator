package cmb_credit

import (
	"encoding/json"
	"fmt"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"io/ioutil"
	"log"
	"os"
)

type CmbCredit struct {
	Statistics Statistics `json:"statistics,omitempty"`
	LineNum    int        `json:"line_num,omitempty"`
	Orders     []Order    `json:"orders,omitempty"`
}

func New() *CmbCredit {
	return &CmbCredit{
		Statistics: Statistics{},
		LineNum:    0,
		Orders:     make([]Order, 0),
	}
}

func (c *CmbCredit) Translate(filename string) (*ir.IR, error) {
	log.SetPrefix("[Provider-CmbCredit] ")

	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err = json.Unmarshal([]byte(byteValue), &c.Statistics); err != nil {
		log.Fatalln(err)
	}


	for index, line := range c.Statistics.Data.Detail {
		err = c.translateToOrders(line)
		if err != nil {
			return nil, fmt.Errorf("failed to translate bill: line %d: %v", index, err)
		}
	}

	log.Printf("Finished to parse the file %s", filename)
	return c.convertToIR(), nil
}



