package pip

import (
	"fmt"
	"strings"

	"github.com/0x00-ketsu/lazypip/internal/config"
	"github.com/0x00-ketsu/lazypip/internal/contrib/logging"
	"go.uber.org/zap"
)

var (
	conf   *config.Config
	logger *zap.Logger
)

type Package struct {
	Name        string
	Version     string
	Released    string
	Description string
	Link        string
}

func init() {
	conf, _ = config.Load()
	logger = logging.GetLogger()
}

func IsPackageInstalled(name string) (bool, string) {
	installedPkgs := List()
	for _, pkg := range installedPkgs {
		if strings.ToLower(pkg.Name) == strings.ToLower(name) {
			return true, fmt.Sprintf("%s\t%s", pkg.Name, pkg.Version)
		}
	}

	return false, ""
}
