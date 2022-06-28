package console

import "fmt"

func Intro() {
	Print("?Starting Ron!", bPurple)
}

func SettingUp() {
	Print("?Configuring watcher.", bPurple)
}

func Listening() {
	Print("?Listening for changes...", bPurple)
}

func Restarting() {
	Print("?Restarting...", bGreen)
}

func Error(message string) {
	Print(fmt.Sprintf("?[ERROR] ?%s", message), bRed, bPurple)
}
