package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var (
	length       int
	useSpecial   bool
	useNumbers   bool
	useUppercase bool
	useLowercase bool
	count        int
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate strong passwords",
	Long: `Generate strong passwords with specified options.
You can customize the length, character types, and number of passwords to generate.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			password := generatePassword(length, useSpecial, useNumbers, useUppercase, useLowercase)
			fmt.Printf("Password %d: %s\n", i+1, password)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Add flags for password generation options
	generateCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	generateCmd.Flags().BoolVarP(&useSpecial, "special", "s", true, "Include special characters")
	generateCmd.Flags().BoolVarP(&useNumbers, "numbers", "n", true, "Include numbers")
	generateCmd.Flags().BoolVarP(&useUppercase, "uppercase", "u", true, "Include uppercase letters")
	generateCmd.Flags().BoolVarP(&useLowercase, "lowercase", "w", true, "Include lowercase letters")
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of passwords to generate")
}

// generatePassword creates a strong password based on the specified options
func generatePassword(length int, useSpecial, useNumbers, useUppercase, useLowercase bool) string {
	var chars string

	if useLowercase {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if useUppercase {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useNumbers {
		chars += "0123456789"
	}
	if useSpecial {
		chars += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}

	// Default to lowercase if nothing is selected
	if chars == "" {
		chars = "abcdefghijklmnopqrstuvwxyz"
	}

	return generateRandomString(length, chars)
}

// generateRandomString creates a random string of specified length using the provided character set
func generateRandomString(length int, charset string) string {
	result := strings.Builder{}
	result.Grow(length)

	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			panic(err)
		}
		result.WriteByte(charset[randomIndex.Int64()])
	}

	return result.String()
}
