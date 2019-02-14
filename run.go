// Package nmap provides the needed functions and structures to easily queue and manage nmap scans
package nmap

import (
	"context"
	"encoding/xml"
	"github.com/kulinacs/cliarg"
	"os/exec"
)

// Scan represents that options available when running an nmap scan
type Scan struct {
	Path      string `cliarg:"arg" json:"path"`
	Service   bool   `cliarg:"-sV" json:"service"`
	Ports     string `cliarg:"-p" json:"ports"`
	XMLOutput string `cliarg:"-oX" json:"xmloutput"`
	Host      string `cliarg:"arg" json:"host"`
}

// Run runs the scan and parses the output
func (s Scan) Run() (Result, []byte, error) {
	var result Result
	s.XMLOutput = "-"
	cmdline := cliarg.Marshal(s)
	output, err := exec.Command(cmdline[0], cmdline[1:]...).Output()
	if err != nil {
		return result, output, err
	}
	err = xml.Unmarshal(output, &result)
	return result, output, err
}

// RunCtx runs the scan with a given context
func (s Scan) RunCtx(ctx context.Context) (Result, []byte, error) {
	var result Result
	s.XMLOutput = "-"
	cmdline := cliarg.Marshal(s)
	output, err := exec.CommandContext(ctx, cmdline[0], cmdline[1:]...).Output()
	if err != nil {
		return result, output, err
	}
	err = xml.Unmarshal(output, &result)
	return result, output, err
}

// RunRaw runs the scan and returns the raw stdout output
func (s Scan) RunRaw() ([]byte, error) {
	cmdline := cliarg.Marshal(s)
	output, err := exec.Command(cmdline[0], cmdline[1:]...).Output()
	return output, err
}

// RunRawCtx runs the scan and returns the raw stdout output
func (s Scan) RunRawCtx(ctx context.Context) ([]byte, error) {
	cmdline := cliarg.Marshal(s)
	output, err := exec.CommandContext(ctx, cmdline[0], cmdline[1:]...).Output()
	return output, err
}
