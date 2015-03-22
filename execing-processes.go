package main

import (
	"os"
	"os/exec"
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	lsPath, pathErr := exec.LookPath("ls")
	check(pathErr)

	args := []string{"ls", "-a", "-h", "-t", "-l", "/home"}

	env := os.Environ()

	lsErr := syscall.Exec(lsPath, args, env)
	check(lsErr)
}
