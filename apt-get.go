package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/exp/slices"
)

var banned_args = []string{"--no-install-recommends"}

func main() {

	args := []string{}

	for i := range os.Args {
		if i >= 1 {
			if !slices.Contains(banned_args, os.Args[i]) {
				args = append(args, os.Args[i])
			}

		}
	}

	if !slices.Contains(args, "-y") {
		args = append(args, "-y")
	}

	if slices.Contains(args, "install") {
		args = append(args, "--skip-broken")
	}

	fmt.Println("apt-get-shim -> dnf", args)

	cmd := exec.Command("dnf", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
