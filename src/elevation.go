package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func checkRoot() bool {
	http.DefaultTransport.(*http.Transport).ResponseHeaderTimeout = time.Second * 1
	_, err := http.Get("http://localhost:7070/health")

	if err != nil {

		fmt.Println("Server not running, starting it up")

		fmt.Println("Asking user for password")

		var password = ""
		var widget = widgets.NewQWidget(nil, 0)
		var dialog = widgets.NewQInputDialog(widget, core.Qt__Dialog)
		dialog.SetWindowTitle("ThreeFold Network Connector")
		dialog.SetLabelText("ThreeFold Network Connector would like to automatically\nset up your connection to the ThreeFold Network.\n\nTo do this,  please provide the password for \"" + getUsername() + "\"")
		dialog.SetTextEchoMode(widgets.QLineEdit__Password)
		dialog.SetInputMethodHints(core.Qt__ImhNone)

		dialog.ConnectAccepted(func() {
			fmt.Println("Accepted")
			password = dialog.TextValue()
			dialog.Close()
		})

		dialog.ConnectRejected(func() {
			fmt.Println("Rejected")
			os.Exit(1)
		})

		dialog.Exec()

		fmt.Println("Starting server as root")
		elevateMyself(password)
	}
	time.Sleep(2 * time.Second)
	_, err2 := http.Get("http://localhost:7070/health")
	if err2 != nil {
		checkRoot()
	}
	return false
}

func cleanupYggdrasilSock(password string) string {

	cmd := "echo " + password + " | sudo -S rm -rf /var/run/yggdrasil.sock"
	// fmt.Println(cmd)
	stdout, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
	return strings.TrimSpace(string(stdout))
}

func elevateMyself(password string) {
	ex, errp := os.Executable()
	if errp != nil {
		panic(errp)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	cmd := "echo " + password + " | sudo -S \"" + ex + "\" -server"
	rcmd := exec.Command("bash", "-c", cmd)
	err := rcmd.Start()

	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}

}

func getUsername() string {
	return os.Getenv("USER")
}
