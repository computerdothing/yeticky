// package cmd is where the commands are
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"yeticky/csvdata"
	"yeticky/util"

	"github.com/spf13/cobra"
)

var HasHeader bool
var DoDeDup bool

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "csv command is for outputting another csv file",
	Long: `Manipulate and format the input csv file into an output
				csv file. It's fun.
				yeticky csv <in-file> <out-file> "how,you,,want,format" --flags-perhaps`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		infile := args[0]
		outfile := args[1]
		format := args[2]

		// open input file like a dumb guy
		in, err := os.Open(infile)
		if err != nil {
			fmt.Printf("Error opening file %s : %s", infile, err.Error())
			return
		}
		defer in.Close()

		// read that file into a CSV struct thing
		c := csvdata.NewCSV()
		c.ReadAll(csv.NewReader(in), HasHeader)

		// Get that formatted CSV from the Format()
		oc := c.Format(format)

		// if dodeeedoop then do it
		if DoDeDup {
			oc = oc.DeDup()
		}

		// open output file like a dumb guy
		if util.Exists(outfile) {
			fmt.Printf("Error! File %s already exists.", outfile)
		}
		of, err := os.Create(outfile)
		if err != nil {
			fmt.Printf("Error opening file %s : %s", outfile, err.Error())
			return
		}
		defer of.Close()

		// If there's a header, stick it on the output CSV
		if HasHeader {
			oc.Header2Row()
		}

		// write our final joint to the outfile
		if err = oc.Write(csv.NewWriter(of)); err != nil {
			fmt.Println("Error writing to outfile %s : %s", outfile, err.Error())
		}
	},
}

func init() {
	csvCmd.Flags().BoolVarP(&HasHeader, "hasheader", "H", false, "If set, will treat first line of CSV infile as the header")
	csvCmd.Flags().BoolVarP(&DoDeDup, "dedup", "d", false, "Set this flag to only get unique rows after formatting")
}
