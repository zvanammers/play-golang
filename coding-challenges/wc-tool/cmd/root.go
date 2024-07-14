/*
Copyright © 2024 Zoe
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type ThingToCount int

const (
	bytes ThingToCount = iota
	words
	runes
	lines
	all
)

const ARG_FLAG = "ARG_FLAG"

func countBytes(text []byte) int {
	return len(text)
}

func countChars(text []byte) int {
	return utf8.RuneCountInString(string(text))
}

func countLines(textSource []byte) int {
	text := string(textSource)

	lines := strings.Split(text, "\n")
	return len(lines)
}

func countWords(text []byte) int {
	processed := []byte{}
	for _, r := range text {
		if unicode.IsSpace(rune(r)) {
			processed = append(processed, ' ')
		} else {
			processed = append(processed, r)
		}
	}

	words := strings.Split(string(processed), " ")
	count := 0

	for _, word := range words {
		if word != "" {
			count++
		}

	}
	return count
}

func getFileName(reader io.Reader, args []string, bytesFile string, charsFile string, linesFile string, wordsFile string) (text []byte, thingToCount ThingToCount, file string) {

	if term.IsTerminal(int(os.Stdin.Fd())) && len(args) == 0 && (bytesFile == ARG_FLAG || charsFile == ARG_FLAG || linesFile == ARG_FLAG || wordsFile == ARG_FLAG) {
		fmt.Println("Error: Cannot provide flag, but no file or stdin")
		os.Exit(1)
	} else if term.IsTerminal(int(os.Stdin.Fd())) && len(args) == 0 {
		fmt.Println("Error: Must provide stdin or file")
		os.Exit(1)
	}

	thingToCount = bytes

	if charsFile != "" {
		thingToCount = runes
	} else if linesFile != "" {
		thingToCount = lines
	} else if wordsFile != "" {
		thingToCount = words
	} else if bytesFile != "" {
		thingToCount = bytes
	} else {
		thingToCount = all
	}

	// file is provided
	if len(args) == 1 && (bytesFile == ARG_FLAG || charsFile == ARG_FLAG || linesFile == ARG_FLAG || wordsFile == ARG_FLAG) {
		fileName := args[0]
		text, errReadFile := os.ReadFile(fileName)

		if errReadFile != nil {
			fmt.Printf("Error when reading file with flag: %s\n", errReadFile)
			os.Exit(1)
		}
		return text, thingToCount, fileName
	} else if len(args) == 1 {
		text, errReadFile := os.ReadFile(args[0])

		if errReadFile != nil {
			fmt.Printf("Error when reading file with no flags provided: %s\n", errReadFile)
			os.Exit(1)
		}
		return text, all, args[0]
	} else if len(args) > 1 {
		fmt.Println("Cannot provide more than one file")
		os.Exit(1)
	}

	// std input
	text, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return text, thingToCount, ""
}

var fileToCountBytes string
var fileToCountChars string
var fileToCountLines string
var fileToCountWords string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc-tool",
	Short: "Calculates word count of a text file",
	Long:  `Given a text file, calculates number of words, chars, and bytes`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		num := 0
		var text, thingToCount, fileName = getFileName(cmd.InOrStdin(), args, fileToCountBytes, fileToCountChars, fileToCountLines, fileToCountWords)

		if thingToCount == bytes {
			num = countBytes(text)
		} else if thingToCount == runes {
			num = countChars(text)
		} else if thingToCount == words {
			num = countWords(text)
		} else if thingToCount == lines {
			num = countLines(text)
		} else if thingToCount == all {
			byteCount := countChars(text)
			lineCount := countLines(text)
			wordCount := countWords(text)
			fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, fileName)
			os.Exit(0)
		}

		fmt.Printf("%d %s\n", num, fileName)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wc-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&fileToCountBytes, "c", "c", "", "Provide file to calculate number of bytes")
	rootCmd.Flags().StringVarP(&fileToCountChars, "m", "m", "", "Provide file to calculate number of chars")
	rootCmd.Flags().StringVarP(&fileToCountLines, "l", "l", "", "Provide file to calculate number of lines")
	rootCmd.Flags().StringVarP(&fileToCountWords, "w", "w", "", "Provide file to calculate number of words")
	rootCmd.Flags().Lookup("c").NoOptDefVal = ARG_FLAG
	rootCmd.Flags().Lookup("m").NoOptDefVal = ARG_FLAG
	rootCmd.Flags().Lookup("l").NoOptDefVal = ARG_FLAG
	rootCmd.Flags().Lookup("w").NoOptDefVal = ARG_FLAG
}
