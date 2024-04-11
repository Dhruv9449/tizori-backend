package main

import "github.com/GDGVIT/Tizori-backend/cmd"

func main() {
	tizoriApp := cmd.NewTizoriCliApp()
	tizoriApp.Run()
}
