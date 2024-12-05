package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func countLines(input string) int {
	return len(strings.Split(input, "\n"))
}

func wordCount(input string) int {
	return len(strings.Fields(input))
}

// Function to count characters
func countChars(input string) int {
	return len(input)
}

func ReadInputFromFileOrStdin(args []string) (string, error) {
	if len(args) > 0 {

		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return "", err
		}
		return string(data), nil
	}
	fmt.Println("Enter Input : ")
	var input string
	fmt.Scan(&input)
	return input, nil
}

func countFromInput(input string, CounCharFlag , CountLinesFlag , CountWordsFlag bool) {
	if CounCharFlag {
		fmt.Println("Characters Count Is :", countChars(input))
	}
	if CountLinesFlag {
		fmt.Println("Lines Count Is :", countLines(input))
	}
	if CountWordsFlag {
		fmt.Println("Words Count Is :", wordCount(input))
	}
}

var CountLinesFlag, CountWordsFlag, CountCharsFlag bool

func wcMainFunc(cmd *cobra.Command, args []string) error {
	input, err := ReadInputFromFileOrStdin(args)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return err
	}
	countFromInput(input, CountCharsFlag, CountLinesFlag, CountWordsFlag)
	return nil
}

var rootCmd = cobra.Command{
	Use:   "GO-WC",
	Short: "A simple word count tool",
	RunE:  wcMainFunc,
}

func init() {
	rootCmd.Flags().BoolVarP(&CountCharsFlag, "chars", "c", false, "count characters")
	rootCmd.Flags().BoolVarP(&CountLinesFlag, "lines", "l", false, "count lines")
	rootCmd.Flags().BoolVarP(&CountWordsFlag, "words", "w", false, "count words")
}

func main() {
	fmt.Println("Building WC in GO")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

}
