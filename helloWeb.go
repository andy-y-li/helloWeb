package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    weather "github.com/andy-y-li/helloWeb/weather"
    "io/ioutil"
)

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

type info struct {
    city string `json:"city"`
    cityid string `json:"cityid"`
    temp1 string `json:"temp1"`
    temp2 string `json:"temp2"`
    weather string `json:"weather"`
    img1 string `json:"img1"`
    img2 string `json:"img2"`
    ptime string `json:"ptime"`
}

type weatherData struct {
    weatherInfo info `json:"weatherinfo"`
}

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

func query(city string) (weatherData, error) {
    wCode, err := weather.GetWeatherCodeByCounty(city)
    if err != nil {
        fmt.Println(err)
        return weatherData{},err
    }
    qstr := "http://www.weather.com.cn/data/cityinfo/" + wCode + ".html"
    fmt.Println(qstr)
    resp, err := http.Get(qstr)
    if err != nil {
        return weatherData{}, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return weatherData{}, err
    }
    var s info
    json.Unmarshal([]byte(string(body)), &s)
    fmt.Println(string(body))
    fmt.Println(s)

    var d weatherData

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
        return weatherData{}, err
    }

    return d, nil
}

