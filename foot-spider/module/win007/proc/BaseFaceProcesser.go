package proc

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"io/ioutil"
	"regexp"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"time"

	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
)

type BaseFaceProcesser struct {
	service.BFScoreService
	service.BFBattleService
	service.BFJinService
	service.BFFutureEventService
	service2.LeagueService
	//是否是单线程
	SingleThread       bool
	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

type FutureMatches struct {
	ID         int    `json:"id"`
	MatchTime  string `json:"matchTime"`
	LeagueID   int    `json:"leagueId"`
	LeagueName string `json:"leagueName"`
	HomeTeamID string `json:"homeTeamId"`
	HomeTeam   string `json:"homeTeam"`
	AwayTeamID string `json:"awayTeamId"`
	AwayTeam   string `json:"awayTeam"`
	Separator  int    `json:"seperator"`
}

type FutureHomeMatches struct {
	TeamID   int             `json:"teamId"`
	TeamName string          `json:"teamName"`
	Matches  []FutureMatches `json:"matches"`
}

type TeamPoints struct {
	TeamId      int    `json:"teamId"`
	TeamName    string `json:"teamName"`
	ReductPoint int    `json:"reductPoint"`
	Points      []struct {
		Name       string  `json:"name"`
		Total      int     `json:"total"`
		Win        int     `json:"win"`
		Draw       int     `json:"draw"`
		Loss       int     `json:"loss"`
		GetGoal    int     `json:"getGoal"`
		LossGoal   int     `json:"lossGoal"`
		NetGoal    int     `json:"netGoal"`
		Point      int     `json:"point"`
		Rank       int     `json:"rank"`
		WinScale   float64 `json:"winScale"`
		PointsKind string  `json:"pointsKind"`
	} `json:"points"`
	SclassId    int    `json:"sclassId"`
	CurSeason   string `json:"curSeason"`
	SclassSubId int    `json:"sclassSubId"`
}

type Match struct {
	ID           int         `json:"id"`
	LeagueID     int         `json:"leagueId"`
	LeagueName   string      `json:"leagueName"`
	MatchTime    string      `json:"matchTime"`
	HomeTeam     Team        `json:"homeTeam"`
	AwayTeam     Team        `json:"awayTeam"`
	LetGoal      []BetResult `json:"letgoal"`
	OU           []BetResult `json:"ou"`
	IsNeutrality bool        `json:"isNeutrality"`
	LeaguesKind  int         `json:"leaugsKind"`
}

type Team struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Score      int    `json:"score"`
	HalfScore  int    `json:"halfScore"`
	Corner     int    `json:"corner"`
	YellowCard int    `json:"yellowCard"`
	RedCard    int    `json:"redCard"`
	HalfCorner int    `json:"halfCorner"`
}

type BetResult struct {
	Kind   string  `json:"kind"`
	PanKou float64 `json:"panKou"`
	Result string  `json:"result"`
}

type BattleDataTeam struct {
	HomeTeamID int     `json:"homeTeamId"`
	AwayTeamID int     `json:"awayTeamId"`
	Matches    []Match `json:"matches"`
}

type JinDataTeam struct {
	TeamId   int     `json:"teamId"`
	TeamName string  `json:"teamName"`
	Matches  []Match `json:"matches"`
}

func GetBaseFaceProcesser() *BaseFaceProcesser {
	processer := &BaseFaceProcesser{}
	processer.Init()
	return processer
}

func (this *BaseFaceProcesser) Init() {
	//初始化参数值
	this.Win007idMatchidMap = map[string]string{}
}

func (this *BaseFaceProcesser) Setup(temp *BaseFaceProcesser) {
	//设置参数值
}

func (this *BaseFaceProcesser) Startup() {

	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "BaseFaceProcesser")
	for i, v := range this.MatchLastList {

		if !this.SingleThread && i%1000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			//先将前面的spider启动
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
			newSpider.SetThreadnum(8).Run()

			processer = GetBaseFaceProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "BaseFaceProcesser"+strconv.Itoa(i))
		}

		temp_flag := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(temp_flag)
		matchLastExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchLastExt)
		win007_id := matchLastExt.Sid

		processer.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_BASE_FACE_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
	newSpider.SetThreadnum(8).Run()

}

