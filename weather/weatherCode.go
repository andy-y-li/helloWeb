package weatherCode

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	//"strconv"
)

type StringResources struct {
	XMLName        xml.Name         `xml:"China"`
	ResourceString []ResourceString `xml:"province"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"province"`
	StringID   string   `xml:"id,attr"`
	StringName string   `xml:"name,attr"`
	CityList   []City   `xml:"city"`
	//InnerText string `xml:",innerxml"`
}

type City struct {
	CityName   xml.Name     `xml:"city"`
	StringID   string       `xml:"id,attr"`
	StringName string       `xml:"name,attr"`
	CountyList []CountyCode `xml:"county"`
}

type CountyCode struct {
	CountyName  xml.Name `xml:"county"`
	StringID    string   `xml:"id,attr"`
	StringName  string   `xml:"name,attr"`
	WeatherCode string   `xml:"weatherCode,attr"`
}

type Area struct {
	ProvinceName string
	CityName     string
	CountyName   string
}

func ReadCityCode() {
	content, err := ioutil.ReadFile("cityCodes.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(result)
	//log.Println(result.ResourceString)
	for _, line := range result.ResourceString {
		log.Println("province-id:" + line.StringID + " name:" + line.StringName)
		for _, citys := range line.CityList {
			log.Println("city-id:" + citys.StringID + " name:" + citys.StringName)
			for _, counties := range citys.CountyList {
				log.Println("county id:" + counties.StringID + " name:" + counties.StringName + " weatherCode:" + counties.WeatherCode)
			}

		}
	}
}

func GetWeatherCode(where Area) (wCode string, err error) {
	if where.CountyName == "" {
		where.CountyName = where.CityName
	}
	content, err := ioutil.ReadFile("cityCodes.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range result.ResourceString {
		if line.StringName != where.ProvinceName {
			continue
		}
		for _, citys := range line.CityList {
			if citys.StringName != where.CityName {
				continue
			}
			for _, counties := range citys.CountyList {
				if counties.StringName == where.CountyName {
					return counties.WeatherCode, nil
				}
			}
		}
	}
	return "", errors.New("Not Found!")
}
