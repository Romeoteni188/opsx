package doctor

import (
	"os/exec"
)

func CheckDevelopment() Category {
	return Category{
		Name: "Development",
		Checks: []CheckResult{
			checkGo(),
			checkGit(),
		},
	}
}

func checkGo() CheckResult {
	return checkCommand("Go", "go", "version")
}

func checkGit() CheckResult {
	return checkCommand("Git", "git", "--version")
}

func checkCommand(name string, command string, args ...string) CheckResult {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		return CheckResult{
			Name:    name,
			Passed:  false,
			Message: "not found",
		}
	}

	return CheckResult{
		Name:    name,
		Passed:  true,
		Message: string(output),
	}
}
