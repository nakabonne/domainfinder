package main

import (
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("synonyms"),
	exec.Command("sprinkle"),
	exec.Command("coolify"),
	exec.Command("domainify"),
	exec.Command("available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout
}