func (this *BaseFaceProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:", request.Url, p.Errormsg())
		return
	}

	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(request.Url), ".")[0]
	matchId := this.Win007idMatchidMap[win007Id]

	this.data_process(matchId, p)
}

//处理获取积分榜数据
func (this *BaseFaceProcesser) data_process(matchId string, p *page.Page) {

	//if strings.TrimSpace(matchId) != "2483939" {
	//	return
	//}
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(p.GetBodyStr()))
	dom.Find("script").Each(func(index int, item *goquery.Selection) {

		scriptContent := item.Text()
		if strings.Contains(scriptContent, "var jsonData") {
			var jsonDataMap map[string]interface{}
			jsonRegex := regexp.MustCompile(`var jsonData\s*=\s*(.+?);`)
			jsonData := jsonRegex.FindStringSubmatch(scriptContent)

			if len(jsonData) == 2 {
				jsonDataStr := jsonData[1]
				if err := json.Unmarshal([]byte(jsonDataStr), &jsonDataMap); err != nil {
					panic(err)
				}
				if strings.TrimSpace(matchId) != "2512789" {
					if err := ioutil.WriteFile("D:/jsonData.txt", []byte(jsonDataStr), 0644); err != nil {
						panic(err)
					}
				}
				//积分榜
				scoreSaveList := make([]interface{}, 0)
				scoreModifyList := make([]interface{}, 0)
				scoreList := this.score_process(matchId, jsonDataMap)
				for _, e := range scoreList {
					temp_id, exist := this.BFScoreService.Exist(e)
					if exist {
						e.Id = temp_id
						scoreModifyList = append(scoreModifyList, e)
					} else {
						scoreSaveList = append(scoreSaveList, e)
					}
				}
				this.BFScoreService.SaveList(scoreSaveList)
				this.BFScoreService.ModifyList(scoreModifyList)

				//对战历史
				battleSaveList := make([]interface{}, 0)
				battleModifyList := make([]interface{}, 0)
				battleList := this.battle_process(matchId, jsonDataMap)
				//fmt.Println("battleList len:", len(battleList))
				for _, e := range battleList {
					temp_id, exist := this.BFBattleService.Exist(e)
					if exist {
						e.Id = temp_id
						battleModifyList = append(battleModifyList, e)
					} else {
						battleSaveList = append(battleSaveList, e)
					}
				}

				this.BFBattleService.SaveList(battleSaveList)
				this.BFBattleService.ModifyList(battleModifyList)

				//近期对战
				jinSaveList := make([]interface{}, 0)
				jinModifyList := make([]interface{}, 0)
				jinListHomeMatches := this.jin_process(matchId, "homeMatches", jsonDataMap)
				jinListHomeMatches = removeDuplicatesElementsBFJin(jinListHomeMatches)
				for _, e := range jinListHomeMatches {
					if len(string(e.ScheduleID)) <= 0 {
						continue
					}
					temp_id, exist := this.BFJinService.Exist(e)
					if exist {
						e.Id = temp_id
						jinModifyList = append(jinModifyList, e)
					} else {
						jinSaveList = append(jinSaveList, e)
					}
				}
				this.BFJinService.SaveList(jinSaveList)
				this.BFJinService.ModifyList(jinModifyList)
				jinSaveList = jinSaveList[:0]
				jinModifyList = jinModifyList[:0]

				jinListAwayMatches := this.jin_process(matchId, "awayMatches", jsonDataMap)
				jinListAwayMatches = removeDuplicatesElementsBFJin(jinListAwayMatches)
				for _, e := range jinListAwayMatches {
					if len(string(e.ScheduleID)) <= 0 {
						continue
					}
					temp_id, exist := this.BFJinService.Exist(e)
					if exist {
						e.Id = temp_id
						jinModifyList = append(jinModifyList, e)
					} else {
						jinSaveList = append(jinSaveList, e)
					}
				}
				this.BFJinService.SaveList(jinSaveList)
				this.BFJinService.ModifyList(jinModifyList)

				//未来对战
				futureEventSaveList := make([]interface{}, 0)
				futureEventModifyList := make([]interface{}, 0)
				futureEventListHomeMatches := this.future_event_process(matchId, "homeMatches", jsonDataMap)
				futureEventListHomeMatches = removeDuplicatesElementsBFFutureEvent(futureEventListHomeMatches)
				for _, e := range futureEventListHomeMatches {
					temp_id, exist := this.BFFutureEventService.Exist(e)
					if exist {
						e.Id = temp_id
						futureEventModifyList = append(futureEventModifyList, e)
					} else {
						futureEventSaveList = append(futureEventSaveList, e)
					}
				}
				this.BFFutureEventService.SaveList(futureEventSaveList)
				this.BFFutureEventService.ModifyList(futureEventModifyList)
				futureEventSaveList = futureEventSaveList[:0]
				futureEventModifyList = futureEventModifyList[:0]

				futureEventListAwayMatches := this.future_event_process(matchId, "awayMatches", jsonDataMap)
				futureEventListAwayMatches = removeDuplicatesElementsBFFutureEvent(futureEventListAwayMatches)
				for _, e := range futureEventListAwayMatches {
					temp_id, exist := this.BFFutureEventService.Exist(e)
					if exist {
						e.Id = temp_id
						futureEventModifyList = append(futureEventModifyList, e)
					} else {
						futureEventSaveList = append(futureEventSaveList, e)
					}
				}
				this.BFFutureEventService.SaveList(futureEventSaveList)
				this.BFFutureEventService.ModifyList(futureEventModifyList)

			}
		}
	})
}

