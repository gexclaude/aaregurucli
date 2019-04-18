package console

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"os/exec"
	"runtime"
)

// source: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
var clear map[string]func() //create a map for storing clear funcs
var colored = true

// InitConsole inits the console package and persists the colored flag
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

// BeforeExitConsole is a hook to call before console / program terminates
func BeforeExitConsole() {
	/*if runtime.GOOS == "windows" {
		fmt.Println("Press any key...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}*/
}

// ClearConsole clears previous output based on the given os
func ClearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		fmt.Println(runtime.GOOS)
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// color detection

// SupportsColors was supposed for terminal color detection - not used at the moment
func SupportsColors() {
	os.Getenv("COLORTERM")
}

// ---

// CBgBlue conditionally colored output with blue background
func CBgBlue(arg string) string {
	if isColored() {
		return aurora.BgBlue(arg).String()
	}
	return arg
}

// CBlue conditionally colored output in blue
func CBlue(arg string) string {
	if isColored() {
		return aurora.Blue(arg).String()
	}
	return arg
}

// CRed conditionally colored output in red
func CRed(arg string) string {
	if isColored() {
		return aurora.Red(arg).String()
	}
	return arg
}

// CGreen conditionally colored output in green
func CGreen(arg string) string {
	if isColored() {
		return aurora.Green(arg).String()
	}
	return arg
}

// CGray conditionally colored output in grey
func CGray(arg string) string {
	if isColored() {
		return aurora.Gray(20, arg).String()
	}

	return arg
}

// CBrown conditionally colored output in brown
func CBrown(arg string) string {
	if isColored() {
		return aurora.Brown(arg).String()
	}

	return arg
}

func isColored() bool {
	return colored
}
