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

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func main() {
	branches, err := getBranches()

	if err != nil {
		fmt.Println("Error fetching branches:", err)
		os.Exit(1)
	}

	currentBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Error fetching current branch:", err)
		os.Exit(1)
	}

	fmt.Println("Local Git branches:")
	for _, branch := range branches {
		if branch == currentBranch {
			fmt.Printf("  * %s (active)\n", branch)
		} else {
			fmt.Printf("  - %s\n", branch)
		}
	}
}
