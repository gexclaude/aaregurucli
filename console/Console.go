package console

import (
	"os/exec"
	"os"
	"runtime"
	"bufio"
	"fmt"
)

// source: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go

var clear map[string]func() //create a map for storing clear funcs

func Init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func BeforeExit() {
	if runtime.GOOS == "windows" {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		fmt.Println(runtime.GOOS)
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}