//处理获取积分榜数据
func (this *BaseFaceProcesser) score_process(matchId string, jsonDataMap map[string]interface{}) []*pojo.BFScore {
	data_list_slice := make([]*pojo.BFScore, 0)

	currentPoints := jsonDataMap["currentPoints"].(map[string]interface{})
	homePointsData, err := json.Marshal(currentPoints["homePoints"])
	if err != nil {
		base.Log.Error("转换为JSON出错: matchId", matchId, " err:", err)
		return data_list_slice
	}
	var homePoints TeamPoints
	if err := json.Unmarshal(homePointsData, &homePoints); err != nil {
		base.Log.Error("解析JSON出错: matchId", matchId, " err:", err)
		return data_list_slice
	}
	temp := this.parseTeamPointsData(matchId, homePoints)

	data_list_slice = append(data_list_slice, temp...)
	awayPointsData, err := json.Marshal(currentPoints["awayPoints"])
	if err != nil {
		base.Log.Error("转换为JSON出错: matchId", matchId, " err:", err)
		return data_list_slice
	}
	var awayPoints TeamPoints
	if err := json.Unmarshal(awayPointsData, &awayPoints); err != nil {
		base.Log.Info("解析JSON出错: matchId", matchId, " err:", err)
		return data_list_slice
	}
	temp = this.parseTeamPointsData(matchId, awayPoints)
	data_list_slice = append(data_list_slice, temp...)

	return data_list_slice
}

func (this *BaseFaceProcesser) parseTeamPointsData(matchId string, points TeamPoints) []*pojo.BFScore {
	if len(points.Points) == 0 {
		return nil
	}

	data_list_slice := make([]*pojo.BFScore, 0)
	for _, point := range points.Points {
		temp := new(pojo.BFScore)
		temp.MatchId = matchId
		temp.TeamId = strconv.Itoa(points.TeamId)
		temp.Type = point.Name
		temp.MatchCount = point.Total
		temp.WinCount = point.Win
		temp.DrawCount = point.Draw
		temp.FailCount = point.Loss
		temp.GetGoal = point.GetGoal
		temp.LossGoal = point.LossGoal
		temp.DiffGoal = point.NetGoal
		temp.Score = point.Point
		temp.Ranking = point.Rank
		temp.WinRate = point.WinScale
		data_list_slice = append(data_list_slice, temp)
	}

	return data_list_slice

}

