package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"

	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
)

func main() {
	onExit := func() {
		fmt.Println("onExit")
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Sendshit")
	mOpen := systray.AddMenuItem("Send file", "Send file")
	mScreenshot := systray.AddMenuItem("Send screenshot", "Send screenshot")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				filename, _ := dialog.File().Load()
				if err := clipboard.WriteAll(filename); err != nil {
					panic(err)
				}
			case <-mScreenshot.ClickedCh:
				EvChan := hook.Start()
				defer hook.End()

				var topX int = 0
				var topY int = 0
				dragging := false

				for ev := range EvChan {
					if ev.Kind == hook.MouseHold && int(ev.Button) == 1 {
						topX = int(ev.X)
						topY = int(ev.Y)
						dragging = true

						fmt.Println(ev)
					}

					if dragging && ev.Kind == hook.MouseDown && int(ev.Button) == 1 {
						fmt.Println(ev)

						img, err := screenshot.CaptureRect(image.Rect(topX, topY, int(ev.X)-topX, int(ev.Y)-topY))
						if err != nil {
							panic(err)
						}

						fmt.Println(topX, topY, int(ev.X)-topX, int(ev.Y)-topY)

						file, _ := os.Create("test.png")
						defer file.Close()
						png.Encode(file, img)

						break
					}
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}
