package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type APIResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
}

func main() {
	city := strings.ToLower(strings.Join(os.Args[1:], " "))
	template := `http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=810da73052a9b959fade35cde8894228`
	url := fmt.Sprintf(template, city)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response APIResponse
	json.Unmarshal(body, &response)

	kelvin := response.Main.Temp
	celcius := kelvin - 273.15

	if celcius > 15 {
		fmt.Println("No Hoodie")
	} else {
		fmt.Println("Hoodie")
	}
}
