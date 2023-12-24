package homework_day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command struct {
	FileName string
	Action   string
}

func main() {
	//if err := Run(Command{FileName: "example.txt", Action: "byte_count"}); err != nil {
	//	panic("there are some error: " + err.Error())
	//}
	//
	//if err := Run(Command{FileName: "example.txt", Action: "line_count"}); err != nil {
	//	panic("there are some error: " + err.Error())
	//}
	//
	//if err := Run(Command{FileName: "example.txt", Action: "word_count"}); err != nil {
	//	panic("there are some error: " + err.Error())
	//}

	cmd := getCommand()

	if err := Run(cmd); err != nil {
		panic("there are some error: " + err.Error())
	}
}
func getCommand() Command {
	// TODO implement
	/*
		Link: https://gobyexample.com/command-line-arguments
		command : ./ccwc -c example.txt
	*/
	var action string

	args := os.Args[1:]
	switch args[0] {
	case "-c":
		action = "byte_count"
	case "-l":
		action = "line_count"
	case "-w":
		action = "word_count"
	}

	return Command{
		FileName: args[1],
		Action:   action,
	}
}

func Run(cmd Command) error {
	switch cmd.Action {
	case "byte_count":
		countNumberOfBytes(cmd.FileName)
	case "line_count":
		countNumberOfLines(cmd.FileName)
	case "word_count":
		countNumberOfWords(cmd.FileName)
	default:
		return fmt.Errorf("run error: Not implement action: %s", cmd.Action)
	}

	return nil
}

func countNumberOfBytes(file string) {
	// TODO implement
	fmt.Println("TODO implement countNumberOfBytes")
	readResult := ReadFile(file)
	fmt.Println("NumberOfBytes: ", len(readResult))
}

func countNumberOfLines(file string) {
	// TODO implement
	fmt.Println("TODO implement countNumberOfLines")
	readResult := ReadFile(file)
	line := 0
	for _, char := range readResult {
		if char == '\n' {
			line += 1
		}
	}
	fmt.Println("NumberOfLines: ", line)
}

func countNumberOfWords(file string) {
	// TODO implement
	fmt.Println("TODO implement countNumberOfWords")
	readResult := ReadFile(file)
	words := 0

	for _, char := range readResult {
		if char == ' ' || char == '\n' {
			words += 1
		}
	}

	words += 1
	fmt.Println("NumberOfWords: ", words)
}

func ReadFile(filepath string) string {
	/*
		Link: https://yourbasic.org/golang/read-file-line-by-line/
	*/
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	// Remove last newline character
	result = result[:len(result)-1]

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
