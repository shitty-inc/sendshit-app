package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/0xAX/notificator"
	"github.com/ProtonMail/go-autostart"
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/shitty-inc/sendshit-app/icon"
	"github.com/shitty-inc/sendshit-go"
	"github.com/sqweek/dialog"
)

var notify *notificator.Notificator

func main() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/icon.png",
		AppName:     "Sendshit",
	})

	systray.Run(onReady, nil)
}

func onReady() {
	ex, _ := os.Executable()

	app := &autostart.App{
		Name:        "Sendshit",
		DisplayName: "Sendshit menu bar app",
		Exec:        []string{ex},
	}

	systray.SetIcon(icon.AppIcon)

	mOpen := systray.AddMenuItem("Open File", "Send a file")
	mScreenshot := systray.AddMenuItem("Take Screenshot", "Send a screenshot")

	systray.AddSeparator()

	mToggle := systray.AddMenuItem("Start at Login", "Toggle starting the app on Login")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	if app.IsEnabled() {
		mToggle.SetTitle("Do not start at Login")
	} else {
		mToggle.SetTitle("Start at Login")
	}

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				filename, _ := dialog.File().Load()
				send(filename)
			case <-mScreenshot.ClickedCh:
				randName, _ := sendshit.GenerateRandomString(12)
				tmpFile := fmt.Sprintf("%sscreenshot-%s.png", os.TempDir(), randName)

				cmd := exec.Command("screencapture", "-i", tmpFile)
				err := cmd.Run()

				if err != nil {
					log.Fatalf("Couldn't capture that shit %s\n", err)
				}

				send(tmpFile)
				os.Remove(tmpFile)
			case <-mToggle.ClickedCh:
				if app.IsEnabled() {
					app.Disable()
					mToggle.SetTitle("Start at Login")
				} else {
					app.Enable()
					mToggle.SetTitle("Do not start at Login")
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func send(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Couldn't read that shit %s\n", err)
	}

	key, err := sendshit.GenerateRandomString(24)

	if err != nil {
		log.Fatalf("Couldn't generate a key for that shit %s\n", err)
	}

	encodedStr, err := sendshit.EncryptFile(filepath.Base(path), file, key)

	if err != nil {
		log.Fatalf("Couldn't encrypt that shit %s\n", err)
	}

	response, err := sendshit.UploadFile(encodedStr)

	if err != nil {
		log.Fatalf("Couldn't upload that shit %s\n", err)
	}

	if err := clipboard.WriteAll(fmt.Sprintf("https://sendsh.it/#/%s/%s\n", response.ID, key)); err != nil {
		panic(err)
	}

	notify.Push("", "Link has been copied to the clipboard!", "", notificator.UR_NORMAL)
}
