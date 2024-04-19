package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	
	

)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
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

	fmt.Println("ModifiedText written to", outputFile)
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

	_, err = file.WriteString(text )
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
	text = convertToUpper(text)

	// Rule 4: Lowercase
	text = convertToLower(text)
    //Rule 5:capitalised
	text = capitalised(text)

	// Rule 6: Handle punctuations
	text = fixPunctuation(text)

	// Rule 7: Replace 'a' with 'an'
	text = fixArticles(text)

	// Rule 8: Fix single quotes
	
    text = fixSingleQuotes(text)

	return text
}

// replaceHex Function to replace hexadecimal numbers
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
func replaceBin(text string) string {
	re := regexp.MustCompile(`(\b[01]+) \(bin\)`)
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		word := match[1]
		binValue, err := strconv.ParseInt(word, 2, 64)
		if err == nil {
			text = strings.Replace(text, match[0], fmt.Sprintf("%d", binValue), 1)
		}
	}
	return text
}

func capitalised(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(cap)") {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}else if strings.Contains(words[i], "(cap,") {
			caps, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			for j := i - caps; j < i; j++ {
				words[j] = strings.Title(words[j])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}

func convertToUpper(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up)") {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}else if strings.Contains(words[i], "(up,") {
			caps, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			for j := i - caps; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}

func convertToLower(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(low)") {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}else if strings.Contains(words[i], "(low,") {
			caps, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			for j := i - caps; j <= i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return strings.Join(words, " ")
}

func fixPunctuation(text string) string  {
    punctPattern := regexp.MustCompile(`(\s*)([.,!?:;]{1,3})(\s*)`)
    if !punctPattern.MatchString(text) {
        return text
    }
    text = strings.TrimSpace(punctPattern.ReplaceAllString(text, "$2 "))

    return text
}
  
func fixSingleQuotes (text string) string {
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
		nextWord := words[i+1]
		firstChar := nextWord[0]
		if words[i] == "a" && (firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h') {
			words[i] = "an"
		} else if words[i] == "A" && (firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h') {
			words[i] = "An"
		}
	}
	return strings.Join(words, " ")
}


