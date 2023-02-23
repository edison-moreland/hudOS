package main

import "encoding/xml"

type XMLProtocol struct {
	XMLName    xml.Name       `xml:"protocol"`
	Name       string         `xml:"name,attr"`
	Copyright  string         `xml:"copyright"`
	Interfaces []XMLInterface `xml:"interface"`
}

type XMLInterface struct {
	XMLName     xml.Name       `xml:"interface"`
	Name        string         `xml:"name,attr"`
	Version     string         `xml:"version,attr"`
	Description XMLDescription `xml:"description"`
	Requests    []XMLRequest   `xml:"request"`
	Events      []XMLEvent     `xml:"event"`
	Enums       []XMLEnum      `xml:"enum"`
}

type XMLRequest struct {
	XMLName     xml.Name       `xml:"request"`
	Name        string         `xml:"name,attr"`
	Type        string         `xml:"type,attr"`  // "destructor"
	Since       string         `xml:"since,attr"` // The version this request was added
	Description XMLDescription `xml:"description"`
	Arguments   []XMLArgument  `xml:"arg"`
}

type XMLEvent struct {
	XMLName     xml.Name       `xml:"event"`
	Name        string         `xml:"name,attr"`
	Type        string         `xml:"type,attr"`
	Description XMLDescription `xml:"description"`
	Arguments   []XMLArgument  `xml:"arg"`
}

type XMLArgument struct {
	XMLName   xml.Name `xml:"arg"`
	Name      string   `xml:"name,attr"`
	Summary   string   `xml:"summary,attr"`
	Type      string   `xml:"type,attr"`
	Interface string   `xml:"interface,attr"` // Required if Type="new_id" or Type="object
	Enum      string   `xml:"enum,attr"`
	AllowNull string   `xml:"allow-null,attr"`
}

type XMLEnum struct {
	XMLName     xml.Name       `xml:"enum"`
	Name        string         `xml:"name,attr"`
	Description XMLDescription `xml:"description"`
	Entries     []XMLEntry     `xml:"entry"`
}

type XMLEntry struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
	Summary string   `xml:"summary,attr"`
}

type XMLDescription struct {
	XMLName xml.Name `xml:"description"`
	Summary string   `xml:"summary,attr"`
	Value   string   `xml:",chardata"`
}