//处理对战数据获取
func (this *BaseFaceProcesser) battle_process(matchId string, jsonDataMap map[string]interface{}) []*pojo.BFBattle {
	//request := p.GetRequest()
	data_list_slice := make([]*pojo.BFBattle, 0)

	vsMatches := jsonDataMap["vsMatches"].(map[string]interface{})
	vsMatchesjson, err := json.Marshal(vsMatches)
	if err != nil {
		base.Log.Error("JSON化出错: matchId", matchId, " err:", err)
		return data_list_slice
	}

	var battledata BattleDataTeam
	err = json.Unmarshal(vsMatchesjson, &battledata)
	if err != nil {
		base.Log.Error(" vsMatchesjson 转换为JSON出错: matchId", matchId, " err:", err)
		return data_list_slice
	}

	//fmt.Println("battledata.Matches len:", len(battledata.Matches))
	for _, v := range battledata.Matches {
		temp := new(pojo.BFBattle)
		//temp.Id = matchId
		temp.MatchId = matchId

		matchTime, _ := strconv.ParseInt(v.MatchTime, 10, 64)
		battleMatchDate := time.Unix(matchTime, 0)
		temp.BattleMatchDate = battleMatchDate
		temp.BattleLeagueId = strconv.Itoa(v.LeagueID)
		temp.BattleMainTeamId = strconv.Itoa(v.HomeTeam.ID)
		temp.BattleGuestTeamId = strconv.Itoa(v.AwayTeam.ID)
		temp.BattleMainTeamName = v.HomeTeam.Name
		temp.BattleGuestTeamName = v.AwayTeam.Name
		temp.BattleMainTeamHalfGoals = v.HomeTeam.HalfScore
		temp.BattleGuestTeamHalfGoals = v.AwayTeam.HalfScore
		temp.BattleMainTeamGoals = v.HomeTeam.Score
		temp.BattleGuestTeamGoals = v.AwayTeam.Score
		temp.BattleMainTeamRedCard = v.HomeTeam.RedCard
		temp.BattleMainTeamYellowCard = v.HomeTeam.YellowCard
		temp.BattleMainTeamCorner = v.HomeTeam.Corner
		temp.BattleGuestTeamRedCard = v.AwayTeam.RedCard
		temp.BattleGuestTeamYellowCard = v.AwayTeam.YellowCard
		temp.BattleGuestTeamCorner = v.AwayTeam.Corner

		if len(v.LetGoal) == 2 {
			if v.LetGoal[0].Kind == "HALF_TIME" {
				temp.BattleHalfLetgoal = v.LetGoal[0].PanKou
				temp.BattleHalfLetgoalResult = v.LetGoal[0].Result
				temp.BattleLetgoal = v.LetGoal[1].PanKou
				temp.BattleLetgoalResult = v.LetGoal[1].Result
			} else if v.LetGoal[0].Kind == "FULL_TIME" {
				temp.BattleHalfLetgoal = v.LetGoal[1].PanKou
				temp.BattleHalfLetgoalResult = v.LetGoal[1].Result
				temp.BattleLetgoal = v.LetGoal[0].PanKou
				temp.BattleLetgoalResult = v.LetGoal[0].Result
			}
		}
		if len(v.OU) == 2 {
			if v.OU[0].Kind == "HALF_TIME" {
				temp.BattleHalfOu = v.OU[0].PanKou
				temp.BattleHalfOuResult = v.OU[0].Result
				temp.BattleOu = v.OU[1].PanKou
				temp.BattleOuResult = v.OU[1].Result
			} else if v.OU[0].Kind == "FULL_TIME" {
				temp.BattleHalfOu = v.OU[1].PanKou
				temp.BattleHalfOuResult = v.OU[1].Result
				temp.BattleOu = v.OU[0].PanKou
				temp.BattleOuResult = v.OU[0].Result
			}
		}

		data_list_slice = append(data_list_slice, temp)
	}

	return data_list_slice
}

