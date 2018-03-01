package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
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

// Header2Row adds the CSV header as the 1st row of rows
func (c *CSV) Header2Row() {
	var r row
	for i := 0; i < len(c.header); i++ {
		r[c.header[i]] = c.header[i]
	}
	c.rows = append([]row{r}, c.rows)
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
			return fmt.Errorf("Got error while reading csv header: %s", err.Error())
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
			return fmt.Errorf("Got error while reading a csv row: %s", err.Error())
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

// Format returns a new *CSV formatted per format from source *CSV c
func (c *CSV) Format(format string) *CSV {
	rv := NewCSV()

	splitFormat := strings.Split(format, ",")

	// set up rv header
	for i, head := range splitFormat {
		rv.header[i] = head
	}
	for _, rvr := range c.rows {
		r := make(row)
		for i, column := range splitFormat {
			if column == "" {
				r[rv.header[i]] = ""
			} else {
				r[rv.header[i]] = rvr[rv.header[i]]
			}
		}
		rv.rows = append(rv.rows, r)
	}
	return rv
}

// DeDup de duplicates the dupes and returns a nice deduplicated *CSV and is possibly done in an LTO way rn
func (c *CSV) DeDup() *CSV {
	rv := NewCSV()
	rv.header = c.header

	foundMap := make(map[string]bool)
	for _, r := range c.rows {
		var rStr string
		for i := 0; i < len(r); i++ {
			rStr = rStr + r[c.header[i]]
		}
		if !foundMap[rStr] {
			foundMap[rStr] = true
			rv.rows = append(rv.rows, r)
		}
	}
	return rv
}

func (c *CSV) Write(w *csv.Writer) error {
	// get rows to [][]string
	var writeRows [][]string
	for _, row := range c.rows {
		var strRow []string
		for i := 0; i < len(c.header); i++ {
			strRow = append(strRow, row[c.header[i]])
		}
		writeRows = append(writeRows, strRow)
	}

	// write
	w.WriteAll(writeRows)

	if err := w.Error(); err != nil {
		return err
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
