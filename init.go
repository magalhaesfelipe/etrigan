// https://www.gofaq.org/en/how-to-build-a-cli-application-in-go/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("How is it going,", os.Args[2],"?")
	} else {
		fmt.Println("How is it going pal?")
	}
}
