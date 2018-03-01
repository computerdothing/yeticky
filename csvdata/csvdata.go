package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

const HasHeader = true
const NoHeader = false

type row map[string]string

type CSV struct {
	rows   []row
	header map[int]string
}

func NewCSV() *CSV {
	return &CSV{header: make(map[int]string)}
}

// ReadAll reads in everything from a provided csv.Reader csvr, if hasHeader is true, treats first
// line as header and uses header for key in row (blank headers get column number as string as key),
// if false will just use column number as string for key
func (c *CSV) ReadAll(csvr *csv.Reader, hasHeader bool) error {
	// Handle the header
	if hasHeader {
		record, err := csvr.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Got error while reading csv header: %s", err.Error)
		}
		for i, v := range record {
			if v == "" {
				c.header[i] = strconv.Itoa(i)
			} else {
				c.header[i] = v
			}
		}
	}

	// Read dem rows
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Got error while reading a csv row: %s", err.Error)
		}

		// if we didn't set up a header, do it now
		if len(c.header) == 0 {
			for i, _ := range record {
				c.header[i] = strconv.Itoa(i)
			}
		}

		r := make(row)
		for i, v := range record {
			r[c.header[i]] = v
		}
		c.rows = append(c.rows, r)
	}
	return nil
}

// PrintColumns is probably just a toy for testing.
func (c *CSV) PrintColumns(columns ...string) {
	for _, row := range c.rows {
		for i, column := range columns {
			if i == len(columns)-1 {
				fmt.Printf("%s\n", row[column])
				break
			}
			fmt.Printf("%s,", row[column])
		}
	}
}
