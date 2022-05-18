package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type OpenWeatherMap struct {
	Lat      float64    `json:"lat"`
	Lon      float64    `json:"lon"`
	Timezone string     `json:"timezone"`
	Current  Current    `json:"current"`
	Minutely []Minutely `json:"minutely"`
	Hourly   []Hourly   `json:"hourly"`
	Daily    []Daily    `json:"daily"`
	Alerts   []Alerts   `json:"alerts"`
}

type Current struct {
	Temp       float64   `json:"temp"`
	FeelLike   float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	UVI        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	Windspeed  float64   `json:"wind_speed"`
	Winddeg    float64   `json:"wind_deg"`
	Windgust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Minutely struct {
	Dt            int `json:"dt"`
	Precipitation int `json:"precipitation"`
}

type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	Feelslike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	Dewpoint   float64   `json:"dew_point"`
	UVI        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	Windspeed  float64   `json:"wind_speed"`
	Winddeg    float64   `json:"wind_deg"`
	Windgust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	// "weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"
	// "rain":{"1h":0.43}}
}

type Daily struct {
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	Dewpoint  float64   `json:"dew_point"`
	Windspeed float64   `json:"wind_speed"`
	Winddeg   float64   `json:"wind_deg"`
	Weather   []Weather `json:"weather"`
}

type Alerts struct {
	SenderName  string `json:"sender_name"`
	Event       string `json:"event"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Description string `json:"description"`
}

type Weather struct {
	Id          int    `json:"id"`
	Name        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func main() {
	bytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(bytes)) // []byteで受け取るためstringに変換する必要がある
	var allweathers []OpenWeatherMap
	if err := json.Unmarshal(bytes, &allweathers); err != nil {
		log.Fatal(err)
	}

	for _, aw := range allweathers {
		fmt.Printf("%f, %f : %s\n", aw.Lat, aw.Lon, aw.Timezone)
		fmt.Printf("気温：%f°, 風速：%f m,  気圧：%d HP, 湿度：%d ％\n",
			aw.Current.Temp,
			aw.Current.Windspeed,
			aw.Current.Pressure,
			aw.Current.Humidity)
		// fmt.Println(aw.Daily[1].Pressure)
		for _, dayly := range aw.Daily {
			fmt.Println(dayly.Windspeed, dayly.Pressure) // Daily は３時間毎の表示だろう。
		}
		// fmt.Println(aw.Minutely[0])
	}
}
