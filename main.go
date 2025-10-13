package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("py", "-m", "pip", "show", "matplotlib")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))
}
