package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func getBranches() ([]string, map[string]string) {
	cmd := exec.Command("git", "for-each-ref", "--sort=-committerdate", "refs/heads",
		"--format=%(refname:short) %(objectname:short) (%(committerdate:relative))")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("❌ %v", fmt.Errorf("error fetching branches: %w", err))
	}

	lines := strings.Split(out.String(), "\n")
	branches := []string{}
	branchInfo := make(map[string]string)

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		branch := parts[0]
		commitHash := parts[1]
		commitTime := strings.Join(parts[2:], " ")

		info := fmt.Sprintf("%s %s", commitHash, commitTime)
		branchInfo[branch] = info
		branches = append(branches, branch)
	}

	return branches, branchInfo
}

func getCurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("❌ %v", fmt.Errorf("error fetching current branch: %w", err))
	}
	return strings.TrimSpace(out.String())
}

func deleteBranch(branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	branches, branchInfo := getBranches()
	currentBranch := getCurrentBranch()

	deletableBranches := []string{}
	branchDisplayMap := make(map[string]string)

	for _, branch := range branches {
		if branch == currentBranch {
			continue
		}

		commitInfo := branchInfo[branch]
		formattedBranch := fmt.Sprintf("%s %s", color.HiGreenString(branch), color.YellowString(commitInfo))
		branchDisplayMap[formattedBranch] = branch
		deletableBranches = append(deletableBranches, formattedBranch)
	}

	if len(deletableBranches) == 0 {
		fmt.Println(color.HiGreenString("✅ No branches to delete."))
		return
	}

	var selectedBranches []string
	prompt := &survey.MultiSelect{
		Message: "Select branches to delete:",
		Options: deletableBranches,
	}

	err := survey.AskOne(prompt, &selectedBranches)
	if err != nil {
		fmt.Println(color.HiRedString("🛑 Cancelled"))
		return
	}

	if len(selectedBranches) == 0 {
		fmt.Println(color.HiYellowString("⚠️ No branches selected."))
		return
	}

	var confirmDelete bool
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("🔥 Are you sure you want to delete %d branch(es)?", len(selectedBranches)),
		Default: false,
	}

	survey.AskOne(confirmPrompt, &confirmDelete)
	if !confirmDelete {
		fmt.Println(color.HiYellowString("⚠️ Deletion cancelled. No branches were deleted."))
		return
	}

	fmt.Println(color.HiRedString("🔥 Deleting selected branches..."))
	for _, displayBranch := range selectedBranches {
		branch := branchDisplayMap[displayBranch]
		err := deleteBranch(branch)
		if err != nil {
			fmt.Printf("❌ Failed to delete %s: %v\n", branch, err)
		}
	}
}
