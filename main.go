package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Element struct {
	Humidity float64   `json:"humidity"`
	Weather  []Weather `json:"weather"`
}
type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type OneCall struct {
	Date    string
	Current struct {
		Element
		DateUnix CurrentDate `json:"dt"`
		Sunrise  Hour        `json:"sunrise"`
		Sunset   Hour        `json:"sunset"`
		Temp     Temperature `json:"temp"`
	} `json:"current"`
	Hourly []struct {
		Element
		DateUnix HourOnly    `json:"dt"`
		Temp     Temperature `json:"temp"`
	} `json:"hourly"`
	Daily []struct {
		Element
		DateUnix Day `json:"dt"`
		Temp     struct {
			Day Temperature `json:"day"`
			Min Temperature `json:"min"`
			Max Temperature `json:"max"`
		} `json:"temp"`
	} `json:"daily"`
}

func getOneCall(api string) (*OneCall, error) {
	r, err := http.Get("https://api.openweathermap.org/data/2.5/onecall?lat=27.98&lon=86.92&exclude=minutely,alerts&units=metric&appid=" + api)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var d *OneCall
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		return nil, err
	}
	d.Date = time.Time(d.Current.DateUnix).Format("2006-01-02, 15:04:05")
	return d, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("need an openwathermap.org API key as arg & port")
		os.Exit(0)
	}
	apiKey := os.Args[1]
	port := os.Args[2]

	tmpl := template.Must(template.ParseFiles("template.html"))

	fs := http.FileServer(http.Dir("icons"))
	http.Handle("/icons/", http.StripPrefix("/icons/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := getOneCall(apiKey)
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, data)
	})
	fmt.Println("serving at :" + port)
	http.ListenAndServe(":"+port, nil)
}
