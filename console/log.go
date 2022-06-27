package console

import "fmt"

func Intro() {
	Print("?Starting Ron!", bCyan)
}

func SettingUp() {
	Print("?Configuring watcher.", bCyan)
}

func Listening() {
	Print("?Listening for changes...", bCyan)
}

func Restarting() {
	Print("?Restarting...", bCyan)
}

func Error(message string) {
	Print(fmt.Sprintf("?[ERROR] ?%s", message), bRed, bCyan)
}
