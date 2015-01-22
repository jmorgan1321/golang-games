import QtQuick 2.0
import GoExtensions 1.0

Rectangle {
    id: screen
    color: "black"
    focus: true

    MouseArea {
        anchors.fill: parent

        onClicked: console.log("Stop poking me!")
        onPressed: ctrl.handleMouseDown(mouse.x, mouse.y)
        onReleased: ctrl.handleMouseDown(mouse.x, mouse.y)
    }

    Keys.onPressed: {
        var shift_mod = (event.modifiers&Qt.ShiftModifier)!=0;
        var ctrl_mod  = (event.modifiers&Qt.ControlModifier)!=0;
        var alt_mod   = (event.modifiers&Qt.AltModifier)!=0;
        ctrl.handleKey(event.key, shift_mod, ctrl_mod, alt_mod);
    }

    Rectangle {
        width: 250; height: 250
        anchors.centerIn: parent
        color: "red"

        MouseArea {
            anchors.fill: parent
            onClicked: console.log("Stop poking me!")
        }

    }

    Timer {
        id: t1
        objectName: "t1"
        interval: 16; running: false; repeat: true;
        onTriggered: {
            if (running) {
                ctrl.update();
            }
        }
    } // t1

} // screen
