package cmd

import (
	"fmt"
	"os"

	"github.com/0x00-ketsu/lazypip/internal/config"
	"github.com/0x00-ketsu/lazypip/internal/contrib/logging"
	"github.com/0x00-ketsu/lazypip/internal/utils"
	"github.com/0x00-ketsu/lazypip/internal/view"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lazypip",
	Short: "A simple terminal user interface interactive with Python pip.",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Load()
		if err != nil {
			fmt.Printf("Load config settings failed: %v\n", err.Error())
			os.Exit(1)
		}

		app := tview.NewApplication().EnableMouse(conf.App.Mouse)
		layout := view.Load(app)
		if err := app.SetRoot(layout, true).Run(); err != nil {
			panic(err)
		}

	},
}

// Execute executes for root command
func Execute() error {
	if ok, msg := valid(); !ok {
		logger.Error(msg)
		os.Exit(0)
	}
	return rootCmd.Execute()
}

func valid() (bool, string) {
	if !utils.IsCommandExist("pip") {
		msg := "Failed: command `pip` is not found."
		logger.Error(msg)
		return false, msg
	}

	return true, ""
}
