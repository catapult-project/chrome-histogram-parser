package main

import "fmt"
import "os"
import "encoding/xml"
import "io/ioutil"

type histogram struct {
	XMLName  xml.Name `xml:"histogram"`
	Name     string   `xml:"name,attr"`
	Owner    []string `xml:"owner"`
	Summary  string   `xml:"summary"`
	Obsolete string   `xml:"obsolete"`
}

func (histogram histogram) String() string {
	return fmt.Sprintf("%s owners: %s", histogram.Name, histogram.Owner)
}

type histogramConfiguration struct {
	XMLName    xml.Name    `xml:"histogram-configuration"`
	Histograms []histogram `xml:"histograms>histogram"`
}

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])

	var histogramConfiguration histogramConfiguration
	xml.Unmarshal(file, &histogramConfiguration)

	for _, histogram := range histogramConfiguration.Histograms {
		if histogram.Obsolete != "" {
			continue
		}
		fmt.Println(histogram)
	}
}
