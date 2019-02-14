// +build integration

package nmap

import (
	"testing"
)

func TestRun(t *testing.T) {
	s := Scan{Path: "/usr/bin/nmap",
		Service:   true,
		XMLOutput: "-",
		Host:      "127.0.0.1"}
	result, _, _ := s.Run()
	if result.XMLName.Local != "nmaprun" {
		t.Errorf("XML Unmarshal was incorrect, got: %s, want: %s.", result.XMLName.Local, "nmaprun")
	}
}
