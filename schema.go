package main

import "encoding/xml"

// NmapRun represents the results of an nmap scan
type NmapRun struct {
	XMLName          xml.Name  `xml:"nmaprun"`
	Scanner          string    `xml:"scanner,attr"`
	Args             string    `xml:"args,attr"`
	Start            int       `xml:"start,attr"`
	StartStr         string    `xml:"startstr,attr"`
	Version          string    `xml:"version,attr"`
	XMLOutputVersion string    `xml:"xmpoutputversion,attr"`
	ScanInfo         ScanInfo  `xml:"scaninfo"`
	Verbose          Verbose   `xml:"verbose"`
	Debugging        Debugging `xml:"debugging"`
	Hosts            []Host    `xml:"host"`
}

// ScanInfo represents the metadata of an nmap scan
type ScanInfo struct {
	Type        string `xml:"type,attr"`
	Protocol    string `xml:"protocol,attr"`
	NumServices string `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
}

// Verbose represents the verbosity level of an nmap scan
type Verbose struct {
	Level int `xml:"level,attr"`
}

// Debugging represents the debugging level of an nmap scan
type Debugging struct {
	Level int `xml:"level,attr"`
}

// Host represents the results of a host in an nmap scan
type Host struct {
	StartTime int        `xml:"starttime,attr"`
	EndTime   int        `xml:"endtime,attr"`
	Status    Status     `xml:"status"`
	Address   Address    `xml:"address"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Ports     Ports      `xml:"ports"`
	Times     Times      `xml:"times"`
}

// Hostname represents a hostname of a host found in an nmap scan
type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// Status represents the status of a host found in an in nmap scan
type Status struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL int    `xml:"reason_ttl,attr"`
}

// Address represents the address of a host found in an nmap scan
type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

// Ports represents the ports of a host found in an nmap scan
type Ports struct {
	ExtraPorts ExtraPorts `xml:"extraports"`
	Ports      []Port     `xml:"port"`
}

// ExtraPorts represents the extra ports of a host found in an nmap scan
type ExtraPorts struct {
	State        string       `xml:"state,attr"`
	Count        int          `xml:"count,attr"`
	ExtraReasons ExtraReasons `xml:"extrareasons"`
}

// ExtraReasons represents the reasons a set of extra ports was marked as extra in an nmap scan
type ExtraReasons struct {
	Reason string `xml:"reason,attr"`
	Count  int    `xml:"count,attr"`
}

// Port represents a port found on a host in an nmap scan
type Port struct {
	Protocol string  `xml:"protocol,attr"`
	PortID   int     `xml:"portid,attr"`
	State    State   `xml:"state"`
	Service  Service `xml:"services"`
}

// State represents the state of a port found in an nmap scan
type State struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL int    `xml:"reason_ttl,attr"`
}

// Service represents the service of port found in an nmap scan
type Service struct {
	Name   string `xml:"name,attr"`
	Method string `xml:"method,attr"`
	Conf   int    `xml:"conf,attr"`
}

// Times represents the times for a host in an nmap scan
type Times struct {
	SRTT   int `xml:"srtt,attr"`
	RTTVar int `xml:"rttvar,attr"`
	To     int `xml:"to,attr"`
}
