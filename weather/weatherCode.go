package weatherCode

import (
    "encoding/xml"
    "io/ioutil"
    "log"
    "strconv"
)

type StringResources struct {
    XMLName xml.Name `xml:"China"`
    ResourceString []ResourceString `xml:"province"`
}

type ResourceString struct {
    XMLName xml.Name `xml:"province"`
    StringName string `xml:"id,attr"`
    InnerText string `xml:",innerxml"`
}

func ReadCityCode() {
    content, err := ioutil.ReadFile("cityCodes.xml")
    if err != nil {
        log.Fatal(err)
    }
    var  result StringResources
    err = xml.Unmarshal(content, &result)
    if err != nil {
        log.Fatal(err)
    }
    //log.Println(result)
    //log.Println(result.ResourceString)
    for i, line := range result.ResourceString {
        log.Println("string[" + strconv.Itoa(i) + "]:" + line.StringName + "===" + line.InnerText)
    }
}


