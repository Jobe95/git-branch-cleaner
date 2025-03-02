package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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

func deleteBranch(branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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

	deletableBranches := []string{}
	for _, branch := range branches {
		if branch != currentBranch {
			deletableBranches = append(deletableBranches, branch)
		}
	}

	if len(deletableBranches) == 0 {
		fmt.Println("No branches to delete.")
		return
	}

	var selectedBranches []string
	prompt := &survey.MultiSelect{
		Message: "Select branches to delete:",
		Options: deletableBranches,
	}

	err = survey.AskOne(prompt, &selectedBranches)
	if err != nil {
		fmt.Println("Cancelled")
		return
	}

	if len(selectedBranches) == 0 {
		fmt.Println("No branches selected.")
		return
	}

	for _, branch := range selectedBranches {
		err := deleteBranch(branch)
		if err != nil {
			fmt.Printf("Failed to delete %s: %v\n", branch, err)
		} else {
			fmt.Printf("Deleted %s\n", branch)
		}
	}

	fmt.Println("Done.")
}
