package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
void setWindowFlags(long windowId) {
	NSView * view = (NSView *) windowId;
	NSWindow * nswindow = [view window];

	enum {
		NSWindowCollectionBehaviorFullScreenAuxiliary = 1 << 8,
		NSWindowCollectionBehaviorCanJoinAllSpaces = 1 << 0,
		NSWindowCollectionBehaviorStationary = 1 << 4
	};

	[nswindow setTitle:@"test string"];

	[nswindow setCollectionBehavior:NSWindowCollectionBehaviorCanJoinAllSpaces|NSWindowCollectionBehaviorFullScreenAuxiliary];
}
*/
import "C"

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello Widgets Example")

	//window.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	window.SetAttribute(core.Qt__WA_MacAlwaysShowToolWindow, true)

	window.SetWindowFlag(core.Qt__WindowStaysOnTopHint, true)
	window.SetWindowFlag(core.Qt__FramelessWindowHint, true)
	window.SetWindowFlag(core.Qt__Dialog, true)

	window.Show()

	C.setWindowFlags(C.long(window.WinId()))

	app.Exec()
}