//处理对战数据获取
func (this *BaseFaceProcesser) jin_process(matchId string, mainGuess string, jsonDataMap map[string]interface{}) []*pojo.BFJin {
	data_list_slice := make([]*pojo.BFJin, 0)

	nearMatches, ok := jsonDataMap["nearMatches"].(map[string]interface{})
	if ok {
	} else {
		base.Log.Error("nearMatches 不是预期的类型 matchId:", matchId)
		return data_list_slice
	}

	homeMatches, ok := nearMatches[mainGuess].(map[string]interface{})
	if ok {
	} else {
		base.Log.Error("homeMatches 不是预期的类型 matchId:", matchId)
		return data_list_slice
	}

	vsMatchesjson, err := json.Marshal(homeMatches)
	if err != nil {
		base.Log.Error("JSON化出错: matchId", matchId, " err:", err)
		return data_list_slice
	}

	var jinDataTeam JinDataTeam
	err = json.Unmarshal(vsMatchesjson, &jinDataTeam)
	//fmt.Println("jinDataTeam.Matches len:", len(jinDataTeam.Matches))

	for _, v := range jinDataTeam.Matches {
		temp := new(pojo.BFJin)
		temp.ScheduleID = v.ID

		matchTime, _ := strconv.ParseInt(v.MatchTime, 10, 64)
		battleMatchDate := time.Unix(matchTime, 0)
		temp.MatchTimeStr = battleMatchDate.Format("2006-01-02 15:04:05")
		temp.SclassID = v.LeagueID
		temp.SclassName = v.LeagueName
		temp.HomeTeam = v.HomeTeam.Name
		temp.GuestTeam = v.AwayTeam.Name
		temp.HomeHalfScore = v.HomeTeam.HalfScore
		temp.GuestHalfScore = v.AwayTeam.HalfScore
		temp.HomeScore = v.HomeTeam.Score
		temp.GuestScore = v.AwayTeam.Score
		//temp.BattleMainTeamRedCard = v.HomeTeam.RedCard
		//temp.BattleMainTeamYellowCard = v.HomeTeam.YellowCard
		//temp.BattleMainTeamCorner = v.HomeTeam.Corner
		//temp.BattleGuestTeamRedCard = v.AwayTeam.RedCard
		//temp.BattleGuestTeamYellowCard = v.AwayTeam.YellowCard
		//temp.BattleGuestTeamCorner = v.AwayTeam.Corner
		//fmt.Println("v.AwayTeam.Name:", v.AwayTeam.Name)

		if len(v.LetGoal) == 2 {
			if v.LetGoal[0].Kind == "HALF_TIME" {
				temp.FirstLetgoalHalf = v.LetGoal[0].PanKou
				temp.ResultHalf = v.LetGoal[0].Result
				temp.Letgoal = v.LetGoal[1].PanKou
				temp.Result = v.LetGoal[1].Result
			} else if v.LetGoal[0].Kind == "FULL_TIME" {
				temp.FirstLetgoalHalf = v.LetGoal[1].PanKou
				temp.ResultHalf = v.LetGoal[1].Result
				temp.Letgoal = v.LetGoal[0].PanKou
				temp.Result = v.LetGoal[0].Result
			}
		}
		if len(v.OU) == 2 {
			if v.OU[0].Kind == "HALF_TIME" {
				temp.FirstOUHalf = v.OU[0].PanKou
				temp.ResultOUHalf = v.OU[0].Result
				temp.FirstOU = v.OU[1].PanKou
				temp.ResultOU = v.OU[1].Result
			} else if v.OU[0].Kind == "FULL_TIME" {
				temp.FirstOUHalf = v.OU[1].PanKou
				temp.ResultOUHalf = v.OU[1].Result
				temp.FirstOU = v.OU[0].PanKou
				temp.ResultOU = v.OU[0].Result
			}
		}
		data_list_slice = append(data_list_slice, temp)
	}

	return data_list_slice
}

