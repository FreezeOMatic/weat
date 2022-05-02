package main

type Resultstr struct {
	Cod     string `json:"cod"`
	Message string `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
	City    City   `json:"city"`
}

type List struct {
	Dt         int        `json:"dt"`
	Main       Main       `json:"main"`
	Weather    []Weather1 `json:"weather"`
	Clouds     Clouds     `json:"clouds"`
	Wind       Wind       `json:"wind"`
	Visibility int        `json:"visibility"`
	Pop        float64    `json:"pop"`
	Rain       Rain       `json:"rain,omitempty"`
	Sys        Sys        `json:"sys"`
	DtTxt      string     `json:"dt_txt"`
}
type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}
type Weather1 struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Clouds struct {
	All int `json:"all"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}
type Rain struct {
	ThreeH float64 `json:"3h"`
}

type Sys struct {
	Pod string `json:"pod"`
}
