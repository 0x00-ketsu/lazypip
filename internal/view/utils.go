package view

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/0x00-ketsu/lazypip/internal/utils"
	"github.com/creack/pty"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var blankCell = tview.NewTextView()

func makeButton(label string, color tcell.Color, handler func()) *tview.Button {
	button := tview.NewButton(label).
		SetSelectedFunc(handler).
		SetLabelColor(color)

	button.SetBackgroundColor(tcell.ColorCornflowerBlue)

	return button
}

// Execute pip command then wirte result to tview with async
// Invoke callback when execute is finished
func executePipCmdAync(cmdArgs []string, ouput io.Writer, callback func()) {
	ch := make(chan bool)

	cmd := exec.Command("pip", cmdArgs...)
	f, err := pty.Start(cmd)
	if err != nil {
		logger.Error(fmt.Sprintf("Execute pip command failed: %v", err.Error()))
		msg := fmt.Sprintf("Execute pip command: %v, failed: %v", cmdArgs, err.Error())
		statuslineView.showForSeconds(msg, 5)
	}

	go func() {
		w := tview.ANSIWriter(ouput)
		_, err := io.Copy(w, f)
		if err != nil {
			ch <- true
		}
	}()

	go func() {
		done := <-ch
		if done {
			if !utils.IsNil(callback) {
				callback()
			}
		}
	}()
}
