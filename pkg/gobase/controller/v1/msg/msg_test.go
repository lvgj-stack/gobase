package msg

import (
	"fmt"
	"testing"
	"time"
)

func TestGetWeather(t *testing.T) {
	err := get(weatherUrl, WEATHER)
	get(xinZuoUrl, XINZUO)
	get(xiaoHuaUrl, XIAOHUA)
	if err != nil {
		t.Error(err)
	}
	t.Log(constellation.Result.Week.Love)
	t.Log(xiaoHua.Msg)
}

func TestPostMsg(t *testing.T) {
	NewMsgController().PushInfo()

}

func TestTmp(t *testing.T) {
	yearInt := time.Now().Year()
	monthInt := time.Now().Month()
	dayInt := time.Now().Day()
	t.Log(fmt.Sprintf(`今天是：%v 年 %v 月 %v 日，星期%v `, yearInt, int(monthInt), dayInt, int(time.Now().Weekday())))

}
