package main

import (
	"bufio"
	"fmt"
	"io"
	"github.com/charmbracelet/log"
	"os/exec"
)

func main() {

cmd := exec.Command("journalctl", "-f", "-o", "json")

stdout, err := cmd.StdoutPipe()
if err != nil {
	log.Error("Error getting StdoutPipe: %v", err)
}

if err := cmd.Start(); err != nil {
	log.Error("Error starting journalctl: %v", err)
}

// reader buffer
reader := bufio.NewReader(stdout)

log.Info("streaming logs...")

for {
	line, err := reader.ReadString('\n')

	// TODO: Parse certain logs out
	if err != nil {
		if err == io.EOF {
			log.Info("end of stream...")
		} else {
			log.Error("error reading from journalctl:", err)
		}
		
}
}
