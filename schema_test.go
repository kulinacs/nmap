package nmap

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

// https://www.markphelps.me/testing-api-clients-in-go/
func fixture(path string) []byte {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return b
}

func TestNmapRun(t *testing.T) {
	document := fixture("nmaprun.xml")

	var v Result
	err := xml.Unmarshal(document, &v)
	if err != nil {
		t.Error(err)
	}
	if v.XMLName.Local != "nmaprun" {
		t.Errorf("XML Unmarshal was incorrect, got: %s, want: %s.", v.XMLName.Local, "nmaprun")
	}
}
