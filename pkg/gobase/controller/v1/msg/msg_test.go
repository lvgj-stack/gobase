package msg

import "testing"

func TestGetWeather(t *testing.T) {
	err := get(weatherUrl, WEATHER)
	get(xinZuoUrl, XINZUO)
	get(xiaoHuaUrl, XIAOHUA)
	if err != nil {
		t.Error(err)
	}
	//t.Log(constellation.Result.Week.Love)
	t.Log(xiaoHua.Result.List[0].Content)
}
