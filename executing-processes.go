package main

import "syscall"
import "os/exec"
import "os"

func main() {

	binary, _ := exec.LookPath("ls")

	args := []string{"xxxx", "-l", "-h"}
	// actually, the first argument of args will always be ignored
	// because Exec think is as the name of program

	envs := os.Environ()

	err := syscall.Exec(binary, args, envs)
	if err != nil {
		panic(err)
	}

}
