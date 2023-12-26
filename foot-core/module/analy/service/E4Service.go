package service

import (
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"time"
)

type E4Service struct {
	MyAnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *E4Service) ModelName() string {
	return "MyE3"
}

func (this *E4Service) AnalyTest() {
	this.MyAnalyService.AnalyTest(this)
}

/**
计算全部比赛数据
*/
func (this *E4Service) Analy(analyAll bool) {
	this.MyAnalyService.Analy(analyAll, this)
}

/**
计算临近比赛数据
*/
func (this *E4Service) Analy_Near() {
	this.MyAnalyService.Analy_Near(this)
}

func (this *E4Service) Tongji() {

}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
*/
func (this *E4Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyNewResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id

	alist, normal, unnormal := this.AsiaAllTrackService.AnalyOneMatch(matchId)
	if normal == -1 {
		return -1, temp_data
	}

	preResult := -1
	if alist[0].EPanKou > 0 && normal > 10 && normal > unnormal*10 {
		preResult = 3
	} else if normal > 20 {
		preResult = 1
	} else if unnormal > 10 {
		preResult = 0
	}

	var data *entity5.AnalyNewResult
	if len(temp_data.Id) > 0 {
		temp_data.MatchDate = v.MatchDate
		temp_data.PreResult = preResult
		temp_data.HitCount = normal
		temp_data.THitCount = unnormal
		temp_data.LetBall = alist[0].EPanKou
		data = temp_data
		//比赛结果
		data.Result = this.IsNewRight2Option(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyNewResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.SLetBall = alist[0].SPanKou
		data.LetBall = alist[0].EPanKou
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 3
		//比赛结果
		data.Result = this.IsNewRight2Option(v, data)
		return 0, data
	}

}
