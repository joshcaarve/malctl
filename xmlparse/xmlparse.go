package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UrlSetXML struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Url     []UrlXML `xml:"url"`
}
type UrlXML struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	ChangeFreq string   `xml:"changefreq"`
	Priority   string   `xml:"priority"`
}

func GetUrlSetXML(url string) UrlSetXML {
	var result UrlSetXML
	if xmlBytes, err := getXMLHttp(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &result)
	}
	return result
}

func getXMLHttp(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
