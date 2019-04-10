import QtQuick 2.6
import QtQuick.Window 2.2
import QtQuick.Controls 2.1
import Qt.labs.platform 1.0

ApplicationWindow {
    id: root
    visible: true
    width: Screen.desktopAvailableWidth
    height: Screen.desktopAvailableHeight + 500
    title: qsTr("TestApplication")
    opacity: 0.5
    flags: Qt.WindowStaysOnTopHint | Qt.BypassWindowManagerHint | Qt.FramelessWindowHint | Qt.NoDropShadowWindowHint | Qt.ToolTip

    Rectangle{
        id: fullScreenRectangle
        width: Screen.width
        height: Screen.height
        border.color: "red"
        border.width: 5
        z: 1
    }

    SystemTrayIcon {
        id: systemTrayIcon
        visible: true
        iconSource: "qrc:/qml/icon.png"

        menu: Menu {
            enabled: true
            MenuItem {
                text: qsTr("Screenshot")
                onTriggered: root.show()
            }
            MenuItem {
                text: qsTr("Quit")
                onTriggered: Qt.quit()
            }
        }
    }
}
