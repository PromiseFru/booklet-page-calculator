package cmd

import (
	"fmt"
	"os"

	"bookletctl/src"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookletctl",
	Short: "Calculate the Page(s) Order for your Booklet.",
	Long: `Manually Print Booklets on a Regular Single-Sided 
or Double-Sided Printer with no Extra Installation.
Calculate the Page(s) Order for your Booklet.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := src.Calculate("LTR", 2, 40, "DS")
		fmt.Print(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
