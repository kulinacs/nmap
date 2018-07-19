package main

import "encoding/xml"
import "testing"
import "io/ioutil"

// https://www.markphelps.me/testing-api-clients-in-go/
func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestNmapRun(t *testing.T) {
	document := fixture("nmaprun.xml")

	var v NmapRun
	err := xml.Unmarshal([]byte(document), &v)
	if err != nil {
		t.Error(err)
	}
	if v.XMLName.Local != "nmaprun" {
		t.Errorf("XML Unmarshal was incorrect, got: %s, want: %s.", v.XMLName.Local, "nmaprun")
	}
}
