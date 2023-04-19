package hello

import (
	"bufio"
	"fmt"
	"os"
)

// Input function
func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg+ ": ")
	scanner.Scan()
	return scanner.Text()

}
