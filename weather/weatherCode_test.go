package weatherCode

import "testing"

func TestReadCityCodes(t *testing.T) {
    ReadCityCode()
}

func TestGetWeatherCode(t *testing) {
    w Area := {"台湾", "台北", "桃园"}
    wCode := GetWeatherCode(w)
    log.Printf("%v:%s\n", w, wCode)
}

