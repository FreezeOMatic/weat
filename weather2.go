package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://community-open-weather-map.p.rapidapi.com/find?q=Moscow&units=metric"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	Result := &Resultstr{}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, Result)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(Result)

	fmt.Println("За окном:", Result.List[0].Weather[0].Description)
	fmt.Println("Ощущается как:", Result.List[0].Main.FeelsLike)
	fmt.Println("Облачность:", Result.List[0].Clouds.All)
	e := fmt.Sprintf("Ветер направления %d градуса, скоростью %.2f", Result.List[0].Wind.Deg, Result.List[0].Wind.Speed)
	fmt.Println(e)
	fmt.Println("Влажность:", Result.List[0].Main.Humidity)
	fmt.Println("Температура:", Result.List[0].Main.Temp)
}
