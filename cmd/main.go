package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//mygrep "text" file flags

type File struct{
	fileName string
	fileContent []string
}

func scanDir(dir string) ([]File, error) {
	var files []File

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := readFileContents(path)
			if err != nil {
				return err
			}
			files = append(files, File{fileName: path, fileContent: content})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}


func readFileContents(fileName string) ([]string, error) {
	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil,err
	}
	defer file.Close()
	// read the file line by line
	scanner := bufio.NewScanner(file)
	
	var fileContent []string
	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}
	return fileContent, nil
}

func caseSensitiveSearch(fileName string,content []string, strToMatch string,countMode bool) {
	red := "\033[31m"
	reset := "\033[0m"
	countOccurences := 0
	for linenum, line := range content {
		if strings.Contains(line, strToMatch) {
			countOccurences += 1
			highlightedLine := strings.ReplaceAll(line, strToMatch, red+strToMatch+reset)
			if countMode{
				continue
			}
			fmt.Printf("%s %d: %s\n",fileName, linenum+1, highlightedLine)
		}
	}
	fmt.Println("Total occurences:", countOccurences)
	fmt.Printf("\n\n")
}

func caseInsensitiveSearch(fileName string,content []string, strToMatch string, countMode bool) {
    red := "\033[31m"
    reset := "\033[0m"
    countOccurences := 0
    for linenum, line := range content {
        lowerLine := strings.ToLower(line)
        lowerStrToMatch := strings.ToLower(strToMatch)
        if strings.Contains(lowerLine, lowerStrToMatch) {
            countOccurences += 1
            startIndex := strings.Index(lowerLine, lowerStrToMatch)
            endIndex := startIndex + len(strToMatch)
            highlightedLine := line[:startIndex] + red + line[startIndex:endIndex] + reset + line[endIndex:]
            if countMode {
                continue
            }
            fmt.Printf("%s %d: %s\n",fileName, linenum+1, highlightedLine)
        }
    }
	fmt.Println("Total occurences:", countOccurences)
	fmt.Printf("\n\n")
    
}

func searchText(fileName string, content []string, strToMatch string, modes map[string]bool) {	
    
	isCaseSensitive := true
	if _, ok := modes["-i"]; ok {
		isCaseSensitive = false
	}
	countMode := false
	if _, ok := modes["-c"]; ok {
		countMode = true
	}
	
	
    if isCaseSensitive {
		caseSensitiveSearch(fileName,content, strToMatch,countMode)
	} else {
		caseInsensitiveSearch(fileName,content, strToMatch,countMode)
	}
	
}

func main() {
	// parse the command line arguments
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Usage: mygrep <text> <file> <flags>")
		return
	}
	strToMatch := args[1]
	dir := args[2]
	fmt.Println("Searching for:", strToMatch)
	fmt.Println("Directory:", dir)
	var flags []string
	if len(args) > 3 {
		flags = args[3:]
	}
	modes := make(map[string]bool)
	// isCaseSensitive := true
	// fmt.Println("Flags:", flags)
	for _, flag := range flags {
		if flag == "-i" {
			fmt.Println("Ignoring case")
			// isCaseSensitive = false
			modes["-i"] = true
		}
		if flag == "-c" {
			fmt.Println("Counting occurences")
			modes["-c"] = true
		}
	}
	// fmt.Println("Case sensitive:", isCaseSensitive)
	
	files,err := scanDir(dir)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	for _, file := range files {
		fmt.Println("File:", file.fileName)
		content := file.fileContent
		searchText(file.fileName,content, strToMatch, modes)
		// fmt.Println("Content:", content)
	}
}