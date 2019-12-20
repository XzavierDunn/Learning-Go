package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Data struct
type Data struct {
	Coords     Coord     `json:"coord"`
    Weathers   map[Weather]interface{} `json:"weather"`
	Base       string
	Mains      MainData `json:"main"`
	Visibility int
}

// Coord struct
type Coord struct {
	Lon float64
	Lat float64
}

// Weather struct
type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

// MainData struct
type MainData struct {
	Temp      float64
	FeelsLike int
	TempMin   int
	TempMax   int
	Pressure  int
	Humidity  int
}

func main() {
	url := os.Getenv("API_URL")
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

    data := Data{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(data)

//    fmt.Println("Current Weather:", data.Weathers[0].Main, "\n" +
//    "Description:", data.Weathers[0].Description)
}
