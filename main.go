package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--format", "%(refname:short)")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(out.String(), "\n")
	var cleaned []string
	for _, b := range branches {
		b = strings.TrimSpace(b)
		if b != "" {
			cleaned = append(cleaned, b)
		}
	}
	return cleaned, nil
}

func main() {
	branches, err := getBranches()

	if err != nil {
		fmt.Println("Error fetching branches:", err)
		os.Exit(1)
	}

	fmt.Println(branches)
}
