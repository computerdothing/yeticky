// package cmd is where the commands are
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yeticky",
	Short: "yeticky is for chopping CSV files",
	Long: `I just need to chop up some CSV files and I'm not going
				to copy and paste between excel files NO MORE.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(csvCmd)
}
