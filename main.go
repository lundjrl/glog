package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

// TODO: Need other

type SystemLog struct {
	_PID              string `json:"_PID"`
	MESSAGE           string `json:"MESSAGE"`
	SYSLOG_IDENTIFIER string `json:"SYSLOG_IDENTIFIER"`
}

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
		if err != nil {
			if err == io.EOF {
				log.Info("end of stream...")
				os.Exit(0)
			} else {
				log.Error("error reading from journalctl:", err)
			}
		}

		jsonData := []byte(line)

		var sysLog SystemLog
		er := json.Unmarshal(jsonData, &sysLog)

		if er != nil {
			log.Error("error parsing JSON:", er)
		}

		// TODO: Parse certain logs out
		if len(sysLog.MESSAGE) > 0 {
			msg := sysLog.SYSLOG_IDENTIFIER + " " + sysLog.MESSAGE
			log.Info(msg)
		}
	}
}
