import QtQuick 2.6
import QtQuick.Window 2.2
import QtQuick.Controls 2.1
import Qt.labs.platform 1.0

ApplicationWindow {
    id: root
    visible: true
    width: Screen.desktopAvailableWidth
    height: Screen.desktopAvailableHeight + 500
    opacity: 0.5
    flags: Qt.WindowStaysOnTopHint | Qt.FramelessWindowHint | Qt.Dialog

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        hoverEnabled: false

        onPressed: {
            marker.x = mouseArea.mouseX
            marker.y = mouseArea.mouseY
            hoverEnabled = true;
        }

        onPositionChanged: {
            marker.width = mouseArea.mouseX - marker.x
            marker.height = mouseArea.mouseY - marker.y
        }

        onReleased: {
            hoverEnabled = false
            root.visible = false
            areaSelected(marker.x, marker.y, marker.width, marker.height)
        }
    }

    Rectangle {
        id: marker
        width: 0
        height: 0
        color: 'green'
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
