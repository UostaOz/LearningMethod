package keybord

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Keybord() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal()
	}
	input = strings.TrimSpace(input)
	return input
}
