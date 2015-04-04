import QtQuick 2.4
import QtQuick.Controls 1.3
import QtQuick.Window 2.2
import QtQuick.Dialogs 1.2

ApplicationWindow {
    title: qsTr("Hello World")
    width: 360
    height: 200
    visible: true

    GroupBox {
        id: groupBox1
        x: 0
        y: 0
        width: 360
        height: 200
        title: qsTr("Login")

        Label {
            id: label1
            x: 77
            y: 29
            text: qsTr("ID : ")
        }

        Label {
            id: label2
            x: 30
            y: 57
            text: qsTr("Password : ")
        }

        TextField {
            id: idField
            objectName: "idField"
            x: 108
            y: 26
            width: 188
            height: 22
            placeholderText: qsTr("ID")
        }

        TextField {
            id: passwordField
            objectName: "passwordField"
            x: 108
            y: 54
            width: 188
            height: 22
            inputMask: qsTr("")
            echoMode: 2
            placeholderText: qsTr("Password")
        }

        Button {
            id: okBtn
            x: 220
            y: 101
            width: 76
            height: 26
            text: qsTr("OK")
            activeFocusOnPress: false
            isDefault: true
            onClicked: loginObj.handleOk()
        }

        Button {
            id: cancelBtn
            x: 138
            y: 101
            width: 76
            height: 26
            text: qsTr("Cancel")
            onClicked: loginObj.handleCancel()
        }
    }
}
