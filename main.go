package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputText, err := readFromFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	modifiedText := modifyText(inputText)

	err = writeToFile(outputFile, modifiedText)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Modifications complete. Output written to", outputFile)
}

// Function to read text from a file
func readFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var textBuilder strings.Builder
	for scanner.Scan() {
		textBuilder.WriteString(scanner.Text() + "\n")
	}

	return textBuilder.String(), scanner.Err()
}

// Function to write text to a file
func writeToFile(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

// Function to modify the text based on the rules provided
func modifyText(text string) string {
	// Rule 1: Replace hexadecimal numbers
	text = replaceHex(text)

	// Rule 2: Replace binary numbers
	text = replaceBin(text)

	// Rule 3: Uppercase
	text = replaceCase(text, "up")

	// Rule 4: Lowercase
	text = replaceCase(text, "low")

	// Rule 5: Capitalize
	text = replaceCase(text, "cap")

	// Rule 6: Handle punctuations
	text = fixPunctuation(text)

	// Rule 7: Replace 'a' with 'an'
	text = fixArticles(text)

	// Rule 8: Fix single quotes
	text = fixSingleQuotes(text)

	text = fixDoubleQuotes(text)

	return text
}

// Function to replace hexadecimal numbers
func replaceHex(text string) string {
	re := regexp.MustCompile(`(\b[0-9A-Fa-f]+) \(hex\)`)
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		word := match[1]
		hexValue, err := strconv.ParseInt(word, 16, 64)
		if err == nil {
			text = strings.Replace(text, match[0], fmt.Sprintf("%d", hexValue), 1)
		}
	}
	return text
}

// Function to replace binary numbers


func fixPunctuation(text string) string  {
    punctPattern := regexp.MustCompile(`(\s*)([.,!?:;]{1,3})(\s*)`)
    if !punctPattern.MatchString(text) {
        return text
    }
    text = strings.TrimSpace(punctPattern.ReplaceAllString(text, "$2 "))

    return text
}
  
// Function to fix single quotes
func fixSingleQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(\w+)\s*'`)
	text = re.ReplaceAllString(text, "'$1'")
	return text
}


func fixDoubleQuotes (text string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	text = re.ReplaceAllStringFunc(text, func(s string) string {
		words := strings.Fields(s)
		quotedWords := make([]string, 0)
		for _, word := range words {
			if word != "'" {
				quotedWords = append(quotedWords, word)
			}
		}
		return fmt.Sprintf("'%s'", strings.Join(quotedWords, " "))
	})
	return text
}

	
// Function to replace 'a' with 'an' before vowels or 'h'

func fixArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" && (strings.HasPrefix("aeiou", string(words[i+1][0])) || strings.HasPrefix("h", string(words[i+1][0]))) {
			if words[i] == "a" {
				words[i] = "an"
			} else if words[i] == "A" {
				words[i] = "An"
			}
		}
	}
	return strings.Join(words, " ")
}	



