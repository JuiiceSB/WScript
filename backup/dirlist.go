package main

import (
  "os"
  "fmt"
)

func deleteCache() {
  dir := "C:\Users\%AppData%\Local\Microsoft\Terminal Server Client\Cache"

  opendir, _ := os.Open(dir)
  listdir, _ := opendir.Readdir(0)

  for index := range(listdir) {
    listed := listdir[index]

    filename := listed.Name()
    fullpath := dir + filename

    os.Remove(fullpath)
    fmt.Println("File removed:", fullpath)
  }
