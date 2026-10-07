package commands

import (
	"fmt"
	"os"
)

var notFoundMsg = "Command not found"

func Run() {
	if len(os.Args) < 2 {
		fmt.Println(notFoundMsg)
		return
	}

	command := os.Args[1]


	

	



	fmt.Println("Command:", command)
}
