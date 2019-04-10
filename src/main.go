package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.QUrl_FromLocalFile("./qml/main.qml"))
	gui.QGuiApplication_Exec()
}
