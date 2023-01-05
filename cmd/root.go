package cmd

import (
	"fmt"
	"os"

	"bookletctl/src"

	"github.com/spf13/cobra"
)

var Mode int
var Direction string
var Count int
var Type string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookletctl",
	Short: "Calculate the Page(s) Order for your Booklet.",
	Long: `Manually Print Booklets on a Regular Single-Sided 
or Double-Sided Printer with no Extra Installation.
Calculate the Page(s) Order for your Booklet.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Mode < 1 {
			fmt.Print("required flag --mode(-m): must be an integer greater than 1\n\n")
			os.Exit(1)
		} else if Direction != "LTR" && Direction != "RTL" {
			fmt.Print("required flag --direction(-d): must be either LTR(left-to-right) or RTL(right-to-left)\n\n")
			os.Exit(1)
		} else if Type != "SS" && Type != "DS" {
			fmt.Print("required flag --type(-t): must be either SS(Single-Sided) or DS(Double-Sided)\n\n")
			os.Exit(1)
		} else if Count < 1 {
			fmt.Print("required flag --count(-c): must be an integer greater than 1\n\n")
			os.Exit(1)
		} else {
			result := src.Calculate(Direction, Mode, Count, Type)
			fmt.Print(result)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&Mode, "mode", "m", 0, "Page(s) Per Side (required)")
	rootCmd.MarkFlagRequired("mode")
	rootCmd.Flags().StringVarP(&Direction, "direction", "d", "", "LTR(left-to-right) (common) or RTL(right-to-left) (Arabic, Hebrew, Farsi etc.) (required)")
	rootCmd.MarkFlagRequired("direction")
	rootCmd.Flags().StringVarP(&Type, "type", "t", "", "SS(Single-Sided) or DS(Double-Sided) printer (required)")
	rootCmd.MarkFlagRequired("type")
	rootCmd.Flags().IntVarP(&Count, "count", "c", 0, "Number of Page(s) to be printed (required)")
	rootCmd.MarkFlagRequired("count")
}
