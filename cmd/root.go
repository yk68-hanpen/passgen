package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "PassGen - A password generator CLI tool",
	Long: `PassGen is a CLI tool for generating strong passwords.
It can generate multiple passwords with customizable options.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
}
