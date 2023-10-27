package helpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Execute() {
	args := os.Args[1:]
	if verifyArguments(args) {
		var file string
		if len(args) == 1 {
			file = "standard"
		} else {
			file = args[1]
		}
		if len(args[0]) > 0 && TextVerification(args[0]) {
			if args[0] == "\\n" {
				fmt.Println()
			} else {
				lines := strings.Split(args[0], "\\n")
				for i, line := range lines {
					if i < len(lines) - 1{
						DisplayChar(line, file)
					}
				}
				if lines[len(lines)-1] != ""{
					DisplayChar(lines[len(lines)-1], file)

				}
			}
		} else if len(args[0]) > 0 {
			fmt.Println("Votre texte contient des caracteres non pris en charge.")
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
	}

}

func verifyArguments(args []string) bool {
	if len(args) > 0 && len(args) <= 2 {
		return true
	}
	return false
}

func ReadFile(s string) []byte {
	var textByte []byte
	var err error

	if s != "standard" && s != "shadow" && s != "thinkertoy" {
		fmt.Println("INVALID BANNER")
		os.Exit(0)
	} else {
		textByte, err = os.ReadFile(s + ".txt")

		if err != nil {
			println("File not found")
			os.Exit(0)
		}
	}
	return textByte
}

func GetAllChar(file string) map[byte][]string {
	var char string 
	if file == "thinkertoy" {
		char = string(ReadFile(file))[2:]
	}else{
		char = string(ReadFile(file))[1:]
	}
	
	scanner := bufio.NewScanner(strings.NewReader(char))
	scannerTab := []string{}
	for scanner.Scan() {
		scannerTab = append(scannerTab, scanner.Text())
	}

	bowl := []string{}
	compt := 0
	allChar := map[byte][]string{}
	asciiNums := 32

	for _, v := range scannerTab {
		if compt != 8 {
			bowl = append(bowl, v)
			compt++
		} else {
			allChar[byte(asciiNums)] = bowl
			compt = 0
			bowl = []string{}
			asciiNums++
		}
	}
	return allChar
}

func GetSpecificChar(s string, file string) [][]string {
	tabChar := [][]string{}
	allChar := GetAllChar(file)
	for _, v := range s {
		tabChar = append(tabChar, allChar[byte(v)])
	}
	return tabChar
}

func DisplayChar(s string, file string) {
	tabChar := GetSpecificChar(s, file)
	if len(tabChar) > 0 {
		for line := range tabChar[0] {
			for char := range tabChar {
				fmt.Print(tabChar[char][line])
			}
			
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func TextVerification(s string) bool {
	re := regexp.MustCompile(`[^[:ascii:]]`)
	return len(re.FindAllString(s, -1)) == 0
}
