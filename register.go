package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-ecms/generator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("------ Simple Register ------")
	fmt.Println("Input register num: (1:Platform 2:Module)")
	mode, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	switch mode {
	case "1\n":
		generator.RegisterPlatform()
	case "2\n":
		fmt.Println("WIP")
		return
	default:
		fmt.Println("Invalid register")
		return
	}
}
