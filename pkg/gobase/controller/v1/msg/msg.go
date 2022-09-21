package msg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	weatherUrl      = "https://restapi.amap.com/v3/weather/weatherInfo?key=a1fbea6fb02c29d25eb8bfd94c854dce&city=310112&extensions=all&output=json"
	xinZuoUrl       = "https://api.jisuapi.com/astro/fortune?astroid=7&appkey=8fec4338291dc08c"
	historyTodayUrl = "https://api.jisuapi.com/todayhistory/query?appkey=8fec4338291dc08c&month=1&day=2"
	xiaoHuaUrl      = "https://api.jisuapi.com/xiaohua/text?pagenum=1&pagesize=1&sort=addtime&appkey=8fec4338291dc08c"
)

var (
	weather       = &WeatherResp{}
	constellation = &Constellation{}
	xiaoHua       = &XiaoHua{}
	tmp           = make(map[string]interface{})
)

func get(url string, typ int) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	switch typ {
	case WEATHER:
		json.Unmarshal(body, weather)
	case XINZUO:
		json.Unmarshal(body, constellation)
	case XIAOHUA:
		json.Unmarshal(body, xiaoHua)
	default:
		json.Unmarshal(body, tmp)
	}
	fmt.Println(string(body))
	return nil
}
