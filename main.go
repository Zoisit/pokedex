package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		more := scanner.Scan()
		if !more {
			fmt.Print("...")
		} else {
			text := strings.TrimSpace(scanner.Text())
			text = strings.ToLower(text)
			inputs := strings.Fields(text)
			fmt.Printf("Your command was: %s\n", inputs[0]) //careful, index out of bounds possible
		}
	}
}
