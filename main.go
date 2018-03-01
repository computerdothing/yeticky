package main

import (
	"encoding/csv"
	"strings"

	"yeticky/csvdata"
)

func main() {
	in1 := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	in2 := `"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	c1 := csvdata.NewCSV()
	c1.ReadAll(csv.NewReader(strings.NewReader(in1)), csvdata.HasHeader)
	c1.PrintColumns("last_name", "username")
	c1.PrintColumns("first_name", "username")
	c1.PrintColumns("username")

	c2 := csvdata.NewCSV()
	c2.ReadAll(csv.NewReader(strings.NewReader(in2)), csvdata.NoHeader)
	c2.PrintColumns("1", "2")
	c2.PrintColumns("0", "2")
	c2.PrintColumns("2")
}
