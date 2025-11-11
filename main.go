package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetRequirements(pkg string) []string {
	cmd := exec.Command("py", "-m", "pip", "show", pkg)
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i := range strings.SplitSeq(string(result), "\r\n") {
		if strings.HasPrefix(i, "Requires: ") {
			return strings.Split(i[10:], ", ")
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
	var pkg string
	fmt.Scanln(&pkg)

	req := GetRequirements(pkg)
	fmt.Println(createGraph(pkg, req))
}
//Achievement
