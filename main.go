package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    weather "github.com/andy-y-li/helloWeb/weather"
    "io/ioutil"
)



type WeatherInfo map[string]string

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
        city := strings.SplitN(r.URL.Path, "/", 3)[2]
        //fmt.Println("city:", city)
        data, err := query(city)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
    })

    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

/*
{
	"weatherinfo":{
		"city":"深圳",
		"cityid":"101280601",
		"temp1":"14℃",
		"temp2":"23℃",
		"weather":"晴",
		"img1":"n0.gif",
		"img2":"d0.gif",
		"ptime":"18:00"
		}
}
*/
func query(city string) (WeatherInfo, error) {
    wCode, err := weather.GetWeatherCodeByCounty(city)
    if err != nil {
        fmt.Println(err)
        return WeatherInfo{},err
    }
    qstr := "http://www.weather.com.cn/data/cityinfo/" + wCode + ".html"
    fmt.Println(qstr)
    resp, err := http.Get(qstr)
    if err != nil {
        return WeatherInfo{}, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return WeatherInfo{}, err
    }
    str := string(body)
    var s map[string]WeatherInfo
    fmt.Println(str)
    err = json.Unmarshal(body, &s)
    if err != nil {
        return WeatherInfo{}, err
    }
    info := s["weatherinfo"]
    fmt.Println(info)
    fmt.Println(info["weather"])

    return info, nil
}

