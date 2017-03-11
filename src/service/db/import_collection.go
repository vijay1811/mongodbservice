package db

import (
	"bytes"
	"os/exec"
	"service/protocol"
)

func (Default) ImportCollection(cip *protocol.ImportParam) (string, error) {
	cmdArgs := []string{
		// database
		"--db", cip.Database,
		// collection
		"--collection", cip.Collection,
		// file path
		"--file", cip.File,
	}
	cmd := exec.Command("mongoimport", cmdArgs...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stderr.String(), err
}
