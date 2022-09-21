package msg

const (
	WEATHER = iota
	XINZUO
	XIAOHUA
	HISTORYTODAY
)

type WeatherResp struct {
	Status    string
	Count     int
	Info      string
	Infocode  int64
	Forecasts []*Forecasts
}
type Forecasts struct {
	City       string
	Adcode     int
	Province   string
	Reporttime string
	Casts      []*Cast
}

type Cast struct {
	Date         string
	Week         int
	Dayweather   string
	Nightweather string
	Daytemp      string
	Nighttemp    string
	Daywind      string
	Nightwind    string
	Daypower     string
	Nightpower   string
}

type Constellation struct {
	Status int
	Msg    string
	Result Astro
}

type Astro struct {
	Astroid   int
	Astroname string
	Week      AstroWeek
	Today     AstroToday
}

type AstroWeek struct {
	Date   string
	Money  string
	Career string
	Love   string
	Health string
	Job    string
}

type AstroToday struct {
	Date       string
	PreSummary string
	star       string
	Color      string
	Number     int
	Summary    int
	Money      int
	Career     int
	Love       int
	Health     int
}
