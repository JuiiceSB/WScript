package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Deletes all files on this path.
	x := exec.Command("cmd", "/C")

	x.Stdout = os.Stdout
	stdin, err := x.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, `del /Q "C:\Users\%UserName%\AppData\Local\Microsoft\Terminal Server Client\Cache\*.*"`)
	}()
	err = x.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
