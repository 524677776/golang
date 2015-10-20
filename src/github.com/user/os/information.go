package main 

import "fmt"
import "os"
import "strings"

func main() {
	var env []string

	env = os.Environ()

	fmt.Println("List of Environtment variables:\n")

	for  index, value := range env{
		name := strings.Split(value, "=")

		fmt.Printf("[%d]%s: %v\n", index, name[0], name[1])
	}
}