/**
将让球转换类型
*/
func (this *BaseFaceProcesser) ConvertLetball(letball string) float64 {
	var lb_sum float64
	slb_arr := strings.Split(letball, "/")
	slb_arr_0, _ := strconv.ParseFloat(slb_arr[0], 10)
	if len(slb_arr) > 1 {
		if strings.Index(slb_arr[0], "-") != -1 {
			lb_sum = slb_arr_0 - 0.25
		} else {
			lb_sum = slb_arr_0 + 0.25
		}
	} else {
		lb_sum = slb_arr_0
	}

	return lb_sum
}

//处理获取示来对战数据
func (this *BaseFaceProcesser) future_event_process(matchId string, mainGuess string, jsonDataMap map[string]interface{}) []*pojo.BFFutureEvent {
	data_list_slice := make([]*pojo.BFFutureEvent, 0)

	future3Matches, ok := jsonDataMap["future3Matches"].(map[string]interface{})
	if ok {
	} else {
		base.Log.Info("future3Matches 不是预期的类型 matchId:", matchId)
		return data_list_slice
	}

	homeMatches, ok := future3Matches[mainGuess].(map[string]interface{})
	if ok {
	} else {
		base.Log.Info("mainGuess 不是预期的类型 matchId:", matchId)
		return data_list_slice
	}

	vsMatchesjson, err := json.Marshal(homeMatches)
	if err != nil {
		base.Log.Info("JSON化出错:", err)
		return data_list_slice
	}

	var futureHomeMatches FutureHomeMatches
	err = json.Unmarshal(vsMatchesjson, &futureHomeMatches)
	//fmt.Println("jinDataTeam.Matches len:", len(futureHomeMatches.Matches))

	for _, v := range futureHomeMatches.Matches {
		temp := new(pojo.BFFutureEvent)
		temp.TeamId = strconv.Itoa(futureHomeMatches.TeamID)
		temp.TeamName = futureHomeMatches.TeamName
		temp.MatchId = strconv.Itoa(v.ID)
		temp.EventLeagueId = strconv.Itoa(v.LeagueID)
		temp.EventLeagueName = v.LeagueName
		matchTime, _ := strconv.ParseInt(v.MatchTime, 10, 64)
		battleMatchDate := time.Unix(matchTime, 0)
		temp.EventMatchDate = battleMatchDate
		temp.EventMainTeamId = v.HomeTeamID
		temp.EventGuestTeamId = v.AwayTeamID
		temp.EventMainTeamName = v.HomeTeam
		temp.EventGuestTeamName = v.AwayTeam
		temp.IntervalDay = int(v.Separator)

		data_list_slice = append(data_list_slice, temp)
	}
	return data_list_slice
}

func (this *BaseFaceProcesser) Finish() {
	base.Log.Info("基本面分析抓取解析完成 \r\n")

}

func removeDuplicatesElementsBFJin(elements []*pojo.BFJin) []*pojo.BFJin {
	encountered := map[string]struct{}{}
	result := []*pojo.BFJin{}

	for _, v := range elements {
		key := fmt.Sprintf("%d-%d-%d", v.MatchTimeStr, v.HomeTeam, v.GuestTeam)
		if _, ok := encountered[key]; !ok {
			encountered[key] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func removeDuplicatesElementsBFFutureEvent(elements []*pojo.BFFutureEvent) []*pojo.BFFutureEvent {
	encountered := map[string]struct{}{}
	result := []*pojo.BFFutureEvent{}

	for _, v := range elements {
		key := fmt.Sprintf("%d-%d-%d", v.MatchId, v.TeamId, v.EventMatchDate)
		if _, ok := encountered[key]; !ok {
			encountered[key] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
