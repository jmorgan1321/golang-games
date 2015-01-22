import QtQuick 2.0

Rectangle {
    id: screen
    color: "black"
    focus: true

    Timer {
        id: t1
        objectName: "t1"
        interval: 1000; running: false; repeat: true;
        onTriggered: {
            if (running) {
                gameCore.step();
            }
        }
    } // gameTimer

    MouseArea {
        anchors.fill: parent

        onClicked: console.log("Stop poking me!")
        onPressed: inputMngr.handleMouseDown(mouse.x, mouse.y)
        onReleased: inputMngr.handleMouseDown(mouse.x, mouse.y)
    }

    Keys.onPressed: {
        var shift_mod = (event.modifiers&Qt.ShiftModifier)!=0;
        var ctrl_mod  = (event.modifiers&Qt.ControlModifier)!=0;
        var alt_mod   = (event.modifiers&Qt.AltModifier)!=0;
        inputMngr.handleKey(event.key, shift_mod, ctrl_mod, alt_mod);
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


}
