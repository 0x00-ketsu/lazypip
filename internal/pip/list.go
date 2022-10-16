package pip

import (
	"os/exec"
	"strings"
)

func List() []Package {
	packages := []Package{}

	out, err := exec.Command("pip", "list").Output()
	if err != nil {
		logger.Error("[pip list] Execute command `pip list` failed: " + err.Error())
		return packages
	}

	items := strings.Split(string(out[:]), "\n")
	if len(items) <= 2 {
		return packages
	}

	for i := 2; i < len(items); i++ {
		item := strings.TrimSpace(items[i])
		if len(item) > 0 {
			rows := strings.Fields(item)
			packages = append(packages, Package{
				Name:    rows[0],
				Version: rows[1],
			})
		}
	}

	return packages
}
