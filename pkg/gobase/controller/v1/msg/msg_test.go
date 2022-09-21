package msg

import "testing"

func TestGetWeather(t *testing.T) {
	err := get(weatherUrl, WEATHER)
	get(xinZuoUrl, XINZUO)
	if err != nil {
		t.Error(err)
	}
	t.Log(constellation.Result.Week.Love)
}
