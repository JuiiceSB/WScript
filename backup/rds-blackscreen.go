package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// Deletes all files on this path.
	x := exec.Command("cmd", "/C", "DEL", "/Q", `%userprofile%\AppData\Local\Microsoft\Terminal Server Client\Cache\*.*`)
	var out bytes.Buffer
	var stderr bytes.Buffer
	x.Stdout = &out
	x.Stderr = &stderr
	err := x.Run()
	if err != nil {
		fmt.Println(x)
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())

}
