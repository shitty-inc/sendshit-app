package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	icon "./icon"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/shitty-inc/sendshit-go"
	"github.com/sqweek/dialog"
)

func main() {
	onExit := func() {
		fmt.Println("onExit")
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.AppIcon)
	mOpen := systray.AddMenuItem("Open File", "Send a file")
	mScreenshot := systray.AddMenuItem("Take Screenshot", "Send a screenshot")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				filename, _ := dialog.File().Load()
				send(filename)
			case <-mScreenshot.ClickedCh:
				tmpFile := fmt.Sprintf("%sscreen.png", os.TempDir())

				cmd := exec.Command("screencapture", "-i", tmpFile)
				err := cmd.Run()

				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}

				send(tmpFile)
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
}
