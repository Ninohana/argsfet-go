package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		processName := os.Args[1]
		commandLine := ObtainProcessCommandArgs(processName)
		fmt.Print(commandLine)
		file, _ := os.Create(fmt.Sprintf("command_line_%s.txt", processName))
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(commandLine)
		writer.Flush()
	}
}
