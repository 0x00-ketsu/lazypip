package pip

import (
	"os/exec"
	"strings"
)

// GetCommands extracts pip command info from `pip help`
func GetCommands() []string {
	helpInfos := executeHelp()
	start, end := 0, 0
	for idx, line := range helpInfos {
		if line == "Commands:" {
			start = idx + 1
		}

		if line == "General Options:" {
			end = idx
		}
	}

	commands := []string{}
	commandInfos := helpInfos[start:end]
	for _, item := range commandInfos {
		trimCommand := strings.TrimSpace(item)
		if len(trimCommand) > 0 {
			commands = append(commands, strings.Fields(trimCommand)[0])
		}
	}

	return commands
}

// Execute command `pip help`
// Return result lines
func executeHelp() []string {
	lines := []string{}

	out, err := exec.Command("pip", "help").Output()
	if err != nil {
		logger.Error("[pip] Execute command `pip help` failed: " + err.Error())
		return lines
	}

	return strings.Split(string(out[:]), "\n")
}
