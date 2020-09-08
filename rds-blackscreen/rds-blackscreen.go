package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"golang.org/x/sys/windows/registry"
)

func main() {
	// Deletes all files under path.
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
	// Creates registry entry.
	y, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Terminal Server Client\Default`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
	    log.Fatal(err)
	}
	if err := y.SetDWordValue("bitmapCacheSize", 0x00000000); err != nil {
		log.Fatal(err)
	}
	if err := y.Close(); err != nil {
	    log.Fatal(err)
	}
}
