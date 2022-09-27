package msg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	"github.com/gin-gonic/gin"
)

const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
	TimeStart   = "2022-01-20 15:04:05"
)
const (
	weatherUrl      = "https://restapi.amap.com/v3/weather/weatherInfo?key=a1fbea6fb02c29d25eb8bfd94c854dce&city=310112&extensions=all&output=json"
	xinZuoUrl       = "https://api.jisuapi.com/astro/fortune?astroid=7&appkey=8fec4338291dc08c"
	historyTodayUrl = "https://api.jisuapi.com/todayhistory/query?appkey=8fec4338291dc08c&month=1&day=2"
	xiaoHuaUrl      = "https://api.jisuapi.com/xiaohua/text?pagenum=1&pagesize=1&sort=addtime&appkey=8fec4338291dc08c"
	feishuUrl       = "https://open.feishu.cn/open-apis/bot/v2/hook/43417ab6-4cf4-4983-a743-73e9db34b002"
	feishuUrlTest   = "https://open.feishu.cn/open-apis/bot/v2/hook/b19295fc-cf7c-49f8-bed5-f67c59b3b383"
	eachDayASeq     = "http://open.iciba.com/dsapi/"
)

var (
	weather        = &WeatherResp{}
	constellation  = &Constellation{}
	xiaoHua        = &XiaoHua{}
	tmp            = make(map[string]interface{})
	eachDayContent = &EachDayContent{}
)

var WeekDayMap = map[string]string{
	"Monday":    "一",
	"Tuesday":   "二",
	"Wednesday": "三",
	"Thursday":  "四",
	"Friday":    "五",
	"Saturday":  "六",
	"Sunday":    "日",
}

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
	case MEIRIYIJU:
		json.Unmarshal(body, eachDayContent)
	default:
		json.Unmarshal(body, &tmp)
	}
	return nil
}

func (m *MsgController) PushInfo(c *gin.Context) {
	jsonStr := strings.NewReader(generateTodayMsg())
	resp, err := http.Post(feishuUrl, "application/json", jsonStr)
	defer resp.Body.Close()

	if err != nil {
		log.Error("Post msg err", "err", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Read msg error", "err", err)
	}
	log.Info("Body", body)
}

func NewMsgController() *MsgController {
	return &MsgController{}
}

func generateTodayMsg() string {
	get(weatherUrl, WEATHER)
	get(xinZuoUrl, XINZUO)
	get(xiaoHuaUrl, XIAOHUA)
	get(eachDayASeq, MEIRIYIJU)

	yearInt := time.Now().Year()
	monthInt := time.Now().Month()
	dayInt := time.Now().Day()
	weathe := weather.Forecasts[0].Casts[0]
	constel := constellation.Result
	beginT, _ := time.Parse(TIME_LAYOUT, TimeStart)

	ret := fmt.Sprintf(`{
		"msg_type": "interactive",
		"card": {
			"header": {
			  "template": "blue",
			  "title": {
				"i18n": {
				  "zh_cn": "💕亲爱的，早上好💕"
				},
				"tag": "plain_text"
			  }
			},
			"i18n_elements": {
			  "zh_cn": [
				{
				  "tag": "div",
				  "text": {
					"content": "☀️ **<font color='green'> 天气预报来咯 </font>**\n  今天是：%v 年 %v 月 %v 日，星期%v \n  城市：上海闵行区\n  天气：%v\n  最高气温：%v℃\n  最低温度：%v℃\n  今天是我们在一起的第 %v 天❤️\n  Tips: TODO",
					"tag": "lark_md"
				  }
				},
				{
				  "tag": "hr"
				},
				{
				  "tag": "div",
				  "text": {
					"content": "🧙‍♀️ **<font color='red'> **星座运势（ ♎天秤座 ）** </font>**\n  🪄**今日运势**🪄:%v\n\n  🪄**本周运势**🪄\n    💰金钱：%v\n    🎗事业：%v\n    ❤️爱情：%v\n    🚴🏼‍♂️身体：%v",
					"tag": "lark_md"
				  }
				},
				{
				  "tag": "hr"
				},
				{
				  "tag": "div",
				  "text": {
					"content": "\"* %v *\"\n&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&#45;&#45;&#45;&#45; <font color='grey'> 💕 * From LvGJ & Miss U ~* 💕</font>",
					"tag": "lark_md"
				  }
				}
			  ]
			}
		  }
	
}`, yearInt, int(monthInt), dayInt, WeekDayMap[time.Now().Weekday().String()], weathe.Dayweather,
		weathe.Daytemp, weathe.Nighttemp, SubDays(beginT), constel.Today.PreSummary,
		constel.Week.Money, constel.Week.Career, constel.Week.Love, constel.Week.Health, eachDayContent.Content)
	log.Info("ret", ret)
	return ret
}

func SubDays(begin time.Time) (day int) {
	day = int(time.Now().Sub(begin).Hours() / 24)
	return
}
