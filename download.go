package main

import (
	"errors"
	"log"
	"os/exec"
)

func downloadyou(name string) (string, error) {
	log.Println("Downloading started, arg: \"", name+"\"")
	out, _ := exec.Command("python3", "downloadyou.py", name).Output()
	log.Println("Executed")
	if string(out) == "" {
		return "", errors.New("couldn't download " + name)
	}
	return string(out), nil
}
