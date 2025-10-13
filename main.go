package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetRequirements() []string {
	cmd := exec.Command("py", "-m", "pip", "show", "matplotlib")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i := range strings.SplitSeq(string(result), "\n") {
		if strings.HasPrefix(i, "Requires: ") {
			return strings.Split(strings.TrimLeft(i, "Requires: "), ", ")
		}
	}

	return nil
}

func main() {
	for _, i := range GetRequirements() {
		fmt.Println(i)
	}
}
