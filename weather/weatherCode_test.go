package weatherCode

import (
	"fmt"
	"testing"
)

func TestReadCityCodes(t *testing.T) {
	ReadCityCode()
}

func TestGetWeatherCode(t *testing.T) {
	w := Area{"台湾", "台北", "桃园"}
	wCode, err := GetWeatherCode(w)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v, weatherCode:%s\n", w, wCode)
	}
}
