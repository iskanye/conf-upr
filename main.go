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

func createGraph(name string, requirements []string) string {
	pkg := name
	result := "digraph " + pkg + " {"
	for _, v := range requirements {
		result += pkg + " -> " + v + ";\n"
	}
	result += "}"

	return result
}

func main() {
	req := GetRequirements()
	for _, i := range req {
		fmt.Println(i)
	}

	fmt.Println(createGraph("matplotlib", req))
}
