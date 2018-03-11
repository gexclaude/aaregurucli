package console

import (
	"os/exec"
	"os"
	"runtime"
	"fmt"
	. "github.com/logrusorgru/aurora"
)

// source: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
var clear map[string]func() //create a map for storing clear funcs
var colored = true

func InitConsole(coloredParam bool) {
	colored = coloredParam
	
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

func BeforeExitConsole() {
	/*if runtime.GOOS == "windows" {
		fmt.Println("Press any key...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}*/
}

func ClearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		fmt.Println(runtime.GOOS)
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// color detection

func SupportsColors() {
	os.Getenv("COLORTERM")
}

// ---

func CBgBlue(arg string) string {
	if isColored() {
		return BgBlue(arg).String()
	} else {
		return arg
	}
}

func CBlue(arg string) string {
	if isColored() {
		return Blue(arg).String()
	} else {
		return arg
	}
}

func CRed(arg string) string {
	if isColored() {
		return Red(arg).String()
	} else {
		return arg
	}
}

func CGreen(arg string) string {
	if isColored() {
		return Green(arg).String()
	} else {
		return arg
	}
}

func CGray(arg string) string {
	if isColored() {
		return Gray(arg).String()
	} else {
		return arg
	}
}

func CBrown(arg string) string {
	if isColored() {
		return Brown(arg).String()
	} else {
		return arg
	}
}

func isColored() bool {
	return colored
